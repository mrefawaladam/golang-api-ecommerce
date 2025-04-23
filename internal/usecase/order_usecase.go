package usecase

import (
	"ecommerce-api/internal/domain"
	"ecommerce-api/internal/dto/response"
	"ecommerce-api/internal/repository"
)

type OrderUsecase interface {
	Checkout(userID uint) (*response.OrderResponse, error)
}

type orderUsecase struct {
	orderRepo repository.OrderRepository
	cartRepo  repository.CartRepository
}

func NewOrderUsecase(orderRepo repository.OrderRepository, cartRepo repository.CartRepository) OrderUsecase {
	return &orderUsecase{orderRepo, cartRepo}
}

func (u *orderUsecase) Checkout(userID uint) (*response.OrderResponse, error) {
	// Retrieve all cart items for the user
	cartItems, err := u.cartRepo.GetCartByUserID(userID)
	if err != nil {
		return nil, err
	}

	// Calculate the total price from the cart items
	total := 0.0
	for _, item := range cartItems {
		total += float64(item.Quantity) * item.Product.Price
	}

 	order := domain.Order{
		UserID: userID,
		Total:  total,
		Status: "completed", 
	}

	err = u.orderRepo.CreateOrder(&order)
	if err != nil {
		return nil, err
	}

	// Clear the cart after order completion
	if err := u.cartRepo.ClearCartByUserID(userID); err != nil {
		return nil, err
	}

	return &response.OrderResponse{
		ID:        order.ID,
		UserID:    order.UserID,
		Total:     order.Total,
		Status:    order.Status,
		CreatedAt: order.CreatedAt.String(),
		UpdatedAt: order.UpdatedAt.String(),
	}, nil
}
