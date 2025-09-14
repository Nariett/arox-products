package images

import (
	"arox-products/internal/models"
	"context"
)

func (s *store) GetImagesWithId(ctx context.Context, idProduct int64) ([]*models.Image, error) {
	var images []*models.Image

	query := `SELECT * FROM images WHERE id_product = $1 AND is_active = true ORDER BY images.is_main DESC NULLS LAST`

	err := s.db.SelectContext(ctx, &images, query, idProduct)
	if err != nil {
		return nil, err
	}

	return images, nil
}
