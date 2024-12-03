package main

import (
	"log"
	"net/http"
	"os"

	"github.com/deshmukhpurushothaman/go-restaurant-management/internal/config"
	"github.com/deshmukhpurushothaman/go-restaurant-management/internal/database"
	"github.com/deshmukhpurushothaman/go-restaurant-management/internal/handlers"
	"github.com/joho/godotenv"
)

var app config.AppConfig

func main() {
	db, err := run()
	if err != nil {
		log.Fatal(err)
	}
	defer db.SQL.Close() // won't lose the connection until the application is closed/crashed

	srv := &http.Server{
		Addr:    ":8080",
		Handler: routes(),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func run() (*database.DB, error) {
	// Load the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v\n", err)
	}

	connetionString := os.Getenv("DATABASE_URL")
	db, err := database.ConnectSQL(connetionString)
	if err != nil {
		log.Fatal("Cannot connect to database!", err)
	}
	log.Println("Connected to database!")

	handlers.NewConfig(&app, db)

	return db, nil
}
