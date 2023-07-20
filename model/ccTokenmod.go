package model

type CCToken struct {
	UserID          int    `json:"user_id`
	CardNumber      string `json:"credit_number" binding:"required"`
	ExpirationMonth string `json:"expiration_month" binding:"required"`
	ExpirationYear  string `json:"expiration_year" binding:"required"`
	CVC             string `json:"cvc" binding:"required"`
	CustomerID      string `json:"customer_id"`
	TokenStripe     string `json:"token_stripe"`
}
