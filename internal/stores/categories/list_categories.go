package categories

import (
	"arox-products/internal/models"
	"context"
)

func (s *store) ListCategories(ctx context.Context) ([]models.Category, error) {
	var categories []models.Category
	query := `SELECT * FROM categories`

	err := s.db.SelectContext(ctx, &categories, query)
	if err != nil {
		return nil, err
	}

	return categories, nil
}
