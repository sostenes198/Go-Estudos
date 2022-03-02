package contracts

import "implementacao/entitity"

type UserRepository interface {
	Create(user *entitity.User) (int, error)
	List() ([]entitity.User, error)
	Get(id int) (*entitity.User, error)
	Update(user *entitity.User) error
	Delete(id int) error
}
