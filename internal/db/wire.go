package db

import (
	"github.com/Pyrki4/go-shop/internal/config"
	"github.com/google/wire"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var ProviderSet = wire.NewSet(
	NewDB,
)

func NewDB(cfg *config.Config) (*gorm.DB, error) {
	return gorm.Open(postgres.Open(cfg.DBURL), &gorm.Config{})
}
