//go:build unittest

package pkg_sql_connections_config_test

import (
	pkg_sql_connections_config "3-Estudos-Sqlx/Estudos-Implementacao/pkg/sql/connections-config"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSqlConnectionConfig(t *testing.T) {
	t.Parallel()

	t.Run("Should validate SqlDatabaseType", func(t *testing.T) {
		t.Parallel()

		// arrange

		// act
		mySql := pkg_sql_connections_config.TypeMySql
		postgres := pkg_sql_connections_config.TypePostgres
		sqlMock := pkg_sql_connections_config.TypeSqlMock

		//assert
		assert.Equal(t, string(mySql), "mysql")
		assert.Equal(t, string(postgres), "postgres")
		assert.Equal(t, string(sqlMock), "sqlmock")
	})
}
