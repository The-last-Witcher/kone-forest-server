package routes

import (
	"encoding/json"
	"fmt"
	"kone-server/config"
	"net/http"
)

func Chalets(writer http.ResponseWriter, req *http.Request){
	ctx := req.Context()

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
			GROUP BY 
				c.id, c.title, c.description, c.address, 
				c.price, c.rating, c.guestsmax, c.bedrooms, 
				c.bathrooms, c.is_active
			ORDER BY c.title;
	`
	rows, err := config.Pool.Query(ctx, query)

	if err != nil{
		http.Error(writer, "Ошибка при чтении базы данных chalets.go", http.StatusInternalServerError)
		fmt.Fprintf(writer, "Статус ошибки: %v\n", err)
		return 
	}
	defer rows.Close()

	var chalets []config.Chalet

	for rows.Next() {
		var chalet config.Chalet
		var amenitiesJSON []byte

		err := rows.Scan(
			&chalet.ID, &chalet.Title, &chalet.Description, &chalet.Address, 
			&chalet.Price, &chalet.Rating, &chalet.GuestsMax, &chalet.Bedrooms, 
			&chalet.Bathrooms, &chalet.IsActive, &amenitiesJSON,
		)

		if err != nil{
			http.Error(writer, "Ошибка при считывании", http.StatusInternalServerError)
			return
		}

		if err := json.Unmarshal(amenitiesJSON, &chalet.Amenities); err != nil {
			http.Error(writer, "Ошибка парсинга JSON", http.StatusInternalServerError)
			return
		}

		chalets = append(chalets, chalet)
	}

	writer.Header().Set("Contnent-Type", "application/json; charset=utf-8")
	writer.WriteHeader(http.StatusOK)

	json.NewEncoder(writer).Encode(chalets)
}