package tables

import (
	model "shopping/model"
)

type CreditCard struct {
			CardNumber   string `sql:"card_number"`
			ExpMonth     int `sql:"expiration_month"`
			ExpYear      int `sql:"expiration_month"`
			CVV          string `sql:"cvv"`
}

func (card *CreditCard) MapToModule() model.Order {
	return model.Order{
		CardNumber:card.CardNumber
		ExpMonth:card.ExpMonth
		ExpYear:card.ExpYear
		CVV:card.CVV
	}
}

func (c *CreditCard) Fill(card *model.CreditCard) *CreditCard {
	return &CreditCard{
		CardNumber:card.CardNumber
		ExpMonth:card.ExpMonth
		ExpYear:card.ExpYear
		CVV:card.CVV
	}
}
