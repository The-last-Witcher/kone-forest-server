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
	router.HandleFunc("/about", About)
	router.HandleFunc("/contacts", Contacts)
	router.HandleFunc("/account", Account)

	return router
}