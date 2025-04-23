package response

type OrderResponse struct {
	ID        uint    `json:"id"`
	UserID    uint    `json:"user_id"`
	Total     float64 `json:"total"`
	Status    string  `json:"status"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
}
