package categories

import (
	"arox-products/internal/models"
	"context"
	"github.com/jmoiron/sqlx"
)

type Store interface {
	ListCategories(ctx context.Context) ([]models.Category, error)
}

type store struct {
	db *sqlx.DB
}

func NewStore(db *sqlx.DB) Store {
	return &store{db: db}
}
