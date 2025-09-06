package categories

import (
	"arox-products/internal/models"
	"github.com/jmoiron/sqlx"
)

type Store interface {
	List() ([]models.Category, error)
}

type store struct {
	db *sqlx.DB
}

func NewStore(db *sqlx.DB) Store {
	return &store{db: db}
}
