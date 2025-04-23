package handler

import (
	"ecommerce-api/internal/domain"
	"ecommerce-api/internal/dto/request"
	"ecommerce-api/internal/dto/response"
	"strconv"

	"github.com/labstack/echo/v4"
)

type AccountHandler struct {
	usecase domain.AccountUsecase
}

func NewAccountHandler(u domain.AccountUsecase) *AccountHandler {
	return &AccountHandler{usecase: u}
}

func (h *AccountHandler) Deposit(c echo.Context) error {
	var req request.TransactionRequest
	if err := c.Bind(&req); err != nil {
		res := response.BadRequest("Format request tidak valid")
		return c.JSON(res.StatusCode, res)
	}

	userID, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		res := response.BadRequest("User ID tidak valid")
		return c.JSON(res.StatusCode, res)
	}

	if err := h.usecase.Deposit(userID, req.Amount); err != nil {
		res := response.BadRequest("Gagal melakukan deposit: " + err.Error())
		return c.JSON(res.StatusCode, res)
	}

	res := response.Success("Deposit berhasil", nil)
	return c.JSON(res.StatusCode, res)
}
func (h *AccountHandler) SimulateConcurrent(c echo.Context) error {
	userID, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		res := response.BadRequest("User ID tidak valid")
		return c.JSON(res.StatusCode, res)
	}

	var req request.SimulationRequest
	if err := c.Bind(&req); err != nil {
		res := response.BadRequest("Format simulasi tidak valid")
		return c.JSON(res.StatusCode, res)
	}

	finalBalance, expectedBalance, err := h.usecase.SimulateConcurrent(
		userID,
		req.InitialBalance,
		req.DepositAmount,
		req.WithdrawAmount,
		req.NumGoroutines,
	)
	
	if err != nil {
		res := response.InternalServerError("Gagal menjalankan simulasi")
		return c.JSON(res.StatusCode, res)
	}

	res := response.Success("Simulasi selesai", map[string]interface{}{
		"expected_balance": expectedBalance,
		"final_balance":    finalBalance,
	})
	return c.JSON(res.StatusCode, res)
}

func (h *AccountHandler) Withdraw(c echo.Context) error {
	var req request.TransactionRequest
	if err := c.Bind(&req); err != nil {
		res := response.BadRequest("Format request tidak valid")
		return c.JSON(res.StatusCode, res)
	}

	userID, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		res := response.BadRequest("User ID tidak valid")
		return c.JSON(res.StatusCode, res)
	}

	if err := h.usecase.Withdraw(userID, req.Amount); err != nil {
		res := response.BadRequest("Gagal melakukan withdraw: " + err.Error())
		return c.JSON(res.StatusCode, res)
	}

	res := response.Success("Withdraw berhasil", nil)
	return c.JSON(res.StatusCode, res)
}

func (h *AccountHandler) GetBalance(c echo.Context) error {
	userID, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		res := response.BadRequest("User ID tidak valid")
		return c.JSON(res.StatusCode, res)
	}

	balance, err := h.usecase.GetBalance(userID)
	if err != nil {
		res := response.InternalServerError("Gagal mengambil saldo")
		return c.JSON(res.StatusCode, res)
	}

	res := response.Success("Saldo berhasil diambil", map[string]interface{}{
		"balance": balance,
	})
	return c.JSON(res.StatusCode, res)
}