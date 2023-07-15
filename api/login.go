package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	auth "shopping/authMiddleware"
	database "shopping/database/implement"
	model "shopping/model"
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

	err = db.QueryRow(`SELECT *FROM public."user" WHERE email =$1 And password =$2;`, user.Email, user.Password).Scan(&user.UserID, &user.UserName, &user.Email)

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"error": err.Error()})
}
