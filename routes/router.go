package routes

import (
	"kone-server/config"
	"kone-server/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

func Router() *mux.Router{
	router := mux.NewRouter()

	router.PathPrefix("/static/").Handler(
		http.StripPrefix("/static/", http.FileServer(http.Dir("./uploads"))),
	)

	router.Use(middlewares.AuthMiddleware)
	router.Use(config.CORSConfig().Handler)

	router.HandleFunc("/", Home)
	router.HandleFunc("/contacts", Contacts).Methods("POST")
	router.HandleFunc("/account", Account)
	router.HandleFunc("/chalets", Chalets)
	router.HandleFunc("/chalets/{id}", CurrentChalet)

	return router
}