package inmemory

import "github.com/graphictects/ecommerce/internal/application/domain/market"

type OrderMap struct {
	orders map[string]market.Order
}

func NewOrderMap() *OrderMap {
	return &OrderMap{
		orders: make(map[string]market.Order),
	}
}

func (o *OrderMap) Store(order market.Order) error {
	o.orders[order.GetID().String()] = order
	return nil
}