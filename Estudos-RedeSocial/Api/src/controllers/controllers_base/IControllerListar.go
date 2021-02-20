package controllers_base

import "net/http"

type IControllerListar interface {
	Listar(w http.ResponseWriter, r *http.Request)
}
