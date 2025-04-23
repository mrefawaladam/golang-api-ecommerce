package handler

import (
	"ecommerce-api/internal/dto/response"
	"ecommerce-api/internal/repository"
	"ecommerce-api/internal/usecase"
	"strconv"

	"github.com/labstack/echo/v4"
)

type OrderHandler struct {
	orderUsecase usecase.OrderUsecase
}

func NewOrderHandler(orderUsecase usecase.OrderUsecase, cartRepo repository.CartRepository) *OrderHandler {
	return &OrderHandler{orderUsecase}
}

func (h *OrderHandler) Checkout(c echo.Context) error {
	userID, err := strconv.ParseUint(c.Param("user_id"), 10, 64)
	if err != nil {
		res := response.BadRequest("User ID tidak valid")
		return c.JSON(res.StatusCode, res)
	}

	order, err := h.orderUsecase.Checkout(uint(userID))
	if err != nil {
		res := response.InternalServerError("Checkout gagal")
		return c.JSON(res.StatusCode, res)
	}

	res := response.Success("Checkout berhasil", order)
	return c.JSON(res.StatusCode, res)
}
