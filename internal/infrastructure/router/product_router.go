package router

import (
	"ecommerce-api/internal/handler"
	"ecommerce-api/internal/middleware"
	"ecommerce-api/internal/repository"
	"ecommerce-api/internal/usecase"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func ProductRouter(e *echo.Group, db *gorm.DB) {
	productRepo := repository.NewProductRepository(db)
	productUC := usecase.NewProductUsecase(productRepo)
	productHandler := handler.NewProductHandler(productUC)

	e.GET("/products", productHandler.GetProducts, middleware.JWTMiddleware())
}
