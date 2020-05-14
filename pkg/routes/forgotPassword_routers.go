package routes

import "github.com/ashi5lab/EduLab/pkg/middlewares"

//AddForgotPasswordRouters function
func (server *Server) AddForgotPasswordRouters() {
	server.Router.HandleFunc("/forgot-password", middlewares.SetMiddlewareJSON(server.Handler.ForgotPassword)).Methods("POST")
	// server.Router.HandleFunc("/teachers", middlewares.SetMiddlewareJSON(server.Handler.GetAllTeacher)).Methods("GET")
	// server.Router.HandleFunc("/teachers/{id}", middlewares.SetMiddlewareJSON(server.Handler.GetTeacher)).Methods("GET")
	// server.Router.HandleFunc("/teachers/{id}", middlewares.SetMiddlewareJSON(server.Handler.UpdateTeacher)).Methods("PUT")
	// server.Router.HandleFunc("/teacher/{id}", middlewares.SetMiddlewareJSON(server.Handler.DeleteTeacher)).Methods("DELETE")

}
