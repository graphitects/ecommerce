package billing

import (
	"errors"
	"regexp"
)

type Card struct {
	// number is the card number.
	number string

	// holder is the card holder.
	holder string

	// expirationDate is the expiration date.
	expirationDate string

	// cvv is the cvv.
	cvv string
}

// CardNumber returns the card number.
func (c Card) CardNumber() string {
	return c.number
}

// Holder returns the card holder.
func (c Card) Holder() string {
	return c.holder
}

// ExpirationDate returns the expiration date.
func (c Card) ExpirationDate() string {
	return c.expirationDate
}

// CVV returns the cvv.
func (c Card) CVV() string {
	return c.cvv
}

// NewCard creates a new card.
func NewCard(number, holder, expirationDate, cvv string) (Card, error) {
	// validate the card number
	rx := regexp.MustCompile(`^\d{16}$`)
	if !rx.MatchString(number) {
		return Card{}, errors.New("invalid card number")
	}

	// validate the expiration date
	rx = regexp.MustCompile(`^\d{2}/\d{2}$`)
	if !rx.MatchString(expirationDate) {
		return Card{}, errors.New("invalid expiration date")
	}

	// validate the cvv
	rx = regexp.MustCompile(`^\d{3}$`)
	if !rx.MatchString(cvv) {
		return Card{}, errors.New("invalid cvv")
	}

	return Card{
		number:         number,
		holder:         holder,
		expirationDate: expirationDate,
		cvv:            cvv,
	}, nil
}
