package routes

//AddUserRouters function
func (server *Server) AddUserRouters() {
	server.Router.HandleFunc("/users", server.Handler.CreateUser).Methods("POST")
	server.Router.HandleFunc("/users", server.Handler.GetUsers).Methods("GET")
	server.Router.HandleFunc("/users/{id}", server.Handler.GetUser).Methods("GET")
	server.Router.HandleFunc("/users/{id}", server.Handler.UpdateUser).Methods("PUT")
	server.Router.HandleFunc("/users/{id}", server.Handler.DeleteUser).Methods("DELETE")
}
