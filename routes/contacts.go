package routes

import (
	"fmt"
	"net/http"
)

func Contacts(writer http.ResponseWriter, req *http.Request){
	fmt.Fprint(writer, "Contacts")
}