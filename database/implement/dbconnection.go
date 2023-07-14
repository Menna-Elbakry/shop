package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "12345"
	dbname   = "postgres"
)

// GetDB establishes database connection
func GetDB() (*sql.DB, error) {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the database: %w", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to ping the database: %w", err)
	}

	// Create the orders table if it doesn't exist
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS orders (
			order_id SERIAL PRIMARY KEY,
			product_id INTEGER REFERENCES products(product_id),
			quantity INTEGER,
			price FLOAT
		)
	`)
	if err != nil {
		return nil, fmt.Errorf("failed to create orders table: %w", err)
	}

	// Create the products table if it doesn't exist
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS products (
			product_id SERIAL PRIMARY KEY,
			product_name TEXT,
			quantity INTEGER,
			price FLOAT
		)
	`)
	if err != nil {
		return nil, fmt.Errorf("failed to create products table: %w", err)
	}

	// Create the users table if it doesn't exist
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			user_id SERIAL PRIMARY KEY,
			user_name TEXT,
			email TEXT,
			password TEXT
		)
	`)
	if err != nil {
		return nil, fmt.Errorf("failed to create users table: %w", err)
	}

	return db, nil
}
