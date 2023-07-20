package auth

import (
	"github.com/dgrijalva/jwt-go"
)

// GenerateToken generates a JWT token with the provided user ID
func GenerateToken(userID int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
	})

	return token.SignedString([]byte("Secret"))
}

// // Query the database to retrieve the stored password for the given email
// func GenerateToken(useId int) (string, error) {
// 	claims := &model.Token{
// 		UserID: useId,
// 		StandardClaims: jwt.StandardClaims{
// 			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
// 		},
// 	}

// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

// 	// Sign the token with the secret key
// 	tokenString, err := token.SignedString([]byte("Secret"))
// 	if err != nil {
// 		return "", err
// 	}

// 	return tokenString, nil
// }
