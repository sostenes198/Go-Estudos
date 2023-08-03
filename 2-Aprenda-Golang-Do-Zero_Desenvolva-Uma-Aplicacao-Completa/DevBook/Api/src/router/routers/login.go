package routers

import (
	"devbook/src/controllers"
	"net/http"
)

var LoginRouters = router{
	URI:                    "/login",
	Method:                 http.MethodPost,
	Func:                   controllers.Login,
	RequiredAuthentication: false,
}
