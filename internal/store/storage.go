package store

import (
	"arox-products/internal/store/products"
	"go.uber.org/fx"
)

func Construct() fx.Option {
	return fx.Provide(
		products.NewStore,
	)
}
