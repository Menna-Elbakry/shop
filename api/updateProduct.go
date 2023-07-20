package api

import (
	"log"
	"net/http"
	database "shopping/database/implement"
	model "shopping/model"

	"github.com/gin-gonic/gin"
)

// UpdateProduct to modify product in database
func UpdateProduct(c *gin.Context) {
	db, err := database.GetDB()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to database"})
		return
	}
	defer db.Close()

	id := c.Param("id")

	var product model.Product

	//Update Query
	_, err = db.Exec(`UPDATE public."product" SET product_name=$1,quantity=$2,price=$3 WHERE product_id=$4;`, product.ProductName, product.Quantity, product.Price, id)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product updated successfully"})
}
