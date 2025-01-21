package services

import "github.com/google/wire"

var ProviderSet = wire.NewSet(
	NewProductService,
	wire.Bind(new(ProductServiceInterface), new(*ProductService)),
)
