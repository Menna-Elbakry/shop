package api

import (
	"log"
	"net/http"
	database "shopping/database/implement"

	"github.com/gin-gonic/gin"
)

// DeleteProduct deletes the product record from the database
func DeleteProduct(c *gin.Context) {
	db, err := database.GetDB()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to database"})
		return
	}
	defer db.Close()

	id := c.Param("id")

	// Delete Query
	_, err = db.Exec(`DELETE FROM public."product" WHERE product_id = $1`, id)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}
