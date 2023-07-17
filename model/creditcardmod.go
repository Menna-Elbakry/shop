package model


type CreditCard struct {
	CardNumber      string `json:"card_number" binding:"required"`
	ExpirationMonth int    `json:"expiration_month" binding:"required"`
	ExpirationYear  int    `json:"expiration_year" binding:"required"`
	CVV             string `json:"cvv" binding:"required"`
}

