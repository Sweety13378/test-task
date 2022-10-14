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
		Username:    "postgres",
		Password:    "root",
		Host:        "localhost",
		DbName:      "postgres",
		Port:        5432,
		Timeout:     1000,
		MaxAttempts: 5,
	})
	if err != nil {
		log.Fatal(err)
	}

	_, err = database.Conn.Exec(ctx, `
		DROP SCHEMA clients CASCADE;
		CREATE SCHEMA clients
    	CREATE TABLE clients.CLIENTS (
        	id SERIAL PRIMARY KEY NOT NULL,
        	balance int
    	)`)
	if err != nil {
		log.Fatal(err)
	}

	routes.SetupRoutes()
}
