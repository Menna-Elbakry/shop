package api

import (
	"log"
	"net/http"
	auth "shopping/authMiddleware"
	database "shopping/database/implement"
	model "shopping/model"
	//"strconv"

	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	Email string `json:"email"`
	Token string `json:"token"`
}

// Login authenticates the user and generates a JWT token
func Login(c *gin.Context) {
	db, err := database.GetDB()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to database"})
		return
	}
	defer db.Close()

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}
	var request LoginRequest

	// Query the database to retrieve the user based on the provided email
	var user model.User
	err = db.QueryRow(`SELECT user_id, email, password FROM public."user" WHERE email = $1;`, request.Email).Scan(&user.UserID, &user.Email, &user.Password)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	// Authenticate the user
	if !auth.AuthenticateUser(user.Password, request.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User Doesn't Exist"})
		return
	}
	// Generate JWT token
	tokenString, err := GenerateTokenString(user.UserID)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	_, err = db.Exec(`UPDATE public."token" SET token = $1 WHERE user_id = $2;`, tokenString, user.UserID)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update token"})
		return
	}

	response := LoginResponse{
		Email: user.Email,
		Token: tokenString,
	}

	c.JSON(http.StatusOK, response)}
