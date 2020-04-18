package routes

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ashi5lab/EduLab/pkg/handlers"
	"github.com/gorilla/mux"
)

type Server struct {
	Handler handlers.Server
	Router  *mux.Router
}

func (server *Server) InitializeRoutes() {
	server.Router = mux.NewRouter()

	// Login Route
	server.Router.HandleFunc("/login", server.Handler.Login).Methods("POST")

}

func (server *Server) Run(addr string) {
	fmt.Println("Listening to port 8080")
	log.Fatal(http.ListenAndServe(addr, server.Router))
}
