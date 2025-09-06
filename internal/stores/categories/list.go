package categories

import "arox-products/internal/models"

func (s *store) List() ([]models.Category, error) {
	var categories []models.Product
	query := `SELECT * FROM categories ORDER BY id DESC`

	err := s.db.Select(&categories, query)
	if err != nil {
		return nil, err
	}

	return categories, nil
}
