package products

import (
	"arox-products/internal/models"
	"arox-products/tests"
	"context"
	"database/sql"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"sort"
	"testing"
	"time"
)

func TestListProducts(t *testing.T) {
	db := tests.CreateTestContainerWithMigrations(t)
	store := NewStore(db)

	categories := []models.Category{
		{
			Name: "Кроссовки",
			Slug: "sneakers",
		},
		{
			Name: "Футболки",
			Slug: "t-shirts",
		},
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

	products := []models.Product{
		{
			Brand:       "Nike",
			Name:        "base",
			CategoryId:  2,
			Price:       12000,
			Description: sql.NullString{String: "good t-shirt", Valid: true},
			Sizes:       []byte(sizesStr),
			IsActive:    true,
			CreatedAt:   time.Now().In(time.UTC),
		},
		{
			Brand:       "Adidas",
			Name:        "Superstar",
			CategoryId:  1,
			Price:       12000,
			Description: sql.NullString{String: "good sneakers", Valid: true},
			Sizes:       []byte(sizesStr),
			IsActive:    true,
			CreatedAt:   time.Now().In(time.UTC),
		},
	}

	for _, category := range categories {
		createdCategory, err := tests.InsertCategory(db, category)
		require.NoError(t, err)
		assert.True(t, createdCategory.Id > 0)
	}

	for _, product := range products {
		createdProduct, err := tests.InsertProduct(db, product)
		require.NoError(t, err)
		assert.True(t, createdProduct.Id > 0)
	}

	response, err := store.ListProducts(context.Background())
	require.NoError(t, err)

	sort.Slice(response, func(i, j int) bool {
		return response[i].Id < response[j].Id
	})

	for i, v := range response {
		assert.Equal(t, products[i].Brand, v.Brand)
		assert.Equal(t, products[i].Name, v.Name)
		assert.Equal(t, products[i].CategoryId, v.CategoryId)
		assert.Equal(t, products[i].Price, v.Price)
		assert.Equal(t, products[i].Description.String, v.Description.String)
		assert.JSONEq(t, string(products[i].Sizes), string(v.Sizes))
		assert.Equal(t, products[i].IsActive, v.IsActive)
		assert.WithinDuration(t, products[i].CreatedAt, v.CreatedAt, time.Millisecond)
	}
}
