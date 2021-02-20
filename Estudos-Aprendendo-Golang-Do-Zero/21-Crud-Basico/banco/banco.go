package banco

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func Conectar()(db *sql.DB, erro error){
	stringConexao := "golang:golang@/devbook?charset=utf8&parseTime=True&loc=Local"

	db, erro = sql.Open("mysql", stringConexao)
	if erro != nil{
		return nil, erro
	}
	if erro = db.Ping(); erro != nil{
		return nil, erro
	}
	return db, nil
}
