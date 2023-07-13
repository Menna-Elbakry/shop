package tables

import (
	model "shopping/model"
)

type Order struct {
	tableName struct{} `sql:"order"`
	ID        int      `sql:"id"`
	Name      string   `sql:"name"`
	Quantity  int      `sql:"quantity"`
	Price     float32  `sql:"price"`
}

func (ordr *Order) MapToModule() model.Order {
	return model.Order{
		ID:       ordr.ID,
		Name:     ordr.Name,
		Quantity: ordr.Quantity,
		Price:    float32(ordr.Price),
	}
}

func (o *Order) Fill(ordr *model.Order) *Order {
	return &Order{
		ID:       ordr.ID,
		Name:     ordr.Name,
		Quantity: ordr.Quantity,
		Price:    float32(ordr.Price),
	}
}
