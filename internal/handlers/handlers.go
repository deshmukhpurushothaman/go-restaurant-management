package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

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

func (c *Config) CreateCategory(w http.ResponseWriter, r *http.Request) {
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

func (c *Config) GetCategories(w http.ResponseWriter, r *http.Request) {
	categories, err := c.DB.GetAllCategory()
	if err != nil {
		fmt.Println(err)
		helpers.WriteResponse(w, http.StatusBadRequest, err)
		return
	}

	err = helpers.WriteResponse(w, http.StatusOK, categories)
	if err != nil {
		http.Error(w, "Failed to fetch categories", http.StatusInternalServerError)
	}
}

func (c *Config) GetCategoryById(w http.ResponseWriter, r *http.Request) {
	exploded := strings.Split(r.RequestURI, "/")
	categoryId, err := strconv.Atoi(exploded[2])
	if err != nil {
		fmt.Println(err)
		return
	}

	category, err := c.DB.GetCategoryById(categoryId)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = helpers.WriteResponse(w, http.StatusOK, category)
	if err != nil {
		http.Error(w, "Failed to fetch category", http.StatusInternalServerError)
	}
}

func (c *Config) UpdateCategory(w http.ResponseWriter, r *http.Request) {
	UpdateCategory := &models.Category{}
	utils.ParseBody(r, UpdateCategory)

	category, err := c.DB.CreateCategory(UpdateCategory)
	if err != nil {
		fmt.Println(err)
		helpers.WriteResponse(w, http.StatusBadRequest, err)
		return
	}
	fmt.Println(category)

	// Call WriteResponse to write the response
	err = helpers.WriteResponse(w, http.StatusOK, category)
	if err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
	}
}

func (c *Config) DeleteCatogory(w http.ResponseWriter, r *http.Request) {
	exploded := strings.Split(r.RequestURI, "/")
	categoryId, err := strconv.Atoi(exploded[2])
	if err != nil {
		fmt.Println(err)
		return
	}

	category, err := c.DB.DeleteCatogory(categoryId)
	if err != nil {
		fmt.Println(err)
		helpers.WriteResponse(w, http.StatusBadRequest, err)
		return
	}
	fmt.Println(category)

	// Call WriteResponse to write the response
	err = helpers.WriteResponse(w, http.StatusOK, category)
	if err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
	}
}
