package api

import (
	"log"
	"net/http"

	database "shopping/database/implement"
	"shopping/model"

	"github.com/gin-gonic/gin"
)

// BuyProduct to add the product record to order database
func BuyProduct(c *gin.Context) {
	db, err := database.GetDB()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to the database"})
		return
	}
	defer db.Close()

	id := c.Param("id")

	var order model.Order
	err = c.ShouldBindJSON(&order)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	//select from product database
	pull := "SELECT * FROM product WHERE product_id = $1"
	rows, err := db.Query(pull, id)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query product"})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var productID, quantity int
		var price float64
		if err := rows.Scan(&productID, &quantity, &price); err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan product rows"})
			return
		}

		//add the selected product to order database
		insert := "INSERT INTO order (product_id, quantity, price) VALUES ($1, $2, $3)"
		_, err := db.Exec(insert, productID, quantity, price)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert product to order"})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product purchased successfully"})
}
