package categories

import (
	"arox-products/internal/models"
	"context"
	"github.com/jmoiron/sqlx"
)

//go:generate mockgen -source=$GOFILE -destination=./mock/$GOFILE
type Store interface {
	ListCategories(ctx context.Context) ([]*models.Category, error)
	GetCategoryWithId(ctx context.Context, idCategory int64) (*models.Category, error)
}

type store struct {
	db *sqlx.DB
}

func NewStore(db *sqlx.DB) Store {
	return &store{db: db}
}
