package tables

import (
	model "shopping/model"
)

type Order struct {
	tableName struct{} `sql:"order"`
	OrderID   int      `sql:"order_id"`
	ProductID int      `sql:"product_id"`
	Name      string   `sql:"name"`
	Quantity  int      `sql:"quantity"`
	Price     float32  `sql:"price"`
}

func (ordr *Order) MapToModule() model.Order {
	return model.Order{
		OrderID:   ordr.OrderID,
		ProductID: ordr.ProductID,
		Quantity:  ordr.Quantity,
		Price:     float32(ordr.Price),
	}
}

func (o *Order) Fill(ordr *model.Order) *Order {
	return &Order{
		OrderID:   ordr.OrderID,
		ProductID: ordr.ProductID,
		Quantity:  ordr.Quantity,
		Price:     float32(ordr.Price),
	}
}
