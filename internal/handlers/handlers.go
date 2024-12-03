package handlers

import (
	"fmt"
	"net/http"

	"github.com/deshmukhpurushothaman/go-restaurant-management/internal/config"
	"github.com/deshmukhpurushothaman/go-restaurant-management/internal/helpers"
	"github.com/deshmukhpurushothaman/go-restaurant-management/internal/repository"
	"github.com/deshmukhpurushothaman/go-restaurant-management/internal/repository/dbrepo"
	"gorm.io/gorm"
)

var Repo *Config

type Config struct {
	App *config.AppConfig
	DB  repository.DatabaseRepo
}

func NewConfig(a *config.AppConfig, db *gorm.DB) *Config {
	return &Config{
		App: a,
		DB:  dbrepo.NewPostgresRepo(db, a),
	}
}

func (c *Config) DummyTest(w http.ResponseWriter, r *http.Request) {
	// Example data to send in the response
	responseData := map[string]string{
		"message": "Hello, World!",
	}

	users := c.DB.AllUsers()
	fmt.Println(users)

	// Call WriteResponse to write the response
	err := helpers.WriteResponse(w, http.StatusOK, responseData)
	if err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
	}
}
