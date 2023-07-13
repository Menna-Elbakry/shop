package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	database "shopping/database/implement"
)

// buy a product for a user
func BuyProduct(c *gin.Context) {
	// TODO: Implement logic to handle buying a product for a user, including payment processing using Stripe

	c.JSON(http.StatusOK, gin.H{"message": "Product purchased successfully"})
}
	// Calculate the total price of the order
	var totalPrice float64

	_,err = database.QueryRow(`
		SELECT SUM(p.price * o.quantity)
		FROM orders o
		JOIN products p ON o.product_id = p.id
	`).Scan(&totalPrice)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Total Price: $%.2f\n", totalPrice)
