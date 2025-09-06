package stores

import (
	"arox-products/internal/stores/products"
)

type Stores interface {
	Products() products.Store
}

type stores struct {
	products products.Store
}

func New(p products.Store) Stores {
	return &stores{
		products: p,
	}
}

func (s *stores) Products() products.Store { return s.products }
