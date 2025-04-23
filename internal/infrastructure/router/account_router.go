package router

import (
	"ecommerce-api/internal/handler"
	"ecommerce-api/internal/repository"
	"ecommerce-api/internal/usecase"

	"github.com/labstack/echo/v4"
)

func RegisterAccountRoutes(e *echo.Group) {
	repo := repository.NewAccountRepository()
	use := usecase.NewAccountUsecase(repo)
	handler := handler.NewAccountHandler(use)

	e.POST("/account/:user_id/deposit", handler.Deposit)
	e.POST("/account/:user_id/withdraw", handler.Withdraw)
	e.GET("/account/:user_id/balance", handler.GetBalance)
	e.POST("/account/:user_id/simulate", handler.SimulateConcurrent)
}
