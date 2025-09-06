package stores

import (
	"arox-products/internal/stores/categories"
	"arox-products/internal/stores/products"
)

type Stores interface {
	Products() products.Store
	Categories() categories.Store
}

type stores struct {
	products   products.Store
	categories categories.Store
}

func New(p products.Store, c categories.Store) Stores {
	return &stores{
		products:   p,
		categories: c,
	}
}

func (s *stores) Products() products.Store     { return s.products }
func (s *stores) Categories() categories.Store { return s.categories }
