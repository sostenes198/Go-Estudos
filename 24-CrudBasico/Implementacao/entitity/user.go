package entitity

import "implementacao/global"

type User struct {
	Name  string `validate:"required"`
	Email string `validate:"required,email"`
}

func NewUser(name string, email string) (*User, error) {
	user := &User{Name: name, Email: email}
	err := global.Validator().Validate(user)

	if err != nil {
		return nil, err
	}

	return user, nil
}
