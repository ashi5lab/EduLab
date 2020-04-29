package routes

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ashi5lab/EduLab/pkg/handlers"
	"github.com/ashi5lab/EduLab/pkg/middlewares"
	"github.com/gorilla/mux"
)

//Server struct
type Server struct {
	Handler handlers.Server
	Router  *mux.Router
}

//InitializeRoutes method
func (server *Server) InitializeRoutes() {
	server.Router = mux.NewRouter()

	server.Router.HandleFunc("/login", middlewares.SetMiddlewareAuthentication(seserver.Handler.Login)).Methods("POST")
	// User Route
	server.AddUserRouters()

}

//Run function
func (server *Server) Run(addr string) {
	fmt.Println("Listening to port 8080")
	log.Fatal(http.ListenAndServe(addr, server.Router))
}
