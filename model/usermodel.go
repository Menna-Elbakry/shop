package model

import (
	"errors"
)

// User model
type User struct {
	UserID   int    `json:"user_id"`
	UserName string `json:"user_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

func (u *User) Validate() error {
	if u.UserID == 0 {
		return errors.New("should have userId")
	}
	if u.UserName == " " {
		return errors.New("should have Name")
	}
	if u.Email == " " {
		return errors.New("should have Email")
	}
	if u.Password == " " {
		return errors.New("should have Password")
	}
	if u.Role == " " {
		return errors.New("should have Role")
	}
	return nil
}
