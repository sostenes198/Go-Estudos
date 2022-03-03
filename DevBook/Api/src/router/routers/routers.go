package routers

import (
	"net/http"
)

type router struct {
	URI                    string
	Method                 string
	Func                   func(w http.ResponseWriter, r *http.Request)
	RequiredAuthentication bool
}
