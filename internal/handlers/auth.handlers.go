package handlers

import (
	"net/http"

	"github.com/deshmukhpurushothaman/go-restaurant-management/internal/helpers"
	"github.com/deshmukhpurushothaman/go-restaurant-management/internal/inputs"
	"github.com/deshmukhpurushothaman/go-restaurant-management/internal/models"
	"github.com/deshmukhpurushothaman/go-restaurant-management/internal/utils"
	"golang.org/x/crypto/bcrypt"
)

func (c *Config) RegisterUser(w http.ResponseWriter, r *http.Request) {
	User := &inputs.RegisterUserInput{}
	utils.ParseBody(r, User)

	existingUser, err := c.DB.GetUserByEmail(User.Email)
	if err != nil {
		http.Error(w, "Error fetching existing user", http.StatusInternalServerError)
	}

	if existingUser.ID != 0 {
		helpers.WriteResponse(w, http.StatusBadRequest, "User already exists")
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(User.Password), bcrypt.DefaultCost)
	if err != nil {
		helpers.WriteResponse(w, http.StatusBadRequest, err)
	}

	userModel := models.User{
		Name:     User.Name,
		Email:    User.Email,
		Password: string(passwordHash),
		Role:     1,
	}
	user, err := c.DB.CreateUser(&userModel)
	if err != nil {
		helpers.WriteResponse(w, http.StatusBadRequest, err)
		return
	}

	err = helpers.WriteResponse(w, http.StatusOK, user)
	if err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
	}
}
