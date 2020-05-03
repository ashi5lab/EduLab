package routes

import (
	"github.com/ashi5lab/EduLab/pkg/middlewares"
)

//AddStudentRouters function
func (server *Server) AddStudentRouters() {
	server.Router.HandleFunc("/students", middlewares.SetMiddlewareJSON(server.Handler.CreateStudent)).Methods("POST")
	server.Router.HandleFunc("/students", middlewares.SetMiddlewareJSON(server.Handler.GetStudents)).Methods("GET")
	server.Router.HandleFunc("/students/{id}", middlewares.SetMiddlewareJSON(server.Handler.GetStudent)).Methods("GET")
	server.Router.HandleFunc("/students/{id}", middlewares.SetMiddlewareJSON(server.Handler.UpdateStudent)).Methods("PUT")
}
