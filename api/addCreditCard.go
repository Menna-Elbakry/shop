package api

import (
	"log"
	"net/http"
	database "shopping/database/implement"
	model "shopping/model"
	"strings"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/paymentmethod"
)

// AddCreditCard adds a credit card to the database
func AddCreditCard(c *gin.Context) {
	db, err := database.GetDB()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to database"})
		return
	}
	defer db.Close()

	var creditCard model.CreditCard
	err = c.ShouldBindJSON(&creditCard)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	params := &stripe.PaymentMethodParams{
		Type: stripe.String("card"),
		Card: &stripe.PaymentMethodCardParams{
			Number:   stripe.String(creditCard.CardNumber),
			ExpMonth: stripe.String(creditCard.ExpirationMonth),
			ExpYear:  stripe.String(creditCard.ExpirationYear),
			CVC:      stripe.String(creditCard.CVV),
		},
	}

	pm, err := paymentmethod.New(params)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create payment method"})
		return
	}

	// Insert the credit card into the database
	_, err = db.Exec(`
		INSERT INTO public."credit_card" (card_number, exp_month, exp_year, cvv)
		VALUES ($1, $2, $3, $4);
	`, stripCardNumber(creditCard.CardNumber), creditCard.ExpirationMonth, creditCard.ExpirationYear, creditCard.CVV)

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Credit card added successfully"})

	fmt.Println("Payment Method ID:", pm.ID)
	fmt.Println("Payment Method Type:", pm.Type)
	fmt.Println("Payment Method Created:", pm.Created)
}

// Helper function to strip non-numeric characters from the card number
func stripCardNumber(cardNumber string) string {
	return strings.Map(func(r rune) rune {
		if r >= '0' && r <= '9' {
			return r
		}
		return -1
	}, cardNumber)
}
