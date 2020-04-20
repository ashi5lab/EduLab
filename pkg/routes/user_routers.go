package routes

func (server *Server) AddUserRouters() {
	server.Router.HandleFunc("/users", server.Handler.CreateUser).Methods("POST")
	server.Router.HandleFunc("/users", server.Handler.GetUsers).Methods("GET")
	server.Router.HandleFunc("/users/{id}", server.Handler.GetUser).Methods("POST")

}
