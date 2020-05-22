package routes

import "github.com/ashi5lab/EduLab/pkg/middlewares"

//AddTeacherRouters function
func (server *Server) AddTeacherRouters() {
	server.Router.HandleFunc("/teachers", middlewares.SetMiddlewareJSON(server.Handler.CreateTeacher)).Methods("POST", "OPTIONS")
	server.Router.HandleFunc("/teachers", middlewares.SetMiddlewareJSON(server.Handler.GetAllTeacher)).Methods("GET", "OPTIONS")
	server.Router.HandleFunc("/teachers/{id}", middlewares.SetMiddlewareJSON(server.Handler.GetTeacher)).Methods("GET", "OPTIONS")
	server.Router.HandleFunc("/teachers/{id}", middlewares.SetMiddlewareJSON(server.Handler.UpdateTeacher)).Methods("PUT", "OPTIONS")
	// server.Router.HandleFunc("/teacher/{id}", middlewares.SetMiddlewareJSON(server.Handler.DeleteTeacher)).Methods("DELETE")

}
