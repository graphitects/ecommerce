package market

import "github.com/google/uuid"

type ProductAttributes struct {
	Name        string
	Model       string
	Category    string
	Description string
}

type Product struct {
	id uuid.UUID

	attr ProductAttributes

	price float64
}

func (p *Product) GetID() uuid.UUID {
	return p.id
}

func (p *Product) GetPrice() float64 {
	return p.price
}

func (p *Product) GetAttributes() ProductAttributes {
	return p.attr
}

func NewProduct(id uuid.UUID, attr ProductAttributes, price float64) Product {
	return Product{id: id, attr: attr, price: price}
}
