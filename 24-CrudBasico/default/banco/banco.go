package banco

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql" // Driver de conexão com o MySql
)

type Repositorio interface {
	Conectar() (*sql.DB, error)
}

type RepositorioDb struct {
}

func NewRepositorioDb() *RepositorioDb {
	return &RepositorioDb{}
}

func (esDb RepositorioDb) Conectar() (*sql.DB, error) {
	stringConexao := "root:root@/devbook?charset=utf8&parseTime=true&loc=Local"

	db, erro := sql.Open("mysql", stringConexao)
	if erro != nil {
		return nil, erro
	}

	if erro = db.Ping(); erro != nil {
		return nil, erro
	}

	return db, nil
}
