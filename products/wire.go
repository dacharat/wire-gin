package products

import (
	"github.com/google/wire"
)

var Set = wire.NewSet(
	ProvideProductRepository,
	ProvideProductService,
	ProvideProductAPI,
)
