package usecase

import (
	"ecommerce-api/internal/dto/response"
	"ecommerce-api/internal/repository"
)

type ProductUsecase interface {
	GetAllProducts(search string) ([]response.ProductResponse, error)
}

type productUsecase struct {
	productRepo repository.ProductRepository
}

func NewProductUsecase(productRepo repository.ProductRepository) ProductUsecase {
	return &productUsecase{productRepo}
}

func (u *productUsecase) GetAllProducts(search string) ([]response.ProductResponse, error) {
	products, err := u.productRepo.FindAll(search)
	if err != nil {
		return nil, err
	}

	var result []response.ProductResponse
	for _, p := range products {
		result = append(result, response.ProductResponse{
			ID:          p.ID,
			Name:        p.Name,
			Description: p.Description,
			Price:       p.Price,
			Stock:       p.Stock,
		})
	}
	return result, nil
}
