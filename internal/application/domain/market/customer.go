package market

import "github.com/google/uuid"

type Customer struct {
	// id is the unique identifier of the customer.
	id uuid.UUID

	// userID is the unique identifier of the user.
	userID string

	// name is the name of the customer.
	username string
}

func (c *Customer) GetID() uuid.UUID {
	return c.id
}

func (c *Customer) GetUserID() string {
	return c.userID
}

func (c *Customer) GetUsername() string {
	return c.username
}

func NewCustomer(id uuid.UUID, userID, username string) Customer {
	return Customer{id: id, userID: userID, username: username}
}
