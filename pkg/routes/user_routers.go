package routes

func (server *Server) AddUserRouters() {
	server.Router.HandleFunc("/users", server.Handler.CreateUser).Methods("POST")

}
