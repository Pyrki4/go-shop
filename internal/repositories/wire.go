package repositories

import (
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewProductRepository, wire.Bind(new(ProductRepositoryInterface), new(*ProductRepository)))
