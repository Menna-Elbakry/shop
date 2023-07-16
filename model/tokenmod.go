package model

	import "github.com/dgrijalva/jwt-go"

type Token struct {
	UserID int `json:"user_id"`
	jwt.StandardClaims

}
