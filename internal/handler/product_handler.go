package handler

import (
	"ecommerce-api/internal/dto/response"
	"ecommerce-api/internal/usecase"

	"github.com/labstack/echo/v4"
)

type ProductHandler struct {
	productUsecase usecase.ProductUsecase
}

func NewProductHandler(productUsecase usecase.ProductUsecase) *ProductHandler {
	return &ProductHandler{productUsecase}
}

func (h *ProductHandler) GetProducts(c echo.Context) error {
	search := c.QueryParam("search")
	products, err := h.productUsecase.GetAllProducts(search)
	if err != nil {
		res := response.InternalServerError("Gagal mengambil produk")
		return c.JSON(res.StatusCode, res)
	}
	res := response.Success("Berhasil mengambil data produk", products)
	return c.JSON(res.StatusCode, res)
}
