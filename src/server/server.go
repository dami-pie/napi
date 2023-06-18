package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dami-pie/napi/src/config"
	"github.com/dami-pie/napi/src/router"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type Server struct {
	port   string
	routes *mux.Router
}

func NewServer() *Server {
	config.LoadEnv()

	return &Server{
		port:   config.Port,
		routes: router.ConfigRoutes(),
	}
}

func Run() {
	server := NewServer()

	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"})
	origins := handlers.AllowedOrigins([]string{"*"})

	fmt.Println("Server running on PORT:", server.port)

	log.Fatal(http.ListenAndServeTLS(server.port, config.CertFile, config.KeyFile, handlers.CORS(headers, methods, origins)(server.routes)))
}
