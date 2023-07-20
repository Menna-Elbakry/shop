package tables

import (
	model "shopping/model"
)

type Product struct {
	tableName   struct{} `sql:"product"`
	ProductID   int      `sql:"product_id"`
	ProductName string   `sql:"product_name"`
	Quantity    int      `sql:"quantity"`
	Price       float32  `sql:"price"`
}

func (prod *Product) MapToModule() model.Product {
	return model.Product{
		ProductID:   prod.ProductID,
		ProductName: prod.ProductName,
		Quantity:    prod.Quantity,
		Price:       float32(prod.Price),
	}
}

func (p *Product) Fill(prod *model.Product) *Product {
	return &Product{
		ProductID:   prod.ProductID,
		ProductName: prod.ProductName,
		Quantity:    prod.Quantity,
		Price:       float32(prod.Price),
	}
}
