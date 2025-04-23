package router

import (
	"ecommerce-api/internal/handler"
	"ecommerce-api/internal/repository"
	"ecommerce-api/internal/usecase"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func OrderRouter(e *echo.Group, db *gorm.DB) {
	orderRepo := repository.NewOrderRepository(db)
	cartRepo := repository.NewCartRepository(db) 
	orderUC := usecase.NewOrderUsecase(orderRepo, cartRepo)  
	orderHandler := handler.NewOrderHandler(orderUC, cartRepo)

	e.POST("/checkout/:user_id", orderHandler.Checkout)
}
