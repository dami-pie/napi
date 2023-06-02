package router

import (
	"api/src/router/routes"

	"github.com/gorilla/mux"
)

func AddRoutes() *mux.Router {
	r := mux.NewRouter()
	return routes.ConfigRoutes(r)
}
