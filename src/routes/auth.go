package routes

import (
	"net/http"

	"github.com/dami-pie/napi/src/controllers"
)

var AuthRoutes = []Route{
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
