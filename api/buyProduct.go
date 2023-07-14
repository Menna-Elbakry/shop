package api

import (
	"log"
	"net/http"

	database "shopping/database/implement"
	model "shopping/model"

	"github.com/gin-gonic/gin"
)

// BuyProduct adds the product record to the order database and updates the quantity field
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
	// err = c.ShouldBindJSON(&order)
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }

	updateFields := make(map[string]interface{})

	var product model.Product
	// err = c.ShouldBindJSON(&product)
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }

	if product.Quantity > 0 {
		updateFields["quantity"] = product.Quantity
	}

	var currentQuantity int
	err = db.QueryRow(`SELECT quantity FROM public."product" WHERE product_id = $1`, id).Scan(&currentQuantity)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	newQuantity := currentQuantity - 1

	_, err = db.Exec(`UPDATE public."product" SET quantity = $1 WHERE product_id = $2`, newQuantity, id)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update product quantity"})
		return
	}

	_, err = db.Exec(`INSERT INTO public."order" (order_id, product_id, quantity) 
	VALUES ($1, $2, $3);`, order.OrderID, id, 1)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Product purchased successfully"})
}
