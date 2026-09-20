package routes

import (
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

	for _, route := range routers {
		r.HandleFunc(route.URI, route.Function).Methods(route.Method)
	}
	return r
}
