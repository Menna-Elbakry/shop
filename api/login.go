package api

import (
	"log"
	"net/http"

	database "shopping/database/implement"
	model "shopping/model"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	db, err := database.GetDB()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to database"})
		return
	}
	defer db.Close()

	var user model.User
	err = db.QueryRow(`SELECT *FROM public."user" WHERE email =$1 And password =$2;`, user.Email, user.Password).Scan(&user.UserID, &user.UserName, &user.Email)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"error": err.Error()})
}
