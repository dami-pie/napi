package routes

import (
	"github.com/dami-pie/napi/src/controllers"
	"net/http"
)

var apiRoutes = []Route{
	{
		URI:                "/authenticate",
		Metodo:             http.MethodPost,
		Funcao:             controllers.Login,
		RequerAutenticacao: false,
	},
	{
		URI:                "/",
		Metodo:             http.MethodPost,
		Funcao:             controllers.ValidateOTP,
		RequerAutenticacao: true,
	},
}
