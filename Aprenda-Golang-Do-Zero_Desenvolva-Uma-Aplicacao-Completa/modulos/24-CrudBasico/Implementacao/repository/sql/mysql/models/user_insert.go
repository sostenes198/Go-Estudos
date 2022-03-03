package models

import "implementacao/entitity"

type userInsert struct {
	Name  string
	Email string
}

func NewUserInsert(user entitity.User) *userInsert {
	return &userInsert{
		user.Name,
		user.Email,
	}
}
