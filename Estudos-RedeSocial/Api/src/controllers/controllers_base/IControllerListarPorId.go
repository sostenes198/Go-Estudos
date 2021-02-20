package controllers_base

import "net/http"

type IControllerListarPorId interface {
	ListarPorId(w http.ResponseWriter, r *http.Request)
}
