package handlers

import (
	"fmt"
	"net/http"

	"github.com/deshmukhpurushothaman/go-restaurant-management/internal/config"
	"github.com/deshmukhpurushothaman/go-restaurant-management/internal/helpers"
	"github.com/deshmukhpurushothaman/go-restaurant-management/internal/models"
	"github.com/deshmukhpurushothaman/go-restaurant-management/internal/repository"
	"github.com/deshmukhpurushothaman/go-restaurant-management/internal/repository/dbrepo"
	"github.com/deshmukhpurushothaman/go-restaurant-management/internal/utils"
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
	CreateCategory := &models.Category{}
	utils.ParseBody(r, CreateCategory)

	users, err := c.DB.CreateCategory(CreateCategory)
	if err != nil {
		fmt.Println(err)
		helpers.WriteResponse(w, http.StatusBadRequest, err)
		return
	}
	fmt.Println(users)

	// Call WriteResponse to write the response
	err = helpers.WriteResponse(w, http.StatusOK, users)
	if err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
	}
}
