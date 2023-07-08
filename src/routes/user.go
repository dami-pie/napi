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
	{
		URI:                "/user/delete/{id}",
		Metodo:             http.MethodDelete,
		Funcao:             controllers.DeleteUser,
		RequerAutenticacao: false,
	},
	{
		URI:                "/user/update/{id}",
		Metodo:             http.MethodPut,
		Funcao:             controllers.UpdateUser,
		RequerAutenticacao: false,
	},
	{
		URI:                "/user/get/{id}",
		Metodo:             http.MethodGet,
		Funcao:             controllers.GetUser,
		RequerAutenticacao: false,
	},
}
