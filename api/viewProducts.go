package api

import (
	"fmt"
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

	rows, err := db.Query(`SELECT * FROM public."product";`)
	if err != nil {

		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	var products []model.Product

	for rows.Next() {
		var product model.Product

		// Scan the row values into the struct fields
		err := rows.Scan(&product.ProductID, &product.ProductName, &product.Price)
		if err != nil {
return 
		}

		// Append the product to the slice
		products = append(products, product)
	}

	if err = rows.Err(); err != nil {
return 
	}

	// Process the retrieved data
	for _, product := range products {
		fmt.Println(product.ProductID, product.ProductName, product.Price)
	}
	// log.Println(rows)
	c.JSON(http.StatusOK, products)
}
