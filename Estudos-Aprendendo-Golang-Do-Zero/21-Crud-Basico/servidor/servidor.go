package servidor

import (
	"21-Crud-Basico/banco"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strconv"
)

type usuario struct{
	ID uint32 `json:"id"`
	Nome string `json:"nome"`
	Email string `json:"email"`
}

func CriarUsuario(w http.ResponseWriter, r *http.Request){
	corpoRequisicao, erro := ioutil.ReadAll(r.Body)
	if erro != nil{
		w.Write([]byte("Falha ao ler o corpo da requisição"))
		return
	}

	var usuario usuario
	if erro = json.Unmarshal(corpoRequisicao, &usuario); erro != nil{
		w.Write([]byte("Erro ao converter para estruct"))
		return
	}

	db, erro := banco.Conectar()
	defer db.Close()
	if erro != nil{
		w.Write([]byte("Erro ao conectar no banco de dados"))
		return
	}

	// PREPARE STATEMENT
	statement, erro := db.Prepare("insert into usuarios(nome, email) values(?, ?)")
	defer statement.Close()
	if erro != nil{
		w.Write([]byte("Erro ao criar statement"))
		return
	}

	insercao, erro := statement.Exec(usuario.Nome, usuario.Email)
	if erro != nil{
		w.Write([]byte("Erro ao executar o statement"))
		return
	}

	idInserido, erro := insercao.LastInsertId()
	if erro != nil{
		w.Write([]byte("Erro ao obter id inserido"))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("Usuário inserido com sucesso! ID: %d", idInserido)))
}

func BuscarUsuarios(w http.ResponseWriter, r *http.Request){
	db, erro := banco.Conectar()
	defer db.Close()
	if erro != nil{
		w.Write([]byte("Erro ao conectar no banco de dados"))
		return
	}

	linhas, erro := db.Query("select * from usuarios")
	defer linhas.Close()
	if erro != nil{
		w.Write([]byte("Erro ao buscar os usuários"))
		return
	}

	var usuarios []usuario
	for linhas.Next(){
		var usuario usuario

		if erro := linhas.Scan(&usuario.ID, &usuario.Nome, &usuario.Email); erro != nil{
			w.Write([]byte("Erro ao escaner usuário"))
			return
		}
		usuarios = append(usuarios, usuario)
	}

	w.WriteHeader(http.StatusOK)
	if erro := json.NewEncoder(w).Encode(usuarios); erro != nil{
		w.Write([]byte("Erro ao converter os usuários para JSON"))
		return
	}
}

func BuscarUsuario(w http.ResponseWriter, r *http.Request){

	parametros := mux.Vars(r)

	ID, erro := strconv.ParseUint(parametros["id"], 10, 32)
	if erro != nil{
		w.Write([]byte("Erro ao converter o parâmetro para inteiro"))
		return
	}

	db, erro := banco.Conectar()
	defer db.Close()
	if erro != nil{
		w.Write([]byte("Erro ao conectar no banco de dados"))
		return
	}

	linha, erro := db.Query("select * from usuarios where id = ?", ID)
	defer linha.Close()
	if erro != nil{
		w.Write([]byte("Erro ao buscar o usuário"))
		return
	}

	var usuario usuario
	if linha.Next(){
		if erro := linha.Scan(&usuario.ID, &usuario.Nome, &usuario.Email); erro != nil{
			w.Write([]byte("Erro ao escaner usuário"))
			return
		}
	}
	w.WriteHeader(http.StatusOK)
	if erro := json.NewEncoder(w).Encode(usuario); erro != nil{
		w.Write([]byte("Erro ao converter os usuários para JSON"))
		return
	}
}

func AtualizarUsuario(w http.ResponseWriter, r *http.Request){
	parametros := mux.Vars(r)

	ID, erro := strconv.ParseUint(parametros["id"], 10, 32)
	if erro != nil{
		w.Write([]byte("Erro ao converter o parâmetro para inteiro"))
		return
	}

	corpoRequisicao, erro := ioutil.ReadAll(r.Body)
	if erro != nil{
		w.Write([]byte("Falha ao ler o corpo da requisição"))
		return
	}

	var usuario usuario
	if erro = json.Unmarshal(corpoRequisicao, &usuario); erro != nil{
		w.Write([]byte("Erro ao converter para estruct"))
		return
	}

	db, erro := banco.Conectar()
	defer db.Close()
	if erro != nil{
		w.Write([]byte("Erro ao conectar no banco de dados"))
		return
	}

	statement, erro := db.Prepare("update usuarios set nome = ?, email = ? where id = ?")
	defer statement.Close()
	if erro != nil{
		w.Write([]byte("Erro ao criar statement"))
		return
	}

	if _, erro := statement.Exec(usuario.Nome, usuario.Email, ID); erro != nil{
		w.Write([]byte("Erro ao atualizar usuário"))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func DeletarUsuario(w http.ResponseWriter, r *http.Request){
	parametros := mux.Vars(r)

	ID, erro := strconv.ParseUint(parametros["id"], 10, 32)
	if erro != nil{
		w.Write([]byte("Erro ao converter o parâmetro para inteiro"))
		return
	}

	db, erro := banco.Conectar()
	defer db.Close()
	if erro != nil{
		w.Write([]byte("Erro ao conectar no banco de dados"))
		return
	}

	statement, erro := db.Prepare("delete from usuarios where id = ?")
	defer statement.Close()
	if erro != nil{
		w.Write([]byte("Erro ao criar statement"))
		return
	}

	if _, erro := statement.Exec(ID); erro != nil{
		w.Write([]byte("Erro ao excluir usuário"))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
