package router

import (
	"github.com/dami-pie/napi/src/router/routes"
	"github.com/gorilla/mux"
)

func AddRoutes() *mux.Router {
	r := mux.NewRouter()
	return routes.ConfigRoutes(r)
}
