package routes

import (
	"encoding/json"
	"fmt"
	"kone-server/config"
	"net/http"
	"kone-server/middlewares"
)

func Home(writer http.ResponseWriter, req *http.Request){
	ctx := req.Context()

	// поиск пользователя
	userID, _ := ctx.Value(middlewares.UserIDKey).(int)
	var currentUser *config.User

	if userID > 0{
		var tmpUser config.User
		err := config.Pool.QueryRow(ctx, "SELECT id, name FROM users WHERE id = $1", userID).
			Scan(&tmpUser.ID, &tmpUser.Name)

		if err != nil{
			currentUser = &tmpUser
		}
	}

	// получение списка шале
	var chalets []config.Chalet

	rows, err := config.Pool.Query(ctx, "SELECT id, title, description, address, price, rating, guestsmax, bedrooms, bathrooms, is_active FROM chalets;")
	if err != nil {
		http.Error(writer, "Ошибка бд", http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	
	// запись списка в массив
	for rows.Next(){
		var chalet config.Chalet

		err := rows.Scan(&chalet.ID, &chalet.Title, &chalet.Description, &chalet.Address, &chalet.Price, &chalet.Rating, &chalet.GuestsMax, &chalet.Bedrooms, &chalet.Bathrooms, &chalet.IsActive)
		if err != nil{
			fmt.Fprintf(writer, "Error: %v\n", err)
			http.Error(writer, "Ошибка при сканировании полей", http.StatusInternalServerError)
			return
		}

		chalets = append(chalets, chalet)
	}

	response := config.DTO{
		User: currentUser,
		Chalet: chalets,
	}
	
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	writer.WriteHeader(http.StatusOK)

	json.NewEncoder(writer).Encode(response)
}