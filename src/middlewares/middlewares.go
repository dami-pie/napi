package middlewares

import (
	"github.com/dami-pie/napi/src/auth"
	"github.com/dami-pie/napi/src/responses"
	"log"
	"net/http"
)

// Escreve as informações da requisição no terminal
func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		log.Printf("[%s %s %s]\n", req.Method, req.RequestURI, req.Host)
		next(res, req)
	}
}

// Verifica se o usuário está autenticado, se estiver ele irá para próxima função
func AuthenticateUser(next http.HandlerFunc) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		if err := auth.ValidateToken(req); err != nil {
			log.Printf("[%s %s %s] -> %s\n", req.Method, req.RequestURI, req.Host, err.Error())
			responses.Err(res, http.StatusUnauthorized, err)
		} else {
			next(res, req)
		}

	}
}
