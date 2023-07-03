package routes

import (
	"net/http"

	"github.com/dami-pie/napi/src/controllers"
)

var UserRoutes = []Route{
	{
		URI:                "/user/create",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CreateUser,
		RequerAutenticacao: false,
	},
}
