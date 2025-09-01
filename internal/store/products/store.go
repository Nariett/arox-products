package products

import "github.com/jmoiron/sqlx"

type Store interface{}

type store struct {
	db *sqlx.DB
}

func NewStore(db *sqlx.DB) Store {
	return &store{db: db}
}
