package controllers_base

import "net/http"

type IControllerCriar interface {
	Criar(w http.ResponseWriter, r *http.Request)
}
