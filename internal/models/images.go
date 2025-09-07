package models

type Image struct {
	Id        int64  `db:"id"`
	IdProduct int64  `db:"id_product"`
	Url       string `db:"url"`
	IsMain    bool   `db:"is_main"`
	IsActive  bool   `db:"is_active"`
}

type Images struct {
	Images []Image `json:"images"`
}
