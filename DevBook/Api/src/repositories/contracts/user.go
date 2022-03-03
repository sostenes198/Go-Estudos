package contracts

import "devbook/src/models"

type UserListFilter struct {
	NameOrNick string
}

type UserRepository interface {
	List(filter UserListFilter) ([]models.User, error)
	GetById(id uint64) (*models.User, error)
	Create(user *models.User) error
}
