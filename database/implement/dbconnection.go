package database

import (
	"database/sql"
	"fmt"
	"log"
)

// PostgreSQL database connection string
const (
	host     = "localhost"
	port     = 5432
	user     = "your_username"
	password = "your_password"
	dbname   = "your_database_name"
)

// Get the PostgreSQL database connection
func GetDB() (*sql.DB, error) {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	// Create the orders table if it doesn't exist
	_, err = db.Exec(`
    CREATE TABLE IF NOT EXISTS orders (
	order_id SERIAL PRIMARY KEY,
	product_id INTEGER REFERENCES products(product_id),
	quantity INTEGER
)
`)
	if err != nil {
		log.Fatal(err)
	}

	// Create the products table if it doesn't exist
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS products (
			product_id SERIAL PRIMARY KEY,
			product_name TEXT,
			price FLOAT
		)
	`)
	if err != nil {
		log.Fatal(err)
	}

	// Create the users table if it doesn't exist
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS products (
			user_id SERIAL PRIMARY KEY,
			user_name TEXT,
			credit_cards TEXT
		)
	`)
	if err != nil {
		log.Fatal(err)
	}
	return db, nil
}
