package images

import (
	"arox-products/internal/models"
	"context"
	"github.com/jmoiron/sqlx"
)

//go:generate mockgen -source=$GOFILE -destination=./mock/$GOFILE
type Store interface {
	GetImagesWithId(ctx context.Context, idProduct int64) ([]*models.Image, error)
}

type store struct {
	db *sqlx.DB
}

func NewStore(db *sqlx.DB) Store {
	return &store{db: db}
}
