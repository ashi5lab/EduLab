package routes

import (
	"github.com/ashi5lab/EduLab/pkg/middlewares"
)

//AddUserRouters function
func (server *Server) AddUserRouters() {
	server.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(server.Handler.CreateUser)).Methods("POST", "OPTIONS")
	server.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(server.Handler.GetUsers)).Methods("POST", "OPTIONS")
	server.Router.HandleFunc("/getprofile", middlewares.SetMiddlewareJSON(server.Handler.GetProfile)).Methods("GET", "OPTIONS")
	server.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(server.Handler.GetUser))).Methods("GET", "OPTIONS")
	server.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(server.Handler.UpdateUser))).Methods("PUT", "OPTIONS")
	server.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(server.Handler.DeleteUser))).Methods("DELETE", "OPTIONS")
}
