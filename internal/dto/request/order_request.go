package request

type OrderRequest struct {
	UserID uint `json:"user_id" validate:"required"`
}
