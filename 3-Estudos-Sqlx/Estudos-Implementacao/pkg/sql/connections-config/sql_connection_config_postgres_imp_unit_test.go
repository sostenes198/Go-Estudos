//go:build unittest

package pkg_sql_connections_config_test

import (
	pkg_sql_connections_config "3-Estudos-Sqlx/Estudos-Implementacao/pkg/sql/connections-config"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSqlConnectionConfigPostgresImp(t *testing.T) {
	t.Parallel()

	const (
		host         = "HOST"
		port         = 123
		user         = "USER"
		password     = "PASSWORD"
		databaseName = "DATABASE_NAME"
	)

	t.Run("Database Type", func(t *testing.T) {
		t.Parallel()

		t.Run("Should Get Database Type", func(t *testing.T) {
			t.Parallel()

			// arrange
			connectionConfig := pkg_sql_connections_config.NewSqlConnectionConfig(host, port, user, password, databaseName)

			// act
			result := connectionConfig.GetDatabaseType()

			// assert
			assert.Equal(t, string(result), "postgres")
		})

		t.Run("Get Connection String", func(t *testing.T) {
			t.Parallel()

			t.Run("Should get connection string", func(t *testing.T) {
				t.Parallel()

				// arrange
				connectionConfig := pkg_sql_connections_config.NewSqlConnectionConfig(host, port, user, password, databaseName)

				// act
				result, err := connectionConfig.GetConnectionString()

				// assert
				assert.Nil(t, err)
				assert.Equal(t, result, "host=HOST port=123 user=USER password=PASSWORD dbname=DATABASE_NAME sslmode=disable")
			})
		})
	})
}
