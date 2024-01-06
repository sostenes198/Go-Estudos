//go:build unittest

package domain_person_test

import (
	domain_person "3-Estudos-Sqlx/Estudos-Implementacao/src/domain/person"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPerson(t *testing.T) {
	t.Parallel()

	t.Run("Create Person", func(t *testing.T) {
		t.Parallel()

		// arrange
		expectedPerson := domain_person.NewPerson("FirstName", "LastName", "email@email.com")

		// act
		person := domain_person.NewPerson("FirstName", "LastName", "email@email.com")

		// assert
		assert.Equal(t, person, expectedPerson)
	})

	t.Run("Should create empty Person", func(t *testing.T) {
		t.Parallel()

		expectedPerson := domain_person.NewPerson("", "", "")

		person := domain_person.EmptyPerson()

		assert.Equal(t, person, expectedPerson)
	})
}
