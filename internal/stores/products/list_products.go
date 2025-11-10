package products

import (
	"arox-products/internal/models"
	"context"
)

// /подумать над запросом
func (s *store) ListProducts(ctx context.Context) ([]*models.ProductWithImage, error) {
	var products []*models.ProductWithImage

	query := `
			SELECT p.id,
			   p.brand,
			   p.name,
			   p.category_id,
			   p.price,
			   p.description,
			   p.sizes,
			   p.is_active,
			   p.created_at,
			   COALESCE(
				   json_agg(
					   json_build_object(
						   'id', i.id,
						   'url', i.url,
						   'is_main', i.is_main,
						   'is_active', i.is_active
					   )
					   ORDER BY i.is_main DESC NULLS LAST, i.id
				   ) FILTER (WHERE i.id IS NOT NULL),
				   '[]'
			   ) AS images
		FROM products p
		LEFT JOIN images i ON p.id = i.id_product
		GROUP BY p.id
		ORDER BY p.id DESC;`

	err := s.db.SelectContext(ctx, &products, query)
	if err != nil {
		return nil, err
	}

	return products, nil
}
