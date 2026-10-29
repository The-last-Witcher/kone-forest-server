package routes

import (
	"fmt"
	"net/http"
)

func Account(writer http.ResponseWriter, req *http.Request){
	fmt.Fprint(writer, "Account")
}