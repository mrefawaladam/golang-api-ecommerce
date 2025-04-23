// internal/dto/request/simulation_request.go
package request

type SimulationRequest struct {
	InitialBalance  float64 `json:"initial_balance"`
	DepositAmount   float64 `json:"deposit_amount"`
	WithdrawAmount  float64 `json:"withdraw_amount"`
	NumGoroutines   int     `json:"num_goroutines"`
}
