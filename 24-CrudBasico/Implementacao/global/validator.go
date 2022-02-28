package global

import (
	"github.com/go-playground/validator/v10"
	"sync"
)

type internalValidator struct {
}

func (v *internalValidator) Validate(input interface{}) error {
	return playgroundValidator.Struct(input)
}

var once = sync.Once{}
var playgroundValidator *validator.Validate
var interValidator *internalValidator

func Validator() *internalValidator {
	once.Do(func() {
		interValidator = &internalValidator{}
		playgroundValidator = validator.New()
	})
	return interValidator
}
