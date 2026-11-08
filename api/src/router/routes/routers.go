package routes

import (
	"api/src/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

// Route representa todas as rotas da API
type Route struct {
	URI          string
	Method       string
	Function     func(http.ResponseWriter, *http.Request)
	RequiredAuth bool
}

// Insere todas as rotas dentro do route
func Config(r *mux.Router) *mux.Router {
	routers := routersUsers
	routers = append(routers, loginRouter)

	for _, route := range routers {
		if route.RequiredAuth {
			r.HandleFunc(route.URI, middlewares.Logger(middlewares.Auth(route.Function))).Methods(route.Method)
		} else {
			r.HandleFunc(route.URI, middlewares.Logger(route.Function)).Methods(route.Method)
		}
	}
	return r
}
