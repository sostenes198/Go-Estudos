package service_test

import (
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"implementacao/entitity"
	service2 "implementacao/service"
	"testing"
)

type mockUserRepository struct {
	mock.Mock
}

func (m *mockUserRepository) Create(user *entitity.User) (int, error) {
	args := m.Called(user)
	return args.Int(0), args.Error(1)
}

func (m *mockUserRepository) List() ([]entitity.User, error) {
	args := m.Called()
	return args.Get(0).([]entitity.User), args.Error(1)
}

func (m *mockUserRepository) Get(id int) (*entitity.User, error) {
	args := m.Called(id)
	return args.Get(0).(*entitity.User), args.Error(1)
}

func (m *mockUserRepository) Update(user *entitity.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *mockUserRepository) Delete(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestCreate(t *testing.T) {
	t.Parallel()

	t.Run("Create User", func(t *testing.T) {
		t.Parallel()

		// arrange
		user, userErr := entitity.NewUser(entitity.Params{Name: "Test", Email: "Testing@test.com"})
		expectedInsertResult := 5

		mockUserRepository := new(mockUserRepository)
		mockUserRepository.On("Create", user).Return(expectedInsertResult, nil)

		service := service2.NewUserService(mockUserRepository)

		// act
		insertResult, createErr := service.Create(user)

		// assert
		require.NoError(t, userErr)
		require.NoError(t, createErr)
		require.Equal(t, expectedInsertResult, insertResult)

		mockUserRepository.AssertExpectations(t)
		mockUserRepository.AssertNumberOfCalls(t, "Create", 1)
	})

	t.Run("Fail when invalid user", func(t *testing.T) {
		t.Parallel()

		// arrange
		user := entitity.User{}

		mockUserRepository := new(mockUserRepository)

		service := service2.NewUserService(mockUserRepository)

		// act
		insertResult, createErr := service.Create(&user)

		// assert
		require.Equal(t, 0, insertResult)
		require.EqualError(t, createErr.(validator.ValidationErrors), "Key: 'User.Name' Error:Field validation for 'Name' failed on the 'required' tag\nKey: 'User.Email' Error:Field validation for 'Email' failed on the 'required' tag")
	})

}

func TestList(t *testing.T) {
	t.Parallel()

	t.Run("List Users", func(t *testing.T) {
		t.Parallel()

		// arrange
		expectedUsers := []entitity.User{
			{Id: 0, Name: "Teste1", Email: "1@gmail.com"},
			{Id: 1, Name: "Teste2", Email: "2@gmail.com"},
			{Id: 2, Name: "Teste3", Email: "3@gmail.com"},
		}

		mockUserRepository := new(mockUserRepository)
		mockUserRepository.On("List").Return(expectedUsers, nil)

		service := service2.NewUserService(mockUserRepository)

		// act
		users, err := service.List()

		// assert
		require.NoError(t, err)
		require.Equal(t, expectedUsers, users)

		mockUserRepository.AssertExpectations(t)
		mockUserRepository.AssertNumberOfCalls(t, "List", 1)
	})
}

func TestGetById(t *testing.T) {
	t.Parallel()

	t.Run("Get User by id", func(t *testing.T) {
		t.Parallel()

		// arrange
		id := 55
		expectedUser := &entitity.User{Id: id, Name: "Testll", Email: "Testll@g.com"}

		mockUserRepository := new(mockUserRepository)
		mockUserRepository.On("Get", id).Return(expectedUser, nil)

		service := service2.NewUserService(mockUserRepository)

		// act
		user, err := service.Get(id)

		// assert
		require.Nil(t, err)
		require.Equal(t, expectedUser, user)
	})

}
