package routes

import (
	"github.com/ashi5lab/EduLab/pkg/middlewares"
)

//AddClassRouters function
func (server *Server) AddClassRouters() {
	server.Router.HandleFunc("/class", middlewares.SetMiddlewareJSON(server.Handler.CreateClass)).Methods("POST", "OPTIONS")
	server.Router.HandleFunc("/class", middlewares.SetMiddlewareJSON(server.Handler.GetAllClass)).Methods("GET", "OPTIONS")
	server.Router.HandleFunc("/class/{id}", middlewares.SetMiddlewareJSON(server.Handler.GetClass)).Methods("GET", "OPTIONS")
	server.Router.HandleFunc("/class/{id}", middlewares.SetMiddlewareJSON(server.Handler.UpdateClass)).Methods("PUT", "OPTIONS")
	server.Router.HandleFunc("/class/{id}", middlewares.SetMiddlewareJSON(server.Handler.DeleteClass)).Methods("DELETE", "OPTIONS")
}
