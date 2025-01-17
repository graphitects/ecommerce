package auth

import "github.com/graphictects/ecommerce/internal/shared/valueobjects"

type User struct {
	// id is the unique identifier of the user.
	id string

	// name is the name of the user.
	name string

	// email is the email of the user.
	email valueobjects.Email

	// password is the password of the user.
	password valueobjects.Password
}

func (u *User) GetID() string {
	return u.id
}

func (u *User) GetName() string {
	return u.name
}

func (u *User) GetEmail() valueobjects.Email {
	return u.email
}

func (u *User) GetPassword() valueobjects.Password {
	return u.password
}

func NewUser(id, name string, email valueobjects.Email, password valueobjects.Password) User {
	return User{id: id, name: name, email: email, password: password}
}
