package internal

import (
	"github.com/Pyrki4/go-shop/internal/config"
	"github.com/Pyrki4/go-shop/internal/db"
	"github.com/Pyrki4/go-shop/internal/handlers"
	"github.com/Pyrki4/go-shop/internal/repositories"
	"github.com/Pyrki4/go-shop/internal/services"
	"github.com/google/wire"
)

func InitializeContainerManual() (*handlers.ProductHandler, error) {
	wire.Build(
		config.ProviderSet,
		db.ProviderSet,
		repositories.ProviderSet,
		services.ProviderSet,
		handlers.ProviderSet,
	)
	return &handlers.ProductHandler{}, nil
}
