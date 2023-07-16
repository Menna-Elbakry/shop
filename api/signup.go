package api

import (
	"log"
	"net/http"
	auth "shopping/authMiddleware"
	database "shopping/database/implement"
	model "shopping/model"

	"github.com/gin-gonic/gin"
)

// SignUp to add new user to database
func SignUp(c *gin.Context) {
	db, err := database.GetDB()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to database"})
		return
	}
	defer db.Close()

	var user model.User

	err = c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err = db.Exec(`INSERT INTO public."user" (user_id, user_name, email, "password")
	VALUES ($1,$2,$3,$4);`, user.UserID, user.UserName, user.Email, user.Password)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Generate JWT token
	tokenString, err := GenerateTokenString(user.UserID)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	// Store the token in the database
	_, err = db.Exec(`INSERT INTO public."token" (user_id, token) VALUES ($1, $2);`, user.UserID, tokenString)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully", "token": tokenString})
}

func GenerateTokenString(userID int) (string, error) {
	tokenString, err := auth.GenerateToken(userID)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
