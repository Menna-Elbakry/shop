package model

// User model
type User struct {
	UserID   int    `json:"user_id"`
	UserName string `json:"user_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	//CreditCard []string `json:"creditCard"`
}
