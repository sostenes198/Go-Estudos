package servidor

import (
	"crud/banco"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type usuario struct {
	ID    uint32 `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
}

var erro error

func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	corpoRequisicao, erro := ioutil.ReadAll(r.Body)
	if escreverErroSeExistir(w, erro, "Falha ao ler corpo da requisição") {
		return
	}

	banco := banco.NewRepositorioDb()

	var usuario usuario

	erro = json.Unmarshal(corpoRequisicao, &usuario)

	if escreverErroSeExistir(w, erro, "Erro ao converter body da requisição para um usuário") {
		return
	}

	db, erro := banco.Conectar()
	if escreverErroSeExistir(w, erro, "Erro ao conectar banco de dados") {
		return
	}
	defer db.Close()

	statement, erro := db.Prepare("insert into usuario (nome, email) values (?, ?)")
	if escreverErroSeExistir(w, erro, "Erro ao criar statement") {
		return
	}
	defer statement.Close()

	insercao, erro := statement.Exec(usuario.Nome, usuario.Email)
	if escreverErroSeExistir(w, erro, "Erro ao inserir no banco de dados.") {
		return
	}

	idInserido, erro := insercao.LastInsertId()
	if escreverErroSeExistir(w, erro, "Erro ao obter id inserido") {
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("Usuário inserido com sucesso! Id: %d", idInserido)))
}

func escreverErroSeExistir(w http.ResponseWriter, internalErro error, mensagem string) bool {
	if internalErro != nil {
		w.Write([]byte(mensagem))
		return true
	}

	return false
}
