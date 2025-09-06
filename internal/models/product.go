package models

import (
	"database/sql"
	"time"
)

type Product struct {
	ID          int64          `db:"id"`
	Brand       string         `db:"brand"`
	Name        string         `db:"name"`
	Category    int64          `db:"category"`
	Price       int64          `db:"price"`
	Description sql.NullString `db:"description"`
	Sizes       []byte         `db:"sizes"`
	IsActive    bool           `db:"is_active"`
	CreatedAt   time.Time      `db:"created_at"`
}
