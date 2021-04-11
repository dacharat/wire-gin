package products

import (
	"github.com/google/wire"

	"github.com/dacharat/wire-gin/database"
)

var Set = wire.NewSet(
	database.NewDatabase,
	ProvideProductRepository,
	ProvideProductService,
	ProvideProductAPI,
)
