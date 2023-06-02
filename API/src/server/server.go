package server

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type Server struct {
	port   string
	routes *mux.Router
}

func NewServer() *Server {
	return &Server{
		port:   config.Port,
		routes: router.AddRoutes(),
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
