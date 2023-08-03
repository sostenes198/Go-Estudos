package pkg_sql

import (
	connectionsconfig "3-Estudos-Sqlx/Estudos-Implementacao/pkg/sql/connections-config"
	"errors"
	"github.com/jmoiron/sqlx"
	"sync"
)

var once sync.Once
var db *sqlx.DB

type SqlService struct {
	db                     *sqlx.DB
	connectionStringConfig connectionsconfig.SqlConnectionConfig
}

func NewSqlService(connectionStringConfig connectionsconfig.SqlConnectionConfig) SqlService {
	return SqlService{db: db, connectionStringConfig: connectionStringConfig}
}

func (sqlService *SqlService) OpenConnection() error {
	var err error = nil
	connectionString, err := sqlService.connectionStringConfig.GetConnectionString()
	if err != nil {
		return err
	}

	databaseType := sqlService.connectionStringConfig.GetDatabaseType()

	once.Do(func() {
		db, err = sqlx.Open(string(databaseType), connectionString)
		if err != nil {
			return
		}

		err = db.Ping()
		if err != nil {
			return
		}
	})

	return err
}

func (sqlService *SqlService) CloseConnection() error {
	if db != nil {
		err := db.Close()
		return err
	}

	return nil
}

func (sqlService *SqlService) GetDb() (*sqlx.DB, error) {
	if db == nil {
		return nil, errors.New("A conexão com o BD não foi aberta ainda")
	}

	return db, nil
}
