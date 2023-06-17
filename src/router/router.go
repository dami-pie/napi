package router

import (
	"github.com/dami-pie/napi/src/middlewares"
	"github.com/dami-pie/napi/src/routes"
	"github.com/gorilla/mux"
)

// Configurar coloca todas as rotas dentro do router
func ConfigRoutes() *mux.Router {
	r := mux.NewRouter()
	routes := routes.AuthRoutes

	for _, route := range routes {
		if route.RequerAutenticacao {
			r.HandleFunc(route.URI, middlewares.Logger(middlewares.AuthenticateUser(route.Funcao))).Methods(route.Metodo)
		} else {
			r.HandleFunc(route.URI, middlewares.Logger(route.Funcao)).Methods(route.Metodo)
		}

	}

	return r
}
