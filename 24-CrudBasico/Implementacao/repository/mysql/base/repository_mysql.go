package base

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type repositoryMySql struct {
	db *sql.DB
}

func Open() (*repositoryMySql, error) {
	db, err := sql.Open("mysql", "root:root@/devbook?charset=utf8&parseTime=true&loc=Local")
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return &repositoryMySql{db: db}, nil
}

func (repository *repositoryMySql) Prepare(query string) (*sql.Stmt, error) {
	return repository.db.Prepare(query)
}

func (repository *repositoryMySql) Close() {
	err := repository.db.Close()
	if err != nil {
		log.Fatal(err)
		return
	}
}
