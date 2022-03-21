package validators

import (
	"fmt"
)

type StringValidator struct {
	Required bool
	Min      int
	Max      int
}

func (validator StringValidator) Validate(val interface{}) (bool, []error) {
	var errs []error
	isValid := true

	length := len(val.(string))

	if length == 0 && validator.Required {
		errs = append(errs, fmt.Errorf("cannot be null or empty"))
		isValid = false
	}

	if length < validator.Min {
		errs = append(errs, fmt.Errorf("should be at least %v chars long", validator.Min))
		isValid = false
	}

	if length > validator.Max {
		errs = append(errs, fmt.Errorf("should be less than %v chars long", validator.Max))
		isValid = false
	}

	return isValid, errs
}
