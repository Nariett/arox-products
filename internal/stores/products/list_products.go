package products

import (
	"arox-products/internal/models"
	"context"
)

func (s *store) ListProducts(ctx context.Context) ([]*models.Product, error) {
	var products []*models.Product
	query := `SELECT * FROM product ORDER BY id DESC`

	err := s.db.SelectContext(ctx, &products, query)
	if err != nil {
		return nil, err
	}

	return products, nil
}
