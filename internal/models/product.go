package models

import (
	"database/sql"
	"time"
)

type Product struct {
	Id          int64          `db:"id"`
	Brand       string         `db:"brand"`
	Name        string         `db:"name"`
	CategoryId  int64          `db:"category_id"`
	Price       int64          `db:"price"`
	Description sql.NullString `db:"description"`
	Sizes       []byte         `db:"sizes"`
	IsActive    bool           `db:"is_active"`
	CreatedAt   time.Time      `db:"created_at"`
}

type ProductWithImage struct {
	Id          int64          `db:"id"`
	Brand       string         `db:"brand"`
	Name        string         `db:"name"`
	CategoryId  int64          `db:"category_id"`
	Price       int64          `db:"price"`
	Description sql.NullString `db:"description"`
	Sizes       []byte         `db:"sizes"`
	IsActive    bool           `db:"is_active"`
	CreatedAt   time.Time      `db:"created_at"`
	Images      []byte         `db:"images"`
}

type Size struct {
	Size  string `json:"size"`
	Count int64  `json:"count"`
}

type Sizes struct {
	Sizes []Size `json:"sizes"`
}
