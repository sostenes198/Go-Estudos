package test_mocks

import (
	pkg_sql_connections_config "3-Estudos-Sqlx/Estudos-Implementacao/pkg/sql/connections-config"
	"github.com/stretchr/testify/mock"
	"testing"
)

type ConnectionStringConfigMock struct {
	mock.Mock
}

const (
	_nameGetDatabaseType     = "GetDatabaseType"
	_nameGetConnectionString = "GetConnectionString"
)

func (c *ConnectionStringConfigMock) GetDatabaseType() pkg_sql_connections_config.SqlDatabaseType {
	args := c.Called()
	return args.Get(0).(pkg_sql_connections_config.SqlDatabaseType)
}

func (c *ConnectionStringConfigMock) GetConnectionString() (string, error) {
	args := c.Called()
	return args.String(0), args.Error(1)
}

func (c *ConnectionStringConfigMock) MockGetConnectionString(connectionString string, err error) {
	c.On(_nameGetConnectionString, mock.Anything).Return(connectionString, err)
}

func (c *ConnectionStringConfigMock) MockGetDatabaseType(sqlDatabaseType pkg_sql_connections_config.SqlDatabaseType) {
	c.On(_nameGetDatabaseType, mock.Anything).Return(sqlDatabaseType)
}

func (c *ConnectionStringConfigMock) AssertGetConnectionString(t *testing.T, expectedCalls int) {
	if expectedCalls > 0 {
		c.AssertCalled(t, _nameGetConnectionString, mock.Anything)
	}
	c.AssertNumberOfCalls(t, _nameGetConnectionString, expectedCalls)
}

func (c *ConnectionStringConfigMock) AssertGetDatabaseType(t *testing.T, expectedCalls int) {
	if expectedCalls > 0 {
		c.AssertCalled(t, _nameGetDatabaseType, mock.Anything)
	}
	c.AssertNumberOfCalls(t, _nameGetDatabaseType, expectedCalls)
}
