package controllers

import (
	"net/http"
)

type UsuarioController struct{}

func (_ UsuarioController) Criar(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Criando Usuário"))
}

func (_ UsuarioController) Listar(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Listando todos usuários"))
}

func (_ UsuarioController) ListarPorId(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Listando usuário por Id"))
}

func (_ UsuarioController) Atualizar(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualizando usuário"))
}

func (_ UsuarioController) Excluir(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Excluindo usuário"))
}
