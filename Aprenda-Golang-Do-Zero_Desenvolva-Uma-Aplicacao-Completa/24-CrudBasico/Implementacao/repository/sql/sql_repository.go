package sql

import (
	"database/sql"
	"log"
	"sync"
)

var once = sync.Once{}
var db *sql.DB

func Open() *sql.DB {
	once.Do(func() {
		dbInternal, err := sql.Open("mysql", "root:root@/devbook?charset=utf8&parseTime=true&loc=Local")
		if err != nil {
			log.Fatalln(err)
		}

		if err = dbInternal.Ping(); err != nil {
			log.Fatalln(err)
		}

		db = dbInternal
	})

	return db
}

func Close() error {
	err := db.Close()
	if err != nil {
		return err
	}

	once = sync.Once{}
	db = nil

	return nil
}
