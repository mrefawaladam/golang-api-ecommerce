package response

type BalanceResponse struct {
	Balance float64 `json:"balance"`
	Message string  `json:"message"`
}
