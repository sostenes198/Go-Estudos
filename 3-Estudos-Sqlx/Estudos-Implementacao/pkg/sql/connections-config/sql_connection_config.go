package pkg_sql_connections_config

type SqlConnectionConfig interface {
	GetDatabaseType() SqlDatabaseType
	GetConnectionString() (string, error)
}

type SqlDatabaseType string

const (
	TypeMySql    SqlDatabaseType = "mysql"
	TypePostgres SqlDatabaseType = "postgres"
	TypeSqlMock  SqlDatabaseType = "sqlmock"
)
