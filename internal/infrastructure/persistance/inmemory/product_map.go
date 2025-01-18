package inmemory

import (
	"sync"

	"github.com/graphictects/ecommerce/internal/application/domain/market"
)

type ProductMap struct {
	// products is a map that stores the products by id
	// - key: product id
	// - value: quantity of the product
	products map[string]uint64

	mu sync.Locker
}

func NewProductMap() *ProductMap {
	return &ProductMap{
		products: make(map[string]uint64),
		mu:       &sync.Mutex{},
	}
}

func (pm *ProductMap) Adquire(id string, quantity uint64) error {
	pm.mu.Lock()
	defer pm.mu.Unlock()

	available, ok := pm.products[id]
	if !ok {
		return market.ErrProductNotFound
	}

	if quantity > available {
		return market.ErrInsufficientProduct
	}

	pm.products[id] -= quantity

	return nil
}

func (pm *ProductMap) Restore(id string, quantity uint64) error {
	pm.mu.Lock()
	defer pm.mu.Unlock()

	pm.products[id] += quantity

	return nil
}