package routes

import (
	"net/http"

	"api/src/middlewares"

	"github.com/gorilla/mux"
)

type Route struct {
	URI                string
	Metodo             string
	Funcao             func(http.ResponseWriter, *http.Request)
	RequerAutenticacao bool
}

// Configurar coloca todas as rotas dentro do router
func ConfigRoutes(r *mux.Router) *mux.Router {
	routes := apiRoutes

	for _, route := range routes {
		if route.RequerAutenticacao {
			r.HandleFunc(route.URI, middlewares.Logger(middlewares.AuthenticateUser(route.Funcao))).Methods(route.Metodo)
		} else {
			r.HandleFunc(route.URI, middlewares.Logger(route.Funcao)).Methods(route.Metodo)
		}

	}

	return r
}
