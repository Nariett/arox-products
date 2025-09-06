package stores

import (
	"arox-products/internal/stores/products"
	"go.uber.org/fx"
)

func Construct() fx.Option {
	return fx.Provide(
		New,
		products.NewStore,
	)
}
