package main

import (
	"github.com/Pyrki4/go-shop/internal/config"
	"github.com/Pyrki4/go-shop/internal/handlers"
	"github.com/Pyrki4/go-shop/internal/repositories"
	"github.com/Pyrki4/go-shop/internal/services"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Загрузка конфигурации
	cfg := config.Load()

	// Инициализация базы данных
	db, err := gorm.Open(postgres.Open(cfg.DBURL), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	// Инициализация репозитория и сервиса
	productRepo := repositories.NewProductRepository(db)
	productService := services.NewProductService(productRepo)

	// Инициализация сервера
	r := gin.Default()

	// Регистрация маршрутов
	handlers.RegisterRoutes(r, productService)

	// Запуск сервера
	r.Run(cfg.ServerAddress)
}
