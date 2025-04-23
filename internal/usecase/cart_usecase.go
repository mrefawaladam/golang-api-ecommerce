package usecase

import (
	"ecommerce-api/internal/domain"
	"ecommerce-api/internal/dto/request"
	"ecommerce-api/internal/dto/response"
	"ecommerce-api/internal/repository"
)

type CartUsecase interface {
	AddItemToCart(userID uint, req request.CartRequest) (*response.CartResponse, error)
	GetCartItems(userID uint) ([]response.CartResponse, error)
}

type cartUsecase struct {
	cartRepo repository.CartRepository
}

func NewCartUsecase(cartRepo repository.CartRepository) CartUsecase {
	return &cartUsecase{cartRepo}
}

func (u *cartUsecase) AddItemToCart(userID uint, req request.CartRequest) (*response.CartResponse, error) {
	// Create Cart Domain
	cart := domain.Cart{
		UserID:    userID,
		ProductID: req.ProductID,
		Quantity:  req.Quantity,
	}

	err := u.cartRepo.AddToCart(&cart)
	if err != nil {
		return nil, err
	}

	// Return Response
	return &response.CartResponse{
		ID:        cart.ID,
		ProductID: cart.ProductID,
		Quantity:  cart.Quantity,
		CreatedAt: cart.CreatedAt.String(),
		UpdatedAt: cart.UpdatedAt.String(),
	}, nil
}

func (u *cartUsecase) GetCartItems(userID uint) ([]response.CartResponse, error) {
	carts, err := u.cartRepo.GetCartByUserID(userID)
	if err != nil {
		return nil, err
	}

	var cartResponses []response.CartResponse
	for _, cart := range carts {
		cartResponses = append(cartResponses, response.CartResponse{
			ID:        cart.ID,
			ProductID: cart.ProductID,
			Quantity:  cart.Quantity,
			CreatedAt: cart.CreatedAt.String(),
			UpdatedAt: cart.UpdatedAt.String(),
		})
	}

	return cartResponses, nil
}
