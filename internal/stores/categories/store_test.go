package categories

import (
	"arox-products/internal/models"
	"arox-products/tests"
	"context"
	"github.com/stretchr/testify/require"
	"gotest.tools/v3/assert"
	"testing"
)

var (
	categories = []models.Category{
		{
			Name: "Футболки",
			Slug: "t-shirts",
		},
		{
			Name: "Джинсы",
			Slug: "jeans",
		},
		{
			Name: "Куртки",
			Slug: "jackets",
		},
		{
			Name: "Платья",
			Slug: "dresses",
		},
		{
			Name: "Юбки",
			Slug: "skirts",
		},
		{
			Name: "Рубашки",
			Slug: "shirts",
		},
		{
			Name: "Аксессуары",
			Slug: "accessories",
		},
		{
			Name: "Обувь",
			Slug: "shoes",
		},
		{
			Name: "Свитеры",
			Slug: "sweaters",
		},
		{
			Name: "Головные уборы",
			Slug: "headwear",
		},
	}
)

func TestListCategories(t *testing.T) {
	db := tests.CreateTestContainerWithMigrations(t)
	store := NewStore(db)

	for _, category := range categories {
		_, err := tests.InsertCategory(db, category)
		if err != nil {
			t.Fatal(err)
		}
	}

	value, err := store.ListCategories(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, len(categories), len(value))

}

func TestGetCategoryWithId(t *testing.T) {
	db := tests.CreateTestContainerWithMigrations(t)
	store := NewStore(db)

	for _, category := range categories {
		_, err := tests.InsertCategory(db, category)
		if err != nil {
			t.Fatal(err)
		}
	}

	value, err := store.GetCategoryWithId(context.Background(), int64(5))
	require.NoError(t, err)

	category := categories[4]

	assert.Equal(t, category.Name, value.Name)
	assert.Equal(t, category.Slug, value.Slug)
}
