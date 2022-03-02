package entitity

import "implementacao/global"

type User struct {
	Id    int
	Name  string `validate:"required"`
	Email string `validate:"required,email"`
}

type Params struct {
	Id          int
	Name, Email string
}

func NewUser(params Params) (*User, error) {
	user := &User{Id: params.Id, Name: params.Name, Email: params.Email}

	if err := user.Validate(); err != nil {
		return nil, err
	}

	return user, nil
}

func (user *User) Validate() error {
	err := global.Validator().Validate(user)

	if err != nil {
		return err
	}

	return nil
}
