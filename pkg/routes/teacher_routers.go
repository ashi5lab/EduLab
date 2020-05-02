package routes

import "github.com/ashi5lab/EduLab/pkg/middlewares"

//AddTeacherRouters function
func (server *Server) AddTeacherRouters() {
	server.Router.HandleFunc("/teacher", middlewares.SetMiddlewareJSON(server.Handler.CreateTeacher)).Methods("POST")
	server.Router.HandleFunc("/teacher", middlewares.SetMiddlewareJSON(server.Handler.GetAllTeacher)).Methods("GET")
	server.Router.HandleFunc("/teacher/{id}", middlewares.SetMiddlewareJSON(server.Handler.GetTeacher)).Methods("GET")
	server.Router.HandleFunc("/teacher/{id}", middlewares.SetMiddlewareJSON(server.Handler.UpdateTeacher)).Methods("PUT")
	// server.Router.HandleFunc("/teacher/{id}", middlewares.SetMiddlewareJSON(server.Handler.DeleteTeacher)).Methods("DELETE")

}
