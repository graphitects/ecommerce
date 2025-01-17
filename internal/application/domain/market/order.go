package market

import "github.com/google/uuid"

type Order struct {
	// id is the unique identifier of the order.
	id uuid.UUID

	// customerID is the unique identifier of the customer.
	customerID string

	// products is the list of products in the order.
	products []Product
}

func (o *Order) GetID() uuid.UUID {
	return o.id
}

func (o *Order) GetCustomerID() string {
	return o.customerID
}

func (o *Order) GetProducts() []Product {
	return o.products
}

func (o *Order) TotalCost() float64 {
	var result float64
	for i := range o.products {
		result += o.products[i].GetPrice()
	}

	return result
}
