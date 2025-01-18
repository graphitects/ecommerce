package market

import "github.com/google/uuid"

type Order struct {
	// id is the unique identifier of the order.
	id uuid.UUID

	// customerID is the unique identifier of the customer.
	customerID string

	// products is the list of products in the order.
	products map[Product]uint64
}

func (o *Order) GetID() uuid.UUID {
	return o.id
}

func (o *Order) GetCustomerID() string {
	return o.customerID
}

func (o *Order) GetProducts() map[Product]uint64 {
	return o.products
}

func (o *Order) AddProduct(product Product, quantity uint64) {
	o.products[product] = quantity
}

func (o *Order) Total() float64 {
	var total float64
	for product, quantity := range o.products {
		total += product.GetPrice() * float64(quantity)
	}
	return total
}

func NewOrder(id uuid.UUID, customerID string, products map[Product]uint64) *Order {
	return &Order{
		id:         id,
		customerID: customerID,
		products:   products,
	}
}
