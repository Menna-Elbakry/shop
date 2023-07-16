package tables

import (
	model "shopping/model"
)

type Token struct {
	UserID int    `sql:"user_id"`
	Token  string `sql:"token"`
}
func (tok *Token) MapToModule() model.Token {
	return model.Order{
		UserID:   tok.UserID,
		TokenStr: tok.TokenStr,
	}
}

func (t *Token) Fill(tok *model.Token) *Token {
	return &Token{
		UserID:   tok.UserID,
		TokenStr: tok.TokenStr,
	}
}