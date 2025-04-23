package router

import (
	"ecommerce-api/internal/handler"
	"ecommerce-api/internal/repository"
	"ecommerce-api/internal/usecase"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func CartRouter(e *echo.Group, db *gorm.DB) {
	cartRepo := repository.NewCartRepository(db)
	cartUC := usecase.NewCartUsecase(cartRepo)
	cartHandler := handler.NewCartHandler(cartUC)

	e.POST("/cart/:user_id", cartHandler.AddToCart)
	e.GET("/cart/:user_id", cartHandler.GetCartItems)
}
