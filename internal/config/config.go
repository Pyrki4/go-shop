package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
	DBURL         string `mapstructure:"DB_URL"`
}

func Load() Config {
	var cfg Config
	viper.SetConfigFile(".env")
	viper.ReadInConfig()
	viper.Unmarshal(&cfg)
	return cfg
}
