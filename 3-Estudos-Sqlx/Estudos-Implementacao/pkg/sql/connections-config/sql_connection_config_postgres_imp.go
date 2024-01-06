package pkg_sql_connections_config

import "fmt"

type sqlConnectionConfigPostgres struct {
	host         string
	port         int
	user         string
	password     string
	databaseName string
}

func NewSqlConnectionConfig(host string, port int, user string, password string, databaseName string) SqlConnectionConfig {
	return sqlConnectionConfigPostgres{host: host, port: port, user: user, password: password, databaseName: databaseName}
}

func (sql sqlConnectionConfigPostgres) GetDatabaseType() SqlDatabaseType {
	return TypePostgres
}

func (sql sqlConnectionConfigPostgres) GetConnectionString() (string, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", sql.host, sql.port, sql.user, sql.password, sql.databaseName)
	return psqlInfo, nil
}
