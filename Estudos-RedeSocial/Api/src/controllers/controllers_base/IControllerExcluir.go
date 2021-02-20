package controllers_base

import "net/http"

type IControllerExcluir interface {
	Excluir(w http.ResponseWriter, r *http.Request)
}
