package images

import (
	"arox-products/internal/models"
	"arox-products/tests"
	"context"
	"database/sql"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestGetImagesWithId(t *testing.T) {
	ctx := context.Background()

	db := tests.CreateTestContainerWithMigrations(t)
	iStore := NewStore(db)

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
		Brand:       "Carhartt WIP",
		Name:        "T-shirt Hudson Pocket",
		CategoryId:  1,
		Price:       100,
		Description: sql.NullString{String: "Good t-shirt", Valid: true},
		Sizes:       []byte(sizesStr),
		IsActive:    true,
		CreatedAt:   time.Now(),
	}

	images := []models.Image{
		{
			IdProduct: 1,
			Url:       "http://example.com/",
			IsMain:    true,
			IsActive:  true,
		},
		{
			IdProduct: 1,
			Url:       "http://example.com/",
			IsMain:    false,
			IsActive:  true,
		},
		{
			IdProduct: 1,
			Url:       "http://example.com/",
			IsMain:    false,
			IsActive:  true,
		},
		{
			IdProduct: 1,
			Url:       "http://example.com/",
			IsMain:    false,
			IsActive:  false,
		},
	}

	createdCategory, err := tests.InsertCategory(db, category)
	require.NoError(t, err)
	assert.True(t, createdCategory.Id > 0)

	createdProduct, err := tests.InsertProduct(db, product)
	require.NoError(t, err)
	assert.True(t, createdProduct.Id > 0)

	for _, img := range images {
		createdImage, err := tests.InsertImage(db, img)
		require.NoError(t, err)
		assert.True(t, createdImage.Id > 0)
	}

	imagesById, err := iStore.GetImagesWithId(ctx, createdCategory.Id)
	require.NoError(t, err)
	assert.Len(t, imagesById, len(images)-1)
	assert.Equal(t, true, imagesById[0].IsMain)
}
