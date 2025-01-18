package valueobjects

import (
	"errors"
)

type Password struct {
	// value is the password value.
	value string
}

// NewPassword creates a new password.
func NewPassword(value string) (Password, error) {
	// Must be between 8 and 255 characters.
	if len(value) < 8 || len(value) > 255 {
		return Password{}, errors.New("invalid password")
	}

	return Password{value: value}, nil
}

// Value returns the password value.
func (p Password) Value() string {
	return p.value
}
