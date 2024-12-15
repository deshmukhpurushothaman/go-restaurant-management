package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/deshmukhpurushothaman/go-restaurant-management/internal/helpers"
	"github.com/deshmukhpurushothaman/go-restaurant-management/internal/models"
	"github.com/deshmukhpurushothaman/go-restaurant-management/internal/utils"
)

func (c *Config) GetAllFoods(w http.ResponseWriter, r *http.Request) {
	foods, err := c.DB.GetAllFoods()
	if err != nil {
		helpers.WriteResponse(w, http.StatusBadRequest, err)
		return
	}

	err = helpers.WriteResponse(w, http.StatusOK, foods)
	if err != nil {
		http.Error(w, "Failed to fetch all foods", http.StatusInternalServerError)
	}
}

func (c *Config) GetFoodById(w http.ResponseWriter, r *http.Request) {
	exploded := strings.Split(r.RequestURI, "/")
	foodId, err := strconv.Atoi(exploded[2])
	if err != nil {
		fmt.Println(err)
		helpers.WriteResponse(w, http.StatusBadGateway, err)
		return
	}

	food, err := c.DB.GetFoodById(foodId)
	if err != nil {
		helpers.WriteResponse(w, http.StatusBadRequest, err)
		return
	}

	err = helpers.WriteResponse(w, http.StatusOK, food)
	if err != nil {
		http.Error(w, "Failed to fetch the food", http.StatusInternalServerError)
	}
}

func (c *Config) UpdateFood(w http.ResponseWriter, r *http.Request) {
	UpdateFood := &models.Food{}
	utils.ParseBody(r, UpdateFood)

	food, err := c.DB.UpdateFood(UpdateFood)
	if err != nil {
		fmt.Println(err)
		helpers.WriteResponse(w, http.StatusBadRequest, err)
		return
	}

	// Call WriteResponse to write the response
	err = helpers.WriteResponse(w, http.StatusOK, food)
	if err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
	}
}

func (c *Config) CreateFood(w http.ResponseWriter, r *http.Request) {
	CreateFood := &models.Food{}
	utils.ParseBody(r, CreateFood)

	food, err := c.DB.CreateFood(CreateFood)
	if err != nil {
		helpers.WriteResponse(w, http.StatusBadRequest, err)
		return
	}

	err = helpers.WriteResponse(w, http.StatusOK, food)
	if err != nil {
		http.Error(w, "Failed to create food", http.StatusInternalServerError)
	}
}

func (c *Config) DeleteFood(w http.ResponseWriter, r *http.Request) {
	exploded := strings.Split(r.RequestURI, "/")
	foodId, err := strconv.Atoi(exploded[2])
	if err != nil {
		helpers.WriteResponse(w, http.StatusBadRequest, err)
	}

	food, err := c.DB.DeleteFoodById(foodId)
	if err != nil {
		helpers.WriteResponse(w, http.StatusInternalServerError, err)
	}

	err = helpers.WriteResponse(w, http.StatusOK, food)
	if err != nil {
		http.Error(w, "Failed to delete food", http.StatusInternalServerError)
	}
}
