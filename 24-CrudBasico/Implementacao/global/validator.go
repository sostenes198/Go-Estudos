package global

import (
	"github.com/go-playground/validator/v10"
	"sync"
)

var once = sync.Once{}
var interValidator *ValidatorGlobal

type ValidatorGlobal struct {
	validator *validator.Validate
}

func (v *ValidatorGlobal) Validate(input interface{}) error {
	return v.validator.Struct(input)
}

func Validator() *ValidatorGlobal {
	once.Do(func() {
		interValidator = &ValidatorGlobal{validator: validator.New()}
	})
	return interValidator
}
