package routes

import (
	"net/http"

	"api/src/controllers"
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
