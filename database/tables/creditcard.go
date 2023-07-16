package tables

import (
	model "shopping/model"
)

type CreditCard struct {
			CardNumber   string `sql:"card_number"`
			ExpMonth     string `sql:"exp_month"`
			ExpYear      string `sql:"exp_year"`
			CVC          string `sql:"cvc"`
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
