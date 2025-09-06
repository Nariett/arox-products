package products

import (
	"arox-products/internal/models"
	"context"
	"github.com/jmoiron/sqlx"
)

type Store interface {
	ListProducts(ctx context.Context) ([]*models.Product, error)
	GetProductWithId(ctx context.Context, id int64) (*models.Product, error)
}

type store struct {
	db *sqlx.DB
}

func NewStore(db *sqlx.DB) Store {
	return &store{db: db}
}
