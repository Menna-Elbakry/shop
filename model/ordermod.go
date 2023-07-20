package model

// Order model
type Order struct {
	OrderID   int     `json:"order_id"`
	ProductID int     `json:"product_id"`
	Quantity  int     `json:"quantity"`
	Price     float32 `json:"price"`
}
