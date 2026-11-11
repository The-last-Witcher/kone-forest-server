package routes

import (
	"encoding/json"
	"errors"
	"kone-server/config"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5"
)

func CurrentChalet(writer http.ResponseWriter, req *http.Request){
	id := mux.Vars(req)["id"]
	ctx := req.Context()
	var chalet config.Chalet
	var amenitiesJSON []byte

	query := `
		SELECT 
			c.id, 
			c.title, 
			c.description, 
			c.address, 
			c.price, 
			c.rating, 
			c.guestsmax, 
			c.bedrooms, 
			c.bathrooms, 
			c.is_active,
			COALESCE(jsonb_agg(a.title) FILTER (WHERE a.title IS NOT NULL), '[]'::jsonb) AS amenities
		FROM chalets c
		LEFT JOIN chalet_amenities ca ON c.id = ca.chalet_id
		LEFT JOIN amenities a ON ca.amenity_id = a.id
		WHERE c.id = $1
		GROUP BY 
			c.id, c.title, c.description, c.address, 
			c.price, c.rating, c.guestsmax, c.bedrooms, 
			c.bathrooms, c.is_active
		ORDER BY c.title
		;
	`

	err := config.Pool.QueryRow(ctx, query, id).Scan(
		&chalet.ID, &chalet.Title, &chalet.Description, 
		&chalet.Address, &chalet.Price, &chalet.Rating, 
		&chalet.GuestsMax, &chalet.Bedrooms, &chalet.Bathrooms, 
		&chalet.IsActive, &amenitiesJSON,
	)
	
	if err != nil{
		if errors.Is(err, pgx.ErrNoRows){
			http.Error(writer, "Нет такого домика", http.StatusNotFound)
			return
		}
		http.Error(writer, "Ошибка сканирования", http.StatusInternalServerError)
		return
	}

	if err := json.Unmarshal(amenitiesJSON, &chalet.Amenities); err != nil {
		http.Error(writer, "Ошибка парсинга json", http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	writer.WriteHeader(http.StatusOK)

	json.NewEncoder(writer).Encode(chalet)
}