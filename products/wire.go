package products

import (
	"github.com/google/wire"
)

var Set = wire.NewSet(
	wire.Bind(new(ProviderProductRepository), new(*ProductRepository)),
	wire.Bind(new(ProviderProductService), new(*ProductService)),
	wire.Bind(new(ProviderProductAPI), new(*ProductAPI)),
	ProvideProductRepository,
	ProvideProductService,
	ProvideProductAPI,
)
