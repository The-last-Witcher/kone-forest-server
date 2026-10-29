package config

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func DBConfig(){
	dbURL := os.Getenv("DBURL")

	ctx := context.Background()

	pool, err := pgxpool.New(ctx, dbURL)

	if err != nil {
		fmt.Printf("Ошибка при создании пула соединений: %v\n", err)
		os.Exit(1)
	}; defer pool.Close()

	if err := pool.Ping(ctx); err != nil{
		fmt.Fprintf(os.Stderr, "Не удалось подключиться к БД: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Успешное подключеник к БД")
}