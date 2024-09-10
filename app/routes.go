package app

import (
	"github.com/gorilla/mux"
	"github.com/rizkiromadoni/go-shop/app/controllers"
)

func (server *Server) InitializeRoutes() {
	server.Router = mux.NewRouter()
	server.Router.HandleFunc("/", controllers.Home).Methods("GET")
}
