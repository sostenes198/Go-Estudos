package entitity_test

import (
	"github.com/stretchr/testify/assert"
	. "implementacao/entitity"
	"testing"
)

func TestNewUser(t *testing.T) {
	t.Parallel()
	// assert

	// act
	user, err := NewUser("Ss", "ss@hotmail.com")

	// assert
	assert.NotNil(t, user)
	assert.Nil(t, err)
}

func TestReturnErrorWhenUserNotValid(t *testing.T) {
	type errorValidationScenario struct {
		value           interface{}
		expectedMessage string
	}

	t.Parallel()

	t.Run("Name", func(t *testing.T) {
		t.Parallel()
		// arrange
		scenarios := []errorValidationScenario{
			{"", "Key: 'User.Name' Error:Field validation for 'Name' failed on the 'required' tag"},
		}

		// act
		for _, scenario := range scenarios {
			user, err := NewUser(scenario.value.(string), "valid-email@gmail.com")

			// assert
			assert.Nil(t, user)
			if assert.NotNil(t, err) {
				assert.Equal(t, scenario.expectedMessage, err.Error())
			}

		}
	})

	t.Run("Email", func(t *testing.T) {
		t.Parallel()

		// arrange
		scenarios := []errorValidationScenario{
			{"", "Key: 'User.Email' Error:Field validation for 'Email' failed on the 'required' tag"},
			{"invalidemail", "Key: 'User.Email' Error:Field validation for 'Email' failed on the 'email' tag"},
		}

		// act
		for _, scenario := range scenarios {
			user, err := NewUser("ValidName", scenario.value.(string))

			// assert
			assert.Nil(t, user)
			if assert.NotNil(t, err) {
				assert.Equal(t, scenario.expectedMessage, err.Error())
			}
		}
	})
}
