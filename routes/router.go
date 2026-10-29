package routes

import "github.com/gorilla/mux"

func Router() *mux.Router{
	router := mux.NewRouter()

	router.HandleFunc("/", Home)
	router.HandleFunc("/about", About)
	router.HandleFunc("/contacts", Contacts)
	router.HandleFunc("/account", Account)

	return router
}