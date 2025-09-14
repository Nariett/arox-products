package stores

import (
	"arox-products/internal/stores/categories"
	"arox-products/internal/stores/images"
	"arox-products/internal/stores/products"
	"go.uber.org/fx"
)

func Construct() fx.Option {
	return fx.Provide(
		New,
		images.NewStore,
		products.NewStore,
		categories.NewStore,
	)
}
