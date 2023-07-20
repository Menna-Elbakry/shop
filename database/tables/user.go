package tables

import (
	model "shopping/model"
)

type User struct {
	tableName struct{} `sql:"user"`
	UserID    int      `sql:"id,pk,type:uuid"`
	UserName  string   `sql:"user_name"`
	Email     string   `sql:"email"`
	Password  string   `sql:"password"`
	Role      string   `sql:"role"`
}

func (usr *User) MapToModule() model.User {
	return model.User{
		UserID:   usr.UserID,
		UserName: usr.UserName,
		Email:    usr.Email,
		Password: usr.Password,
		Role:     usr.Role,
	}
}

func (u *User) Fill(usr *model.User) *User {
	return &User{
		UserID:   usr.UserID,
		UserName: usr.UserName,
		Email:    usr.Email,
		Password: usr.Password,
		Role:     usr.Role,
	}
}
