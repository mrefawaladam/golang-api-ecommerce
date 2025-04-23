package handler

import (
	"ecommerce-api/internal/dto/request"
	"ecommerce-api/internal/dto/response"
	"ecommerce-api/internal/usecase"
	"strconv"

	"github.com/labstack/echo/v4"
)

type CartHandler struct {
	cartUsecase usecase.CartUsecase
}

func NewCartHandler(cartUsecase usecase.CartUsecase) *CartHandler {
	return &CartHandler{cartUsecase}
}
 
func (h *CartHandler) AddToCart(c echo.Context) error {
	userID, err := strconv.ParseUint(c.Param("user_id"), 10, 64)
	if err != nil {
		res := response.BadRequest("User ID tidak valid")
		return c.JSON(res.StatusCode, res)
	}

	var req request.CartRequest
	if err := c.Bind(&req); err != nil {
		res := response.BadRequest("Format request tidak valid")
		return c.JSON(res.StatusCode, res)
	}

	cart, err := h.cartUsecase.AddItemToCart(uint(userID), req)
	if err != nil {
		res := response.InternalServerError("Gagal menambahkan item ke keranjang")
		return c.JSON(res.StatusCode, res)
	}

	res := response.Created("Item berhasil ditambahkan ke keranjang", cart)
	return c.JSON(res.StatusCode, res)
}

func (h *CartHandler) GetCartItems(c echo.Context) error {
	userID, err := strconv.ParseUint(c.Param("user_id"), 10, 64)
	if err != nil {
		res := response.BadRequest("User ID tidak valid")
		return c.JSON(res.StatusCode, res)
	}

	carts, err := h.cartUsecase.GetCartItems(uint(userID))
	if err != nil {
		res := response.InternalServerError("Gagal mengambil item dari keranjang")
		return c.JSON(res.StatusCode, res)
	}

	res := response.Success("Data keranjang berhasil diambil", carts)
	return c.JSON(res.StatusCode, res)
}