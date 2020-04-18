package routes

import (
	myHandler "github.com/ashi5lab/EduLab/pkg/handlers"

	"github.com/gorilla/mux"
)

type Server struct {
	Handler myHandler.Server
	Router  *mux.Router
}
