//go:build unittest

package pkg_sql

import (
	pkg_sql_connections_config "3-Estudos-Sqlx/Estudos-Implementacao/pkg/sql/connections-config"
	test_fixtures "3-Estudos-Sqlx/Estudos-Implementacao/test/fixtures"
	test_mocks "3-Estudos-Sqlx/Estudos-Implementacao/test/mocks"
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSqlService(t *testing.T) {
	t.Parallel()

	const connectionString string = "UNIT_TEST_CONNECTION_STRING"

	t.Run("Open Connection", func(t *testing.T) {
		t.Parallel()

		t.Run("Should Open Connection", func(t *testing.T) {
			t.Parallel()

			// arrange
			_, dbSqlX, _ := test_fixtures.CreateSqlXDbMock(t)

			connectionStringConfigMock := new(test_mocks.ConnectionStringConfigMock)
			connectionStringConfigMock.MockGetConnectionString(connectionString, nil)
			connectionStringConfigMock.MockGetDatabaseType(test_fixtures.MockDriverType)

			sqlService := NewSqlServiceWithDb(dbSqlX, connectionStringConfigMock)

			// act
			err := sqlService.OpenConnection()

			// assert
			assert.Nil(t, err)
			_assertConnectionStringConfig(connectionStringConfigMock, t, 1, 1)
		})

		t.Run("Should failed to open connection when failed to get connection string", func(t *testing.T) {
			t.Parallel()

			// arrange
			expectedErr := errors.New("UnitTestError")
			_, dbSqlX, _ := test_fixtures.CreateSqlXDbMock(t)

			connectionStringConfigMock := new(test_mocks.ConnectionStringConfigMock)
			connectionStringConfigMock.MockGetConnectionString("", expectedErr)

			sqlService := NewSqlServiceWithDb(dbSqlX, connectionStringConfigMock)

			// act
			err := sqlService.OpenConnection()

			// assert
			assert.ErrorIs(t, err, expectedErr)
			_assertConnectionStringConfig(connectionStringConfigMock, t, 1, 0)
		})

		t.Run("Should failed to open connection when failed to open sqlx connection", func(t *testing.T) {
			t.Parallel()

			// arrange
			connectionStringConfigMock := new(test_mocks.ConnectionStringConfigMock)
			connectionStringConfigMock.MockGetConnectionString(connectionString, nil)
			connectionStringConfigMock.MockGetDatabaseType(pkg_sql_connections_config.TypeMySql)

			sqlService := NewSqlService(connectionStringConfigMock)

			// act
			err := sqlService.OpenConnection()

			// assert
			assert.EqualError(t, err, "sql: unknown driver \"mysql\" (forgotten import?)")
			_assertConnectionStringConfig(connectionStringConfigMock, t, 1, 1)
		})

		t.Run("Should failed to open connection when failed to ping connection", func(t *testing.T) {
			t.Parallel()

			// arrange
			connectionStringConfigMock := new(test_mocks.ConnectionStringConfigMock)
			connectionStringConfigMock.MockGetConnectionString(connectionString, nil)
			connectionStringConfigMock.MockGetDatabaseType(test_fixtures.MockDriverType)

			sqlService := NewSqlService(connectionStringConfigMock)

			// act
			err := sqlService.OpenConnection()

			// assert
			assert.EqualError(t, err, "expected a connection to be available, but it is not")
			_assertConnectionStringConfig(connectionStringConfigMock, t, 1, 1)
		})
	})

	t.Run("Close Connection", func(t *testing.T) {
		t.Parallel()

		t.Run("Should close connection", func(t *testing.T) {
			t.Parallel()

			// arrange
			_, dbSqlX, mock := test_fixtures.CreateSqlXDbMock(t)

			mock.ExpectClose()

			connectionStringConfigMock := new(test_mocks.ConnectionStringConfigMock)

			sqlService := NewSqlServiceWithDb(dbSqlX, connectionStringConfigMock)

			// act
			errCloseConnection := sqlService.CloseConnection()

			// assert
			assert.Nil(t, errCloseConnection)
		})

		t.Run("Should return error when connection was not opened", func(t *testing.T) {
			t.Parallel()

			// arrange
			sqlService := sqlServiceImp{}

			// act
			err := sqlService.CloseConnection()

			// assert
			assert.EqualError(t, err, "Connection is not open")
		})

		t.Run("Should return error when failed to close connection", func(t *testing.T) {
			// arrange
			_, dbSqlX, _ := test_fixtures.CreateSqlXDbMock(t)

			connectionStringConfigMock := new(test_mocks.ConnectionStringConfigMock)

			sqlService := NewSqlServiceWithDb(dbSqlX, connectionStringConfigMock)

			// act
			err := sqlService.CloseConnection()

			// assert
			assert.EqualError(t, err, "all expectations were already fulfilled, call to database Close was not expected")
		})
	})

	t.Run("Get DB", func(t *testing.T) {
		t.Parallel()

		t.Run("Should get DB", func(t *testing.T) {
			t.Parallel()

			// arrange
			_, dbSqlX, _ := test_fixtures.CreateSqlXDbMock(t)

			connectionStringConfigMock := new(test_mocks.ConnectionStringConfigMock)

			sqlService := NewSqlServiceWithDb(dbSqlX, connectionStringConfigMock)

			// act
			db, err := sqlService.GetDb()

			// assert
			assert.Nil(t, err)
			assert.Equal(t, db, dbSqlX)
		})

		t.Run("Should return error when connection was not opened", func(t *testing.T) {
			t.Parallel()

			// arrange
			sqlService := sqlServiceImp{}

			// act
			db, err := sqlService.GetDb()

			// assert
			assert.EqualError(t, err, "The connection DB was not open")
			assert.Nil(t, db)
		})
	})
}

func _assertConnectionStringConfig(connectionStringConfigMock *test_mocks.ConnectionStringConfigMock, t *testing.T,
	expectedCallsGetConnectionString int, expectedCallsGetDatabaseType int) {
	connectionStringConfigMock.AssertExpectations(t)
	connectionStringConfigMock.AssertGetConnectionString(t, expectedCallsGetConnectionString)
	connectionStringConfigMock.AssertGetDatabaseType(t, expectedCallsGetDatabaseType)
}
