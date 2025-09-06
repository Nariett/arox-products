package models

type Category struct {
	Id   int64  `db:"id"`
	Name string `db:"name"`
	Slug string `db:"slug"`
}
