package routes

import (
	"encoding/json"
	"kone-server/config"
	"kone-server/middlewares"
	"net/http"
)

func Contacts(writer http.ResponseWriter, req *http.Request){
	ctx := req.Context()
	writer.Header().Set("Content-Type", "application/json;charset=utf-8")
	
	var reqMsg config.InquiryReq
	if err := json.NewDecoder(req.Body).Decode(&reqMsg); err != nil{
		writer.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writer).Encode(map[string]string{"error": "Нверный формат JSON"})
		return
	}

	if !middlewares.ValidateFields(reqMsg){
		http.Error(writer, "Неверный формат JSON", http.StatusBadRequest)
		return
	}

	switch reqMsg.Topic {
		case "booking_question", "adding_chalet", "rescheduling_canceling", "other":
			
		default:
			writer.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(writer).Encode(map[string]string{"error": "Выбрана несуществующая тема обращения"})
			return
	}

	query := `
		INSERT INTO contact_inquiries (first_name, last_name, email, topic, message)
		VALUES ($1, $2, $3, $4, $5);	
	`

	_, err := config.Pool.Exec(ctx, query, 
		reqMsg.FirstName, 
		reqMsg.LastName, 
		reqMsg.Email, 
		reqMsg.Topic, 
		reqMsg.Message,
	)

	if err != nil {
		http.Error(writer, "Ошибка при записи обращения", http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusOK)
}