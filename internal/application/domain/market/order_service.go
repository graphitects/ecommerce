package market

type OrderService interface {
	// CreateOrder creates a new order.
	CreateOrder(customerID string, products []Product) (*Order, error)
}
