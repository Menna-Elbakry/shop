package api

import (
	"log"
	"net/http"
	database "shopping/database/implement"
	model "shopping/model"
"fmt"
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

	var user model.User
	var logs []model.User
	var request LoginRequest


//databse Connection
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
	

	// Query the database to retrieve the user based on the provided email
	rows,er := db.Query(`SELECT * FROM public."user" WHERE email = $1;`,request.Email)
	if er != nil {
		log.Println(er)
		c.JSON(http.StatusUnauthorized, gin.H{"error": er.Error()})
		return
	}

		// Scan the row values into the struct fields
	for rows.Next() {
		err := rows.Scan(&user.UserID,&user.UserName,&user.Email, &user.Password)
		if err != nil {
return 
		}
		// Append the product to the slice
		logs= append(logs, user)
	}

	if err = rows.Err(); err != nil {
return 
	}
	//print user row
	for _, user := range logs {
		if user.Password == request.Password {
			fmt.Println("Authorized: %v/n %v/n %v/n",user.UserID, user.UserName, user.Email)
			return 
		}
	}
   
	// Generate new token
	tokenString, err := GenerateTokenString(user.UserID)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}


	//update old token with new one
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
