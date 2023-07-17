package auth

import (

	"time"
	model "shopping/model"
	"github.com/dgrijalva/jwt-go"
)

	// Query the database to retrieve the stored password for the given email
func GenerateToken(useId int) (string, error){
	claims := &model.Token{
		UserID: useId ,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	tokenString, err := token.SignedString([]byte("Secret"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
