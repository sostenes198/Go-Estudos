package pkg_sql_connections_config

type SqlConnectionConfig interface {
	GetDatabaseType() SqlDatabaseType
	GetConnectionString() (string, error)
}

type SqlDatabaseType string

const (
	MySql    SqlDatabaseType = "mysql"
	Postgres SqlDatabaseType = "postgres"
)
