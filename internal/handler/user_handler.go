package handler

import (
	"ecommerce-api/internal/dto/request"
	"ecommerce-api/internal/dto/response"
	"ecommerce-api/internal/usecase"
	"ecommerce-api/internal/util"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	Usecase usecase.UserUsecase
}

func NewUserHandler(usecase usecase.UserUsecase) *UserHandler {
	return &UserHandler{usecase}
}

func (h *UserHandler) Register(c echo.Context) error {
	var req request.RegisterRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, response.MessageResponse{Message: "invalid input"})
	}
	if errs := util.ValidateStruct(req); errs != nil {
		return c.JSON(http.StatusBadRequest, response.ValidationError(errs))
	}

	if err := h.Usecase.Register(req); err != nil {
		return c.JSON(http.StatusInternalServerError, response.MessageResponse{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, response.Created("user registered successfully", nil))
}

func (h *UserHandler) Login(c echo.Context) error {
	var req request.LoginRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, response.MessageResponse{Message: "invalid input"})
	}
	if errs := util.ValidateStruct(req); errs != nil {
		return c.JSON(http.StatusBadRequest, response.ValidationError(errs))
	}
	token, err := h.Usecase.Login(req)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, response.MessageResponse{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, response.TokenResponse{Token: token})
}
