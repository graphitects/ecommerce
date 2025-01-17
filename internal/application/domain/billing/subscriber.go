package billing

type subscription struct {
	// Role is the role of the subscription.
	Role string

	// Price is the price of the subscription.
	Price float64

	// Discount is an offered discount.
	Discount float64
}

var (
	// FreeSubscription is a free subscription.
	FreeSubscription = subscription{
		Role:     "free",
		Price:    0,
		Discount: 0,
	}

	// PremiumSubscription is a premium subscription.
	PremiumSubscription = subscription{
		Role:     "premium",
		Price:    10,
		Discount: 0.1,
	}

	// GoldSubscription is a gold subscription.
	GoldSubscription = subscription{
		Role:     "gold",
		Price:    30,
		Discount: 0.45,
	}
)

func NewSubscription(role string) subscription {
	switch role {
	case "premium":
		return PremiumSubscription
	case "gold":
		return GoldSubscription
	default:
		return FreeSubscription
	}
}

type Subscriber struct {
	// id is the unique identifier of the subscriber
	id string

	// userID is the unique identifier of the user.
	userID string

	// subscription is the subscription of the user.
	subscription subscription

	// coins is the coins of the user.
	coins uint64
}

func (s *Subscriber) GetID() string {
	return s.id
}

func (s *Subscriber) GetUserID() string {
	return s.userID
}

func (s *Subscriber) GetSubscription() subscription {
	return s.subscription
}

func (s *Subscriber) GetCoins() uint64 {
	return s.coins
}

func NewSubscriber(id, userID string, subscription subscription, coins uint64) Subscriber {
	return Subscriber{id: id, userID: userID, subscription: subscription, coins: coins}
}
