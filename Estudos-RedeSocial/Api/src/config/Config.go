package config

import (
	"Api/src/core"
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"strconv"
)

const DbUsuario = "DB_USUARIO"
const DbSenha = "DB_SENHA"
const DbNomeBanco = "DB_NOME_BANCO"
const ApiPort = "API_PORT"

const portaPadrao = 9000

var (
	StringConexao = ""
	Porta         = 0
)

func Carregar() {
	carregarAmbiente()
	trataPortaApi()
	tratarStringConexao()
}

func carregarAmbiente(){
	var erro error
	erro = godotenv.Load()
	core.TratarErro(erro)
}

func trataPortaApi() {
	var erro error
	Porta, erro = strconv.Atoi(os.Getenv(ApiPort))
	if erro != nil {
		Porta = portaPadrao
	}
}

func tratarStringConexao() {
	StringConexao = fmt.Sprintf(
		"%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv(DbUsuario), os.Getenv(DbSenha), os.Getenv(DbNomeBanco))
}
