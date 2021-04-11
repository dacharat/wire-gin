// +build wireinject

package di

import (
	"github.com/google/wire"

	"github.com/dacharat/wire-gin/database"
	"github.com/dacharat/wire-gin/products"
)

func New() products.ProductAPI {
	wire.Build(
		database.Set,
		products.Set,
	)

	return products.ProductAPI{}
}
