package api

import (
	"log"
	"net/http"
	"strings"
	"fmt"
	

	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/paymentmethod"
	"github.com/stripe/stripe-go/v72/token"
	"shopping/database/implement"
	"shopping/model"
)

// AddCreditCard adds a credit card to the database
func AddCreditCard(c *gin.Context) {
	db, err := database.GetDB()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to the database"})
		return
	}
	defer db.Close()

	stripeAPIKey := "sk_test_51NUoOjCnFRK4tebumo0BqzNvC97xvaivLl1gzQdpwtap2dW65S9N6SAPyeQmDItsi1of25xUnHUSA1cJbDkuH8AN00SYCebPwU"
	stripe.Key = stripeAPIKey

	var creditCard model.CreditCard
	err = c.ShouldBindJSON(&creditCard)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tokenParams := &stripe.TokenParams{
		Card: &stripe.CardParams{
			Number:   stripe.String(creditCard.CardNumber),
			ExpMonth: stripe.String(fmt.Sprintf("%02d", creditCard.ExpirationMonth)),
			ExpYear:  stripe.String(fmt.Sprintf("%d", creditCard.ExpirationYear)),
			CVC:      stripe.String(creditCard.CVV),
		},
	}
	testToken, err := token.New(tokenParams)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create a test token"})
		return
	}

	params := &stripe.PaymentMethodParams{
		Type: stripe.String("card"),
		Card: &stripe.PaymentMethodCardParams{
			Token: stripe.String(testToken.ID),
		},
	}

	pm, err := paymentmethod.New(params)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create a payment method"})
		return
	}

	_, err = db.Exec(`
		INSERT INTO public.credit_card (card_number, exp_month, exp_year, cvv)
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
