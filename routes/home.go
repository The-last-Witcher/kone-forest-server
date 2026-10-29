package routes

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Home(writer http.ResponseWriter, req *http.Request){
	vars := mux.Vars(req)
	writer.WriteHeader(http.StatusOK)
	fmt.Fprintf(writer, "Все оке. Category: %v\n", vars)
}