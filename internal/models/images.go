package models

type Image struct {
	Id        int64  `db:"id" json:"id"`
	IdProduct int64  `db:"id_product" json:"id_product"`
	Url       string `db:"url" json:"url"`
	IsMain    bool   `db:"is_main" json:"is_main"`
	IsActive  bool   `db:"is_active" json:"is_active"`
}

type Images struct {
	Images []Image `json:"images"`
}
