package main

import (
	"fmt"
	"log"
	a "shopping/api"
	database "shopping/database/implement"

	"github.com/gin-gonic/gin"
)

// "fmt"
// "log"
// a "shopping/api"
// database "shopping/database/implement"

// "github.com/gin-gonic/gin"
// "github.com/stripe/stripe-go"

func main() {

	db, err := database.GetDB()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Database connection established successfully.")
	defer db.Close()

	route := gin.Default()

	// 	// Admin APIs
	admin := route.Group("/admin")
	admin.POST("/product", a.CreateProduct)
	admin.PUT("/product/:id", a.UpdateProduct)
	admin.DELETE("/product/:id", a.DeleteProduct)

	// Public APIs
	public := route.Group("/public")
	public.GET("/products", a.GetAllProducts)
	public.POST("/signup", a.SignUp)
	public.POST("/login", a.Login)

	// 	// User APIs
	user := route.Group("/user")
	user.POST("/buyproduct/:id", a.BuyProduct)
	user.POST("/create-card/:id", a.CreateCard) //userID

	r := route.Run(":8080")

	if r != nil {
		log.Fatal(r)
	}
}
