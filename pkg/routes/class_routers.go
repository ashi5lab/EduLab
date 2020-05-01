package routes

import (
	"github.com/ashi5lab/EduLab/pkg/middlewares"
)

//AddClassRouters function
func (server *Server) AddClassRouters() {
	server.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(server.Handler.CreateClass)).Methods("POST")
	server.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(server.Handler.GetAllClass)).Methods("GET")
	server.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(server.Handler.GetClass)).Methods("GET")
	server.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(server.Handler.UpdateClass)).Methods("PUT")
	server.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(server.Handler.DeleteClass)).Methods("DELETE")
}
