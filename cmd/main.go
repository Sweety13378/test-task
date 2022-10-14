package main

import (
	"context"
	"log"
	"task/api/routes"
	"task/pkg/database"
)

func main() {
	ctx := context.Background()
	err := database.Connect(ctx, database.Config{
		Username:    "root",
		Password:    "root",
		Host:        "localhost",
		Port:        5432,
		Timeout:     1000,
		MaxAttempts: 5,
	})
	if err != nil {
		log.Fatal(err)
	}

	_, err = database.Conn.Exec(ctx, `CREATE SCHEMA clients
    CREATE TABLE clients (
        id SERIAL PRIMARY KEY NOT NULL,
        balance int
        createdAt TIMESTAMP
    )`)
	if err != nil {
		log.Fatal(err)
	}

	routes.SetupRoutes()
}
