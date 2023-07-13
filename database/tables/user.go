package tables

import (
	model "shopping/model"
)

type User struct {
	tableName struct{} `sql:"user"`
	ID        int      `sql:"id"`
	Name      string   `sql:"name"`
	Email     string   `sql:"email"`
	Password  string   `sql:"password"`
	//CreditCard []string
}

func (usr *User) MapToModule() model.User {
	return model.User{
		ID:       usr.ID,
		Name:     usr.Name,
		Email:    usr.Email,
		Password: usr.Password,
		//CreditCard: usr.CreditCard,
	}
}

func (u *User) Fill(usr *model.User) *User {
	return &User{
		ID:       usr.ID,
		Name:     usr.Name,
		Email:    usr.Email,
		Password: usr.Password,
		//CreditCard: usr.CreditCard,
	}
}
