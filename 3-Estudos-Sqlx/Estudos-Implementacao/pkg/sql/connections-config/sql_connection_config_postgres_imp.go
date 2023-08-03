package pkg_sql_connections_config

import "fmt"

type SqlConnectionConfigPostgres struct {
	host         string
	port         int
	user         string
	password     string
	databaseName string
}

func NewSqlConnectionConfig(host string, port int, user string, password string, databaseName string) SqlConnectionConfig {
	return SqlConnectionConfigPostgres{host: host, port: port, user: user, password: password, databaseName: databaseName}
}

func (sql SqlConnectionConfigPostgres) GetDatabaseType() SqlDatabaseType {
	return Postgres
}

func (sql SqlConnectionConfigPostgres) GetConnectionString() (string, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", sql.host, sql.port, sql.user, sql.password, sql.databaseName)
	return psqlInfo, nil
}
