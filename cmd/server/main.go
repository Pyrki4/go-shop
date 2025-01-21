package main

import (
	"log"

	"github.com/Pyrki4/go-shop/internal"
	"github.com/Pyrki4/go-shop/internal/config"
	"github.com/Pyrki4/go-shop/internal/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	// Загрузка конфигурации
	cfg := config.Load()

	// Инициализация контейнера зависимостей
	ProductHandler, err := internal.InitializeContainer()
	if err != nil {
		log.Fatalf("Failed to initialize container: %v", err)
	}

	// Инициализация сервера
	r := gin.Default()

	// Регистрация маршрутов
	registerRoutes(r, ProductHandler)

	// Запуск сервера
	log.Printf("Starting server on :%v", cfg.ServerAddress)
	if err := r.Run(cfg.ServerAddress); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func registerRoutes(r *gin.Engine, productHandler *handlers.ProductHandler) {
	r.GET("/products", productHandler.GetAllProducts)
}
