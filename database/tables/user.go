package tables

import (
	model "shopping/model"
)

type User struct {
	tableName struct{} `sql:"user"`
	UserID    int      `sql:"user_id"`
	UserName  string   `sql:"user_name"`
	Email     string   `sql:"email"`
	Password  string   `sql:"password"`
	//CreditCard []string
}

func (usr *User) MapToModule() model.User {
	return model.User{
		UserID:   usr.UserID,
		UserName: usr.UserName,
		Email:    usr.Email,
		Password: usr.Password,
	}
}

func (u *User) Fill(usr *model.User) *User {
	return &User{
		UserID:   usr.UserID,
		UserName: usr.UserName,
		Email:    usr.Email,
		Password: usr.Password,
	}
}
