package categories

import (
	"arox-products/internal/models"
	"context"
)

func (s *store) GetCategoryWithId(ctx context.Context, idCategory int64) (*models.Category, error) {
	var model models.Category

	query := `SELECT * FROM categories WHERE id = $1`

	err := s.db.GetContext(ctx, &model, query, idCategory)
	if err != nil {
		return nil, err
	}

	return &model, nil
}
