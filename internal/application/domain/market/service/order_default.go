package service

import (
	"github.com/google/uuid"
	"github.com/graphictects/ecommerce/internal/application/domain/billing"
	"github.com/graphictects/ecommerce/internal/application/domain/market"
)

type OrderDefault struct {
	productRepository market.ProductRepository
	orderRepository market.OrderRepository

	paymentService billing.CardPayment
}

func NewOrderDefault(productRepository market.ProductRepository, orderRepository market.OrderRepository, paymentService billing.PaymentService) *OrderDefault {
	return &OrderDefault{
		productRepository: productRepository,
		orderRepository: orderRepository,
		paymentService: paymentService,
	}
}

func (o *OrderDefault) CreateOrder(customerID string, products map[market.Product]uint64, cardInfo market.CardInfo) (newOrder *market.Order, err error) {
	// prepare the order
	order := market.NewOrder(uuid.New(), customerID, make(map[market.Product]uint64))

	// prepare the card
	card, err := billing.NewCard(cardInfo.Number, cardInfo.Holder, cardInfo.Exp, cardInfo.CVV)
	if err != nil {
		return nil, err
	}

	// verify products
	for product, quantity := range products {
		if err = o.validAdquisition(product, quantity); err != nil {
			return nil, err
		}

		order.AddProduct(product, quantity)
	}

	// proceed with the payment
	if err = o.paymentService.Pay(card, order.Total()); err != nil {	
		return nil, err
	}
	
	// save the order
	if err = o.orderRepository.Store(order); err != nil {
		// log error
	}

	return order, nil
}

func (o *OrderDefault) validAdquisition(product market.Product, quantity uint64) (err error) {
	defer func() {
		if err != nil {
			o.productRepository.Restore(product.GetID().String(), quantity)
		}
	}()

	// adquire ownership of n number of this product
	validProduct, err := o.productRepository.Adquire(product.GetID().String(), quantity)
	if err != nil {
		return err
	}

	// check if the product is the same as the one requested
	if validProduct != product {
		return market.ErrProductVersionMismatch
	}

	return nil
}