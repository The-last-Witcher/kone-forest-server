package routes

import (
	"fmt"
	"net/http"
)

func About(writer http.ResponseWriter, req *http.Request){
	fmt.Fprint(writer, "Добра пожиловать на стиринцу ээ О")
}