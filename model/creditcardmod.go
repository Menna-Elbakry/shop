package model

type CreditCard struct {
	CardNumber       string `json:"card_number"`
	ExpirationMonth  string `json:"expiration_month"`
	ExpirationYear   string `json:"expiration_year"`
	CVV              string `json:"cvv"`
}
