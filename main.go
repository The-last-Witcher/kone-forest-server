package main

import (
	"errors"
	"fmt"
	"kone-server/routes"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main(){
	godotenv.Load()

	port := os.Getenv("SPORT")
	listener, err := net.Listen("tcp", port)

	if err != nil{
		log.Fatalf("Ошибка при занятии порта: %v\n", err)
	}

	server := &http.Server{
		Handler: routes.Router(),
	}

	go func ()  {
		if err := server.Serve(listener); err != nil && !errors.Is(err, http.ErrServerClosed){
			log.Fatalf("Ошибка сервера: %v", err)
		}
	}()

	fmt.Printf("Сервер успешно запущен: http://localhost%s\n", port)

	select {}
}