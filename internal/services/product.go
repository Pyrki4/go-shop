package services

import (
	"github.com/Pyrki4/go-shop/internal/models"
	"github.com/Pyrki4/go-shop/internal/repositories"
)

type ProductService struct {
	repo *repositories.ProductRepository
}

func NewProductService(repo *repositories.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) GetAllProducts() []models.Product {
	products, _ := s.repo.GetAll()
	return products
}
