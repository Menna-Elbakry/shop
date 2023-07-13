package model

// User model
type User struct {
	ID       int    `json:"id"`
	Name     string `json:"firstName"`
	Email    string `json:"email"`
	Password string `json:"password"`
	//CreditCard []string `json:"creditCard"`
}
