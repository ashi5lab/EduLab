package routes

import (
	"github.com/ashi5lab/EduLab/pkg/middlewares"
)

//AddStudentCassMapRouters function
func (server *Server) AddStudentCassMapRouters() {
	server.Router.HandleFunc("/studentclassmaps", middlewares.SetMiddlewareJSON(server.Handler.CreateStudentClassMap)).Methods("POST")
	server.Router.HandleFunc("/studentclassmaps", middlewares.SetMiddlewareJSON(server.Handler.GetStudentClassMaps)).Methods("GET")
	server.Router.HandleFunc("/studentclassmaps/{id}", middlewares.SetMiddlewareJSON(server.Handler.GetStudentClassMap)).Methods("GET")
	server.Router.HandleFunc("/studentclassmaps/{id}", middlewares.SetMiddlewareJSON(server.Handler.UpdateStudentClassMap)).Methods("PUT")
}
