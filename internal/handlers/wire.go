package handlers

import "github.com/google/wire"

var ProviderSet = wire.NewSet(
	NewProductHandler,
	wire.Bind(new(ProductHandlerInterface), new(*ProductHandler)),
)
