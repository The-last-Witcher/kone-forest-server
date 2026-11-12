package config

import (
	"net/http"

	"github.com/go-chi/cors"
)

func CORSConfig() *cors.Cors{
	
	newCors := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000", "http://192.168.31.*:3000"},
		AllowedMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodOptions},
		AllowedHeaders: []string{"Content-Type", "Authorization"},
	})

	return newCors
}