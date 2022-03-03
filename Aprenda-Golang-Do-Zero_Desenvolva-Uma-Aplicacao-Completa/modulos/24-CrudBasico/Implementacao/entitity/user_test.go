package entitity_test

import (
	"github.com/stretchr/testify/require"
	"implementacao/entitity"
	"testing"
)

func TestNewUser(t *testing.T) {
	t.Parallel()
	// assert

	// act
	user, err := entitity.NewUser(entitity.Params{Name: "Ss", Email: "ss@hotmail.com"})

	// assert
	require.NotNil(t, user)
	require.Nil(t, err)
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
			user, err := entitity.NewUser(entitity.Params{Name: scenario.value.(string), Email: "valid-email@gmail.com"})

			// assert
			require.Nil(t, user)
			require.NotNil(t, err)
			require.Equal(t, scenario.expectedMessage, err.Error())
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
			user, err := entitity.NewUser(entitity.Params{Name: "ValidName", Email: scenario.value.(string)})

			// assert
			require.Nil(t, user)
			require.NotNil(t, err)
			require.Equal(t, scenario.expectedMessage, err.Error())
		}
	})
}
