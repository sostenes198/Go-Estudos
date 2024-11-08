package validators

type DefaultValidator struct {
}

func (validator DefaultValidator) Validate(interface{}) (bool, []error) {
	var errors []error
	return true, errors
}
