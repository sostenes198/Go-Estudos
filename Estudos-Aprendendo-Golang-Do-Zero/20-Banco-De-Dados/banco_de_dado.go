package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main(){
	stringConexao := "golang:golang@/devbook?charset=utf8&parseTime=True&loc=Local"
	db, erro := sql.Open("mysql", stringConexao)
	defer db.Close()
	printErro(erro)
	printErro(db.Ping())

	linhas, erro := db.Query("select * from usuarios")
	printErro(erro)
	defer linhas.Close()

	fmt.Println(linhas)
}

func printErro(erro error){
	if erro != nil{
		log.Fatal(erro)
	}
}
