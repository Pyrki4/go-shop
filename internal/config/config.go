package config

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
	DBURL         string `mapstructure:"DB_URL"`
}

func Load() *Config {
	var cfg Config

	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Failed to read .env file: %v", err)
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatalf("Failed to unmarshal config: %v", err)
	}

	return &cfg
}
