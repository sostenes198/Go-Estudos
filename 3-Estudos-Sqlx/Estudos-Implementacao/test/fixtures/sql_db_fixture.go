package test_fixtures

import (
	pkg_sql_connections_config "3-Estudos-Sqlx/Estudos-Implementacao/pkg/sql/connections-config"
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"testing"
)

const MockDriverType = pkg_sql_connections_config.TypeSqlMock

func CreateSqlXDbMock(t *testing.T) (*sql.DB, *sqlx.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err.Error())
	}

	sqlxDb := sqlx.NewDb(db, string(MockDriverType))

	return db, sqlxDb, mock
}
