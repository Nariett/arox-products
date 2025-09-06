package products

import (
	"arox-products/internal/models"
)

func (s *store) List() ([]models.Product, error) {
	var products []models.Product
	query := `SELECT * FROM product ORDER BY id DESC`

	err := s.db.Select(&products, query)
	if err != nil {
		return nil, err
	}

	return products, nil
}
