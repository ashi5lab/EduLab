package routes

import (
	"github.com/ashi5lab/EduLab/pkg/middlewares"
)

//AddLoginRouters function
func (server *Server) AddLoginRouters() {
	server.Router.HandleFunc("/login", middlewares.SetMiddlewareAuthentication(server.Handler.Login)).Methods("POST")

}
