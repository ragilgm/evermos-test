package payment

type PaymentRequest struct {
	Token string `json:"token" validate:"required"`
}
