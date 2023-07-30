package models

import (
	"devbook/src/security"
	"errors"
	"github.com/badoux/checkmail"
	"strings"
	"time"
)

const (
	UserStepCreate = "Create"
	UserStepUpdate = "Update"
)

type User struct {
	Id       uint64
	Name     string
	Nick     string
	Email    string
	Password string
	CreateAt time.Time
}

type ParamsUser struct {
	Id       uint64
	Name     string
	Nick     string
	Email    string
	Password string
	CreateAt time.Time
}

// NewUser cria uma instância de usuário
func NewUser(params ParamsUser) *User {
	user := &User{
		Id:       params.Id,
		Name:     params.Name,
		Nick:     params.Nick,
		Email:    params.Email,
		Password: params.Password,
		CreateAt: params.CreateAt,
	}
	user.format()
	return user
}

// Validate valida se é um usuário válido
func (user *User) Validate(step string) []error {
	var errs []error
	if user.Name == "" {
		errs = append(errs, errors.New("Nome do usuário não pode ser nullo ou vazio."))
	}
	if user.Nick == "" {
		errs = append(errs, errors.New("Nick do usuário não pode ser nullo ou vazio."))
	}
	if user.Email == "" {
		errs = append(errs, errors.New("Email do usuário não pode ser nullo ou vazio."))
	}

	if err := checkmail.ValidateFormat(user.Email); err != nil {
		errs = append(errs, errors.New("Email do usuário é inválido."))
	}

	if step == UserStepCreate && user.Password == "" {
		errs = append(errs, errors.New("Password do usuário não pode ser nullo ou vazio."))
	}

	if len(errs) > 0 {
		return errs
	}

	return nil
}

func (user *User) format() {
	user.Name = strings.TrimSpace(user.Name)
	user.Nick = strings.TrimSpace(user.Nick)
	user.Email = strings.TrimSpace(user.Email)
}

func (user *User) CryptPassword() error {
	hashPassword, err := security.Hash(user.Password)
	if err != nil {
		return err
	}
	user.Password = string(hashPassword)

	return nil
}
