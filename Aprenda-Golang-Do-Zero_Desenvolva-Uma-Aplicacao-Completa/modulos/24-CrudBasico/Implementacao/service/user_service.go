package service

import (
	"implementacao/entitity"
	repositoryContracts "implementacao/repository/contracts"
	serviceContracts "implementacao/service/contracts"
)

type UserService struct {
	userRepository repositoryContracts.UserRepository
}

func NewUserService(userRepository repositoryContracts.UserRepository) serviceContracts.UserService {
	return &UserService{userRepository: userRepository}
}

func (service *UserService) Create(user *entitity.User) (int, error) {
	if err := user.Validate(); err != nil {
		return 0, err
	}
	return service.userRepository.Create(user)
}

func (service *UserService) List() ([]entitity.User, error) {
	return service.userRepository.List()
}

func (service *UserService) Get(id int) (*entitity.User, error) {
	return service.userRepository.Get(id)
}

func (service *UserService) Update(user *entitity.User) error {
	return service.userRepository.Update(user)
}

func (service *UserService) Delete(id int) error {
	return service.userRepository.Delete(id)
}
