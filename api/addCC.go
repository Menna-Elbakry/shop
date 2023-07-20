package api

import (
	"fmt"
	"log"
	"net/http"
	database "shopping/database/implement"
	"shopping/model"

	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/card"
	"github.com/stripe/stripe-go/customer"
	"github.com/stripe/stripe-go/token"
)

const (
	strKey = "Your_Secret_Key"
)

func CreateCardToken(c *gin.Context, id string) (string, error) {
	//stripe secret key
	stripe.Key = strKey

	var tok model.CCToken

	//database connection
	db, err := database.GetDB()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Database connection established successfully.")
	defer db.Close()

	//insert card information to credit database
	_, err = db.Exec(`INSERT INTO public."credit_card" 
	(user_id, credit_number, expiration_month, "expiration_yaer","cvc","customer_id","token_stripe","card_id")
	VALUES ($1,$2,$3,$4,$5,$6,$7);`, id, tok.CardNumber, tok.ExpirationMonth, tok.ExpirationYear, tok.CVC, "", "", "")
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return "", nil
	}

	params := &stripe.TokenParams{
		Card: &stripe.CardParams{
			Number:   stripe.String(tok.CardNumber),
			ExpMonth: stripe.String(tok.ExpirationMonth),
			ExpYear:  stripe.String(tok.ExpirationYear),
			CVC:      stripe.String(tok.CVC),
		},
	}
	t, err := token.New(params)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error while creating credit token:": err.Error()})
	}
	c.JSON(http.StatusAccepted, gin.H{"token created successfully ": t.ID})

	//add token stripe id to credit database
	_, err = db.Exec(`UPDATE public."credit_card" SET token_stripe = $1 WHERE user_id = $2`, t.ID, tok.UserID)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update product quantity"})
		return "", nil
	}
	return t.ID, nil

}

func CreateCustomer(c *gin.Context) (string, error) {
	//stripe secret key
	stripe.Key = strKey

	var usr model.User
	params := &stripe.CustomerParams{
		Email: stripe.String(usr.Email),
		Name:  stripe.String(usr.UserName),
	}
	cust, _ := customer.New(params)
	c.JSON(http.StatusAccepted, gin.H{"Customer Created with ID ": cust.ID})

	//database connection
	db, err := database.GetDB()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Database connection established successfully.")
	defer db.Close()

	//get user_id
	id, err := db.Query(`SELECT user_id FROM public."user WHERE email=$1";`, usr.Email)
	if err != nil {

		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return "", nil
	}

	//add Customer Stripe ID to credit database
	_, err = db.Exec(`UPDATE public."credit_card" SET customer_id = $1 WHERE user_id = $2`, cust.ID, id)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update product quantity"})
		return "", nil
	}
	return cust.ID, nil

}

func CreateCard(c *gin.Context) {
	//stripe secrete key
	stripe.Key = strKey

	id := c.Param("id")

	//database connection
	db, err := database.GetDB()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Database connection established successfully.")
	defer db.Close()

	cstID, err := CreateCustomer(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error while creating customer": err.Error()})
		return
	}

	tokID, err := CreateCardToken(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error while creating card token": err.Error()})
		return
	}

	// Attach the card to the customer using the card token
	cardParams := &stripe.CardParams{
		Customer: stripe.String(cstID),
		Token:    stripe.String(tokID),
	}
	card, err := card.New(cardParams)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error while creating card": err.Error()})
		return
	}

	//add Card Stripe ID to credit database
	_, err = db.Exec(`UPDATE public."credit_card" SET card_id = $1 WHERE customer_id = $2`, card.ID, cstID)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update product quantity"})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"Card Created with ID ": card.ID})
}
