package categories

import (
	"arox-products/internal/models"
	"arox-products/tests"
	"context"
	"github.com/jmoiron/sqlx"
	"gotest.tools/v3/assert"
	"testing"
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

func TestListCategories(t *testing.T) {
	db := tests.CreateTestContainerWithMigrations(t)
	store := NewStore(db)

	categories := []models.Category{
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

	for _, category := range categories {
		err := InsertCategory(db, category)
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
