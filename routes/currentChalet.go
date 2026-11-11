package routes

import (
	"encoding/json"
	"kone-server/config"
	"net/http"

	"github.com/gorilla/mux"
)

func CurrentChalet(writer http.ResponseWriter, req *http.Request){
	id := mux.Vars(req)["id"]
	ctx := req.Context()
	var chalet config.Chalet

	err := config.Pool.QueryRow(ctx, "SELECT id, title, description, address, price, rating, guestsmax, bedrooms, bathrooms, is_active FROM chalets WHERE id = $1;", id).
		Scan(&chalet.ID, &chalet.Title, &chalet.Description, &chalet.Address, &chalet.Price, &chalet.Rating, &chalet.GuestsMax, &chalet.Bedrooms, &chalet.Bathrooms, &chalet.IsActive)
	
	if err != nil{
		http.Error(writer, "Нет такого домика", http.StatusNotFound)
		return
	}

	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	writer.WriteHeader(http.StatusOK)

	json.NewEncoder(writer).Encode(chalet)
}