package query_test

import (
	"github.com/stretchr/testify/require"
	. "implementacao/repository/sql/mysql/query"
	"testing"
)

func TestQueries(t *testing.T) {
	t.Parallel()

	t.Run("Query Create User", func(t *testing.T) {
		t.Parallel()

		// arrange
		expectedQuery := "INSERT INTO usuario (nome, email) VALUES (?, ?)"

		// act - assert
		require.Equal(t, expectedQuery, Create)
	})

	t.Run("Query List Users", func(t *testing.T) {
		t.Parallel()

		// arrange
		expectedQuery := "SELECT id, nome, email FROM usuario"

		// act - assert
		require.Equal(t, expectedQuery, List)
	})

	t.Run("Query Get User by id", func(t *testing.T) {
		t.Parallel()

		// arrange
		expectedQuery := "SELECT id, nome, email FROM usuario WHERE id=?"

		// act - assert
		require.Equal(t, expectedQuery, GetById)
	})
}
