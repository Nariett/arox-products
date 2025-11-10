package stores

import (
	"arox-products/internal/stores/categories"
	"arox-products/internal/stores/images"
	"arox-products/internal/stores/products"
)

//go:generate mockgen -source=$GOFILE -destination=./mock/$GOFILE
type Stores interface {
	Images() images.Store
	Products() products.Store
	Categories() categories.Store
}

type stores struct {
	images     images.Store
	products   products.Store
	categories categories.Store
}

func New(i images.Store, p products.Store, c categories.Store) Stores {
	return &stores{
		images:     i,
		products:   p,
		categories: c,
	}
}

func (s *stores) Images() images.Store         { return s.images }
func (s *stores) Products() products.Store     { return s.products }
func (s *stores) Categories() categories.Store { return s.categories }
