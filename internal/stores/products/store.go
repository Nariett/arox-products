package products

import (
	"arox-products/internal/models"
	"github.com/jmoiron/sqlx"
)

type Store interface {
	List() ([]models.Product, error)
}

type store struct {
	db *sqlx.DB
}

func NewStore(db *sqlx.DB) Store {
	return &store{db: db}
}
