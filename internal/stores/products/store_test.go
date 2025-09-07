package products

import (
	"arox-products/internal/models"
	"arox-products/tests"
	"context"
	"database/sql"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func InsertCategory(db *sqlx.DB, category models.Category) error {
	ctx := context.Background()

	query := `INSERT INTO categories (name, slug) VALUES ($1, $2)`

	_, err := db.ExecContext(ctx, query, category.Name, category.Slug)
	if err != nil {
		return err
	}

	return nil
}

func InsertProduct(db *sqlx.DB, product models.Product) (*models.Product, error) {
	ctx := context.Background()

	query := `INSERT INTO products (brand, name, category, price, description, sizes, is_active, created_at) 
              VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
              RETURNING id, brand, name, category, price, description, sizes, is_active, created_at`

	var insertedProduct models.Product
	err := db.QueryRowxContext(ctx, query,
		product.Brand,
		product.Name,
		product.CategoryId,
		product.Price,
		product.Description,
		product.Sizes,
		product.IsActive,
		product.CreatedAt,
	).StructScan(&insertedProduct)

	if err != nil {
		return nil, err
	}

	return &insertedProduct, nil
}

func TestListProducts(t *testing.T) {
	db := tests.CreateTestContainerWithMigrations(t)
	store := NewStore(db)

	category := models.Category{
		Name: "Футболки",
		Slug: "t-shirts",
	}

	sizesStr := `
		{
	  "sizes": [
		{
		  "size": "S",
		  "count": 10
		},
		{
		  "size": "M",
		  "count": 10
		},
		{
		  "size": "l",
		  "count": 5
		}
	  ]
	}
	`

	product := models.Product{
		Brand:       "Nike",
		Name:        "Air Max",
		CategoryId:  1,
		Price:       12000,
		Description: sql.NullString{String: "XD", Valid: true},
		Sizes:       []byte(sizesStr),
		IsActive:    true,
		CreatedAt:   time.Now(),
	}

	err := InsertCategory(db, category)
	if err != nil {
		t.Fatal(err)
	}

	value, err := InsertProduct(db, product)
	if err != nil {
		t.Fatal(err)
	}

	assert.True(t, value.ID > 0)

	response, err := store.ListProducts(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, 1, len(response))
}
