package tables

import (
	model "shopping/model"
)

type Product struct {
	tableName struct{} `sql:"product"`
	ID        int      `sql:"id"`
	Name      string   `sql:"name"`
	Quantity  int      `sql:"quantity"`
	Price     float32  `sql:"price"`
}

func (prod *Product) MapToModule() model.Product {
	return model.Product{
		ID:       prod.ID,
		Name:     prod.Name,
		Quantity: prod.Quantity,
		Price:    float32(prod.Price),
	}
}

func (p *Product) Fill(prod *model.Product) *Product {
	return &Product{
		ID:       prod.ID,
		Name:     prod.Name,
		Quantity: prod.Quantity,
		Price:    float32(prod.Price),
	}
}
