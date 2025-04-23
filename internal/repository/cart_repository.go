package repository

import (
	"ecommerce-api/internal/domain"

	"gorm.io/gorm"
)

type CartRepository interface {
	AddToCart(cart *domain.Cart) error
	GetCartByUserID(userID uint) ([]domain.Cart, error)
	ClearCartByUserID(userID uint) error
}

type cartRepository struct {
	db *gorm.DB
}

func NewCartRepository(db *gorm.DB) CartRepository {
	return &cartRepository{db}
}

func (r *cartRepository) AddToCart(cart *domain.Cart) error {
	return r.db.Create(cart).Error
}

func (r *cartRepository) GetCartByUserID(userID uint) ([]domain.Cart, error) {
	var carts []domain.Cart
	err := r.db.Where("user_id = ?", userID).Preload("Product").Find(&carts).Error
	return carts, err
}

func (r *cartRepository) ClearCartByUserID(userID uint) error {
	return r.db.Where("user_id = ?", userID).Delete(&domain.Cart{}).Error
}
