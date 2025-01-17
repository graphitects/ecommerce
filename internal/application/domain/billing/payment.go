package billing

type UserPayment interface {
	// AdquireCoins adquires coins to the user.
	AdquireCoins(userId string, coins uint64) error
	// RestoreCoins restores coins to the user.
	RestoreCoins(userId string, coins uint64) error
}

type CardPayment interface {
	// PayWithCard pays with a card.
	Pay(card Card, amount float64) error
	// Refund refunds a payment.
	Refund(card Card, amount float64) error
}
