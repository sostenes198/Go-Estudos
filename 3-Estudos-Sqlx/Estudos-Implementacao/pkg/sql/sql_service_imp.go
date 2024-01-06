package pkg_sql

import (
	connectionsconfig "3-Estudos-Sqlx/Estudos-Implementacao/pkg/sql/connections-config"
	"errors"
	"github.com/jmoiron/sqlx"
)

type sqlServiceImp struct {
	db                     *sqlx.DB
	connectionStringConfig connectionsconfig.SqlConnectionConfig
}

func NewSqlService(connectionStringConfig connectionsconfig.SqlConnectionConfig) SqlService {
	return &sqlServiceImp{db: nil, connectionStringConfig: connectionStringConfig}
}

func NewSqlServiceWithDb(db *sqlx.DB, connectionStringConfig connectionsconfig.SqlConnectionConfig) SqlService {
	return &sqlServiceImp{db: db, connectionStringConfig: connectionStringConfig}
}

func (sqlService *sqlServiceImp) OpenConnection() error {
	var err error = nil
	connectionString, err := sqlService.connectionStringConfig.GetConnectionString()
	if err != nil {
		return err
	}

	databaseType := sqlService.connectionStringConfig.GetDatabaseType()

	err = openConnection(sqlService, databaseType, connectionString)
	if err != nil {
		return err
	}

	err = sqlService.db.Ping()
	if err != nil {
		return err
	}

	return nil
}

func (sqlService *sqlServiceImp) CloseConnection() error {
	if sqlService.db != nil {
		err := sqlService.db.Close()
		if err != nil {
			return err
		}

		return nil
	}

	return errors.New("Connection is not open")
}

func (sqlService *sqlServiceImp) GetDb() (*sqlx.DB, error) {
	if sqlService.db == nil {
		return nil, errors.New("The connection DB was not open")
	}

	return sqlService.db, nil
}

func openConnection(sqlService *sqlServiceImp, databaseType connectionsconfig.SqlDatabaseType, connectionString string) error {
	var err error = nil
	if sqlService.db == nil {
		sqlService.db, err = sqlx.Open(string(databaseType), connectionString)
		if err != nil {
			return err
		}
	}
	return nil
}
