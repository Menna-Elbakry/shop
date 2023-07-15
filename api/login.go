package api

import (
	"log"
	"net/http"
	model "shopping/model"
auth "shopping/authMiddleware"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// Verify user credentials
	authenticated := auth.AuthenticateUser(user.UserName, user.Password)
	if !authenticated {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	// Generate JWT token
	token, err := auth.GenerateToken(user.UserName)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
