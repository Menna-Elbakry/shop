package database

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token not provided"})
			c.Abort()
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Verify the token's signing algorithm is matching the one that was used
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			// Change this secret key with your own secret key
			secretKey := []byte("Secret")
			return secretKey, nil
		})

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// Pass the user ID in the context to be used in subsequent handlers
		c.Set("userID", claims["id"])

		// Proceed to the next handler
		c.Next()
	}
}
func AuthenticateUser(db *sql.DB, username, password string) bool {
	query := `
        SELECT COUNT(*) FROM user WHERE username = $1 AND password = $2
    `
	// Execute the query
	var count int
	err := db.QueryRow(query, username, password).Scan(&count)
	if err != nil {
		log.Println(err)
		return false
	}

	// If the count is greater than 0, the username and password are valid
	return count > 0
}

func GenerateJWT(username string) (string, error) {
	// Define the expiration time for the token
	expirationTime := time.Now().Add(24 * time.Hour)

	// the JWT claims include the user and time
	claims := jwt.MapClaims{
		"username": username,
		"exp":      expirationTime.Unix(),
	}

	secretKey := []byte("Secret")

	// Create the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token to generate the final JWT token
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
