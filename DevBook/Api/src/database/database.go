package database

import (
	"database/sql"
	"devbook/src/config"
	_ "github.com/go-sql-driver/mysql"
	"sync"
)

var once sync.Once
var Db *sql.DB

// Open abre conexão com o banco de dados
func Open() error {
	var errClojure error
	once.Do(func() {
		db, err := sql.Open("mysql", config.ConnectionStringDb)
		if err != nil {
			errClojure = err
		}

		if err = db.Ping(); err != nil {
			db.Close()
			errClojure = err
		}

		Db = db
	})

	return errClojure
}

// Close fecha conexão com o banco de dados
func Close() {
	if Db != nil {
		Db.Close()
	}
}
