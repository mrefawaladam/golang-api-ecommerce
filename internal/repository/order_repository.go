package repository

import (
	"ecommerce-api/internal/domain"

	"gorm.io/gorm"
)

type OrderRepository interface {
	CreateOrder(order *domain.Order) error
}

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepository{db}
}

func (r *orderRepository) CreateOrder(order *domain.Order) error {
	return r.db.Create(order).Error
} 
