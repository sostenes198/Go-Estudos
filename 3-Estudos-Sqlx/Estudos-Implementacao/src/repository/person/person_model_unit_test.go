//go:build unittest

package repository_person_test

import (
	repository_person "3-Estudos-Sqlx/Estudos-Implementacao/src/repository/person"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPersonModel(t *testing.T) {
	t.Parallel()

	t.Run("Should create PersonModel", func(t *testing.T) {
		t.Parallel()

		// arrange
		expectedResult := repository_person.NewPersonModel("F", "L", "e@e")

		// act
		result := repository_person.NewPersonModel("F", "L", "e@e")

		// assert
		assert.Equal(t, result, expectedResult)
	})
}
