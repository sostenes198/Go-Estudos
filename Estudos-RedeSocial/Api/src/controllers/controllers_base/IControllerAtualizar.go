package controllers_base

import "net/http"

type IControllerAtualizar interface {
	Atualizar(w http.ResponseWriter, r *http.Request)
}
