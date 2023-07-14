package main

import (
	"fmt"
	"log"
	a "shopping/api"
	database "shopping/database/implement"

	"github.com/gin-gonic/gin"
)

func main() {

	db, err := database.GetDB()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Database connection established successfully.")
	defer db.Close()

	route := gin.Default()

	// Admin APIs
	admin := route.Group("/admin")
	admin.POST("/product", database.AuthMiddleware(), a.CreateProduct)
	admin.PUT("/product/:id", database.AuthMiddleware(), a.UpdateProduct)
	admin.DELETE("/product/:id", database.AuthMiddleware(), a.DeleteProduct)

	// Public APIs
	public := route.Group("/public")
	public.GET("/products", a.GetAllProducts)
	public.POST("/signup", a.SignUp)
	public.POST("/login", database.AuthMiddleware(), a.Login)

	// User APIs
	user := route.Group("/user")
	user.POST("/buy-product", database.AuthMiddleware(), a.BuyProduct)

	r := route.Run(":8080")
	if r != nil {
		log.Fatal(r)
	}
}
