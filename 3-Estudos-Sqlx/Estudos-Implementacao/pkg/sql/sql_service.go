package pkg_sql

import "github.com/jmoiron/sqlx"

type SqlService interface {
	OpenConnection() error
	CloseConnection() error
	GetDb() (*sqlx.DB, error)
}
