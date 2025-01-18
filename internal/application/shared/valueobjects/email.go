package valueobjects

import (
	"errors"
	"regexp"
)

type Email struct {
	// value is the email value.
	value string
}

// NewEmail creates a new email.
func NewEmail(value string) (Email, error) {
	rx := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !rx.MatchString(value) {
		return Email{}, errors.New("invalid email")
	}

	return Email{value: value}, nil
}

// Value returns the email value.
func (e Email) Value() string {
	return e.value
}
