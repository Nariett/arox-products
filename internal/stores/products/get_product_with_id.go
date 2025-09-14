package products

import (
	"arox-products/internal/models"
	"context"
)

func (s *store) GetProductWithId(ctx context.Context, id int64) (*models.ProductWithImage, error) {
	var product models.ProductWithImage

	query := `SELECT * FROM products WHERE id = $1`

	err := s.db.GetContext(ctx, &product, query, id)
	if err != nil {
		return nil, err
	}

	return &product, nil
}
