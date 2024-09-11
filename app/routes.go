package app

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rizkiromadoni/go-shop/app/controllers"
)

func (server *Server) InitializeRoutes() {
	server.Router = mux.NewRouter()
	server.Router.HandleFunc("/", controllers.Home).Methods("GET")

	staticFileDir := http.Dir("./assets/")
	staticFileHandler := http.StripPrefix("/public/", http.FileServer(staticFileDir))
	server.Router.PathPrefix("/public/").Handler(staticFileHandler).Methods("GET")
}
