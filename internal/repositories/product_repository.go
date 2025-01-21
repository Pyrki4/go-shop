package repositories

import (
	"github.com/Pyrki4/go-shop/internal/models"
	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

type ProductRepositoryInterface interface {
	GetAll() ([]models.Product, error)
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) GetAll() ([]models.Product, error) {
	var products []models.Product
	result := r.db.Find(&products)
	return products, result.Error
}
