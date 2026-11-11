package routes

import (
	"kone-server/config"
	"kone-server/middlewares"

	"github.com/gorilla/mux"
)

func Router() *mux.Router{
	router := mux.NewRouter()

	router.Use(middlewares.AuthMiddleware)
	router.Use(config.CORSConfig().Handler)

	router.HandleFunc("/", Home)
	router.HandleFunc("/contacts", Contacts)
	router.HandleFunc("/account", Account)
	router.HandleFunc("/chalets", Chalets)
	router.HandleFunc("/chalets/{id}", CurrentChalet)

	return router
}