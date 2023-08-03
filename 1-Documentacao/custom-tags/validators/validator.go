package validators

import (
	"fmt"
	"reflect"
	"strings"
)

const (
	tagName = "validate"
)

type Validator interface {
	Validate(val interface{}) (bool, []error)
}

func ValidateStruct(val interface{}) (bool, []error) {
	var errs []error
	isValid := true

	reflectValues := reflect.ValueOf(val)

	for i := 0; i < reflectValues.NumField(); i++ {
		tag := reflectValues.Type().Field(i).Tag.Get(tagName)

		if tag == "" || tag == "-" {
			continue
		}

		validator := getValidatorFromTag(tag)

		validatorIsValid, validatorErrs := validator.Validate(reflectValues.Field(i).Interface())

		if !validatorIsValid {
			for _, err := range validatorErrs {
				errs = append(errs, fmt.Errorf("%s %s", reflectValues.Type().Field(i).Name, err.Error()))
			}
			isValid = false
		}
	}

	return isValid, errs
}

func getValidatorFromTag(tag string) Validator {
	args := strings.Split(tag, ",")

	switch args[0] {
	case "string":
		validator := StringValidator{}
		fmt.Sscanf(strings.Join(args[1:], ","), "required=%t,min=%d,max=%d", &validator.Required, &validator.Min, &validator.Max)
		return validator
	}

	return DefaultValidator{}
}
