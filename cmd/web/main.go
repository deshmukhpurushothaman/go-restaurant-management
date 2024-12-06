package main

import (
	"log"
	"net/http"
	"os"

	"github.com/deshmukhpurushothaman/go-restaurant-management/internal/config"
	"github.com/deshmukhpurushothaman/go-restaurant-management/internal/handlers"
	"github.com/deshmukhpurushothaman/go-restaurant-management/internal/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var app config.AppConfig

func main() {
	db, err := run()
	if err != nil {
		log.Fatal(err)
	}

	sqlDb, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	defer sqlDb.Close() // won't lose the connection until the application is closed/crashed

	srv := &http.Server{
		Addr:    ":8080",
		Handler: routes(),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func run() (*gorm.DB, error) {
	// Load the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v\n", err)
	}

	connetionString := os.Getenv("DATABASE_URL")
	// db, err := database.ConnectSQL(connetionString)
	db, err := gorm.Open(postgres.Open(connetionString))
	if err != nil {
		log.Fatal("Cannot connect to database!", err)
	}
	log.Println("Connected to database!")

	handlers.Repo = handlers.NewConfig(&app, db)

	err = models.Migrate(db)
	if err != nil {
		log.Fatal("Migration failed")
	}
	log.Println("Database migration successful!")

	return db, nil
}
