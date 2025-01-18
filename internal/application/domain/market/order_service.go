package market

import "errors"

var (
	// ErrProductVersionMismatch is returned when the product version does not match.
	ErrProductVersionMismatch = errors.New("product version mismatch")
)

type CardInfo struct {
	Number string
	Holder string
	Exp    string
	CVV    string
}

type OrderService interface {
	// CreateOrder creates a new order.
	CreateOrder(customerID string, products []Product, cardInfo CardInfo) (*Order, error)
}