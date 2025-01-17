package market

type OrderRepository interface {
	Store(order *Order) error
}
