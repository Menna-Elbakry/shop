package api

import (
	"log"
	"net/http"
	database "shopping/database/implement"
	model "shopping/model"

	"github.com/gin-gonic/gin"
)

// GetAllProducts to view all the products stored in the database
func GetAllProducts(c *gin.Context) {
	db, err := database.GetDB()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to database"})
		return
	}
	defer db.Close()

	//Select Query
	rows, err := db.Query("SELECT id, name, price FROM products")
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get products"})
		return
	}
	defer rows.Close()

	var products []model.Product
	for rows.Next() {
		var product model.Product
		err := rows.Scan(&product.ProductID, &product.ProductName, &product.Price)
		if err != nil {
			log.Println(err)
			continue
		}
		products = append(products, product)
	}

	c.JSON(http.StatusOK, products)
}
