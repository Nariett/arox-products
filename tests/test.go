package tests

import (
	"arox-products/internal/models"
	"context"
	pkg "github.com/Nariett/arox-pkg/db"
	"github.com/jmoiron/sqlx"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
	"testing"
	"time"
)

func CreateTestContainerWithMigrations(t *testing.T) *sqlx.DB {
	ctx := context.Background()

	pgContainer, err := postgres.RunContainer(ctx,
		testcontainers.WithImage("postgres:17-alpine"),
		postgres.WithDatabase("test-db"),
		postgres.WithUsername("postgres"),
		postgres.WithPassword("postgres"),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).WithStartupTimeout(10*time.Second)),
	)
	if err != nil {
		t.Fatal(err)
	}

	connStr, err := pgContainer.ConnectionString(ctx)
	if err != nil {
		t.Fatal(err)
	}

	connStr += " sslmode=disable"

	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		t.Fatal(err)
	}

	pkg.Migrate(db)

	t.Cleanup(func() {
		db.Close()
		if err := pgContainer.Terminate(ctx); err != nil {
			t.Fatalf("failed to terminate pgContainer: %s", err)
		}
	})

	return db
}

func InsertCategory(db *sqlx.DB, category models.Category) (models.Category, error) {
	ctx := context.Background()

	query := `INSERT INTO categories (name, slug) VALUES ($1, $2) returning *`

	var insertedCategory models.Category
	err := db.QueryRowxContext(ctx, query,
		category.Name,
		category.Slug,
	).StructScan(&insertedCategory)
	if err != nil {
		return models.Category{}, err
	}

	return insertedCategory, nil
}

func InsertImage(db *sqlx.DB, image models.Image) (*models.Image, error) {
	ctx := context.Background()

	query := `INSERT INTO images (id_product, url, is_main, is_active) 
			  VALUES ($1, $2, $3, $4)
			  RETURNING *`

	var insertedImage models.Image
	err := db.QueryRowxContext(ctx, query,
		image.IdProduct,
		image.Url,
		image.IsMain,
		image.IsActive,
	).StructScan(&insertedImage)
	if err != nil {
		return nil, err
	}

	return &insertedImage, nil
}

func InsertProduct(db *sqlx.DB, product models.Product) (*models.Product, error) {
	ctx := context.Background()

	query := `INSERT INTO products (brand, name, category_id, price, description, sizes, is_active, created_at) 
              VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
              RETURNING *`

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
