package routes

import (
	"api/src/controllers"
	"net/http"
)

var loginRouter = Route{
	URI:          "/login",
	Method:       http.MethodPost,
	Function:     controllers.Login,
	RequiredAuth: false,
}
