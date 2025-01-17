package market

type ProductRepository interface {
	Adquire(id string, quantity uint64) (*Product, error)
	Release(id string, quantity uint64) error
}
