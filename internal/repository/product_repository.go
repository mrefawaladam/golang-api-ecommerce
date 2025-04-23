package repository

import (
	"ecommerce-api/internal/domain"

	"gorm.io/gorm"
)

type ProductRepository interface {
	FindAll(search string) ([]domain.Product, error)
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db}
}

func (r *productRepository) FindAll(search string) ([]domain.Product, error) {
	var products []domain.Product
	query := r.db
	if search != "" {
		query = query.Where("name ILIKE ?", "%"+search+"%")
	}
	err := query.Find(&products).Error
	return products, err
}
