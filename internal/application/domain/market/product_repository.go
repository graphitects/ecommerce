package market

import "errors"

var (
	ErrProductNotFound	   = errors.New("product not found")
	ErrInsufficientProduct = errors.New("insufficient product")
)

type ProductRepository interface {
	Adquire(id string, quantity uint64) (Product, error)
	Restore(id string, quantity uint64) error
}
