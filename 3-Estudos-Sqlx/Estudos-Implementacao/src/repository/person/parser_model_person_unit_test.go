//go:build unittest

package repository_person_test

import (
	domain_person "3-Estudos-Sqlx/Estudos-Implementacao/src/domain/person"
	repository_person "3-Estudos-Sqlx/Estudos-Implementacao/src/repository/person"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParserModelPerson(t *testing.T) {
	t.Parallel()

	parser := repository_person.NewParserModelPerson()

	t.Run("Should parse Model to Entity", func(t *testing.T) {
		t.Parallel()

		// arrange
		model := repository_person.NewPersonModel("Name", "LastName", "e@e")
		expectedResult := domain_person.NewPerson("Name", "LastName", "e@e")

		// act
		result := parser.ParseToEntity(model)

		// assert
		assert.Equal(t, result, expectedResult)
	})

	t.Run("Should parse Entity to Model", func(t *testing.T) {
		t.Parallel()

		// arrange
		entity := domain_person.NewPerson("Name", "LastName", "e@e")
		expectedResult := repository_person.NewPersonModel("Name", "LastName", "e@e")

		// act
		result := parser.ParseToModel(entity)

		// assert
		assert.Equal(t, result, expectedResult)
	})
}
