package auth

import (
	"log"
	"net/http"
	"time"

	database "shopping/database/implement"
	model "shopping/model"

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

		token, err := jwt.ParseWithClaims(tokenString, &model.Token{}, func(token *jwt.Token) (interface{}, error) {
			return []byte("your-secret-key"), nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		} 
		
		claims, ok := token.Claims.(*model.Token)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
			c.Abort()
			return
		}

		// Pass the username in the context for further processing
		c.Set("user_id", claims.UserID)

		c.Next()
	}
}

func AuthenticateUser(email, password string) bool {
	// Connect to the database
	db, err := database.GetDB()
	if err != nil {
		log.Println(err)
		return false
	}
	defer db.Close()

	// Query the database to retrieve the stored password for the given email
	var storedPassword string
	err = db.QueryRow(`SELECT password FROM public."user" WHERE email = $1;`, email).Scan(&storedPassword)
	if err != nil {
		log.Println(err)
		return false
	}

	// Compare the stored password with the provided password
	if storedPassword == password {
		return true
	}

	return false
}

func GenerateToken(useId int) (string, error) {
	claims := &model.Token{
		UserID: useId ,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	tokenString, err := token.SignedString([]byte("your-secret-key"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
