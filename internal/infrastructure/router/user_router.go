package router

import (
	"ecommerce-api/internal/handler"
	"ecommerce-api/internal/repository"
	"ecommerce-api/internal/usecase"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func UserRouter(e *echo.Group, db *gorm.DB) {
	userRepo := repository.NewUserRepository(db)
	userUC := usecase.NewUserUsecase(userRepo)
	h := handler.NewUserHandler(userUC)
 
	e.POST("/register", h.Register)
	e.POST("/login", h.Login)
}
