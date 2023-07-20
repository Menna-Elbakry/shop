package tables

import (
	model "shopping/model"
)

type CCToken struct {
	tableName       struct{} `sql:"token"`
	UserID          int      `sql:"user_id"`
	CardNumber      string   `sql:"credit_number"`
	ExpirationMonth string   `sql:"expiration_month"`
	ExpirationYear  string   `sql:"expiration_month"`
	CVC             string   `sql:"cvc"`
	CustomerID      string   `sql:"customer_id"`
	TokenStripe     string   `sql:"token_stripe"`
}

func (t *CCToken) MapToModule() model.CCToken {
	return model.CCToken{
		CardNumber:      t.CardNumber,
		ExpirationMonth: t.ExpirationMonth,
		ExpirationYear:  t.ExpirationYear,
		CVC:             t.CVC,
		CustomerID:      t.CustomerID,
		TokenStripe:     t.TokenStripe,
	}
}

func (c *CCToken) Fill(t *model.CCToken) *CCToken {
	return &CCToken{
		CardNumber:      t.CardNumber,
		ExpirationMonth: t.ExpirationMonth,
		ExpirationYear:  t.ExpirationYear,
		CVC:             t.CVC,
		CustomerID:      c.CustomerID,
		TokenStripe:     c.TokenStripe,
	}
}
