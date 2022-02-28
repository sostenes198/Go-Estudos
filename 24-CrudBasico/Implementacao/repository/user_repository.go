package repository

import "implementacao/entitity"

type UserRepository interface{
	Create(user *entitity.User) (int, error)
}
