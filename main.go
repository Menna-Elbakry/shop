package main

import (
	"fmt"
	"log"
	"net/http"

	api "shopping/api"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// JWT middleware
func authMiddleware() gin.HandlerFunc {
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
			secretKey := []byte("YOUR_SECRET_KEY")
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
	}
}
func main() {
	r := gin.Default()

	apis := r.Group("/api")

	// Admin APIs
	admin := apis.Group("/admin")
	admin.POST("/product", authMiddleware(), api.CreateProduct)
	admin.PUT("/product/:id", authMiddleware(), api.UpdateProduct)
	admin.DELETE("/product/:id", authMiddleware(), api.DeleteProduct)

	// Public APIs
	public := apis.Group("/public")
	public.GET("/products", api.GetAllProducts)
	public.POST("/signup", api.SignUp)
	public.POST("/login", api.Login)

	// User APIs
	user := apis.Group("/user")
	user.PUT("/credit-card", authMiddleware(), api.AddCreditCard)
	user.POST("/buy-product", authMiddleware(), api.BuyProduct)

	err := r.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
