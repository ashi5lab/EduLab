package routes

import (
	"github.com/ashi5lab/EduLab/pkg/middlewares"
)

//AddStudentClassMapRouters function
func (server *Server) AddStudentClassMapRouters() {
	server.Router.HandleFunc("/studentclassmaps", middlewares.SetMiddlewareJSON(server.Handler.CreateStudentClassMap)).Methods("POST", "OPTIONS")
	server.Router.HandleFunc("/studentclassmaps", middlewares.SetMiddlewareJSON(server.Handler.GetStudentClassMaps)).Methods("GET", "OPTIONS")
	server.Router.HandleFunc("/studentclassmaps/{id}", middlewares.SetMiddlewareJSON(server.Handler.GetStudentClassMap)).Methods("GET", "OPTIONS")
	server.Router.HandleFunc("/studentclassmaps/{id}", middlewares.SetMiddlewareJSON(server.Handler.UpdateStudentClassMap)).Methods("PUT", "OPTIONS")
}
