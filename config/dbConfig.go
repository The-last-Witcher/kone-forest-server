package config

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var Pool *pgxpool.Pool

func DBConfig(){
	dbURL := os.Getenv("DBURL")
	
	ctx := context.Background()

	var err error
	Pool, err = pgxpool.New(ctx, dbURL)

	if err != nil {
		fmt.Printf("Ошибка при создании пула соединений: %v\n", err)
		os.Exit(1)
	}

	if err := Pool.Ping(ctx); err != nil{
		fmt.Fprintf(os.Stderr, "Не удалось подключиться к БД: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Успешное подключеник к БД")
}