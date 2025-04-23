package request

type TransactionRequest struct {
	Amount float64 `json:"amount" validate:"required,gt=0"`
}
