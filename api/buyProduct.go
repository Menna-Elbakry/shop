package api

import (
	"net/http"

	"log"
	database "shopping/database/implement"
	model "shopping/model"

	"github.com/gin-gonic/gin"
)

// buy a product for a user
func BuyProduct(c *gin.Context) {
	db, err := database.GetDB()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to database"})
		return
	}
	defer db.Close()

	id := c.Param("id")

	var order model.Product
	err = c.ShouldBindJSON(&order)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	_, err = db.Exec("INSERT INTO order (name,quantity,price) VALUES ($1, $2,$3) SELECT * FROM product WHERE id=%s ", order.Name, order.Quantity, order.Price, id)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Product created successfully"})
	// TODO: Implement logic to handle buying a product for a user, including payment processing using Stripe

	c.JSON(http.StatusOK, gin.H{"message": "Product purchased successfully"})
}
