package routes

import (
	"github.com/ashi5lab/EduLab/pkg/middlewares"
)

//AddUserRouters function
func (server *Server) AddUserRouters() {
	server.Router.HandleFunc("/users", middlewares.SetMiddlewareAuthentication(server.Handler.CreateUser)).Methods("POST")
	server.Router.HandleFunc("/users", middlewares.SetMiddlewareAuthentication(server.Handler.GetUsers)).Methods("GET")
	server.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareAuthentication(server.Handler.GetUser)).Methods("GET")
	server.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareAuthentication(server.Handler.UpdateUser)).Methods("PUT")
	server.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareAuthentication(server.Handler.DeleteUser)).Methods("DELETE")
}
