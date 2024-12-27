package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/deshmukhpurushothaman/go-restaurant-management/internal/helpers"
	"github.com/deshmukhpurushothaman/go-restaurant-management/internal/inputs"
	"github.com/deshmukhpurushothaman/go-restaurant-management/internal/models"
	"github.com/deshmukhpurushothaman/go-restaurant-management/internal/utils"
	"github.com/golang-jwt/jwt/v5"
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

func (c *Config) LoginHandler(w http.ResponseWriter, r *http.Request) {
	User := &inputs.LoginUserInput{}
	utils.ParseBody(r, User)

	user, err := c.DB.GetUserByEmail(User.Email)
	if err != nil {
		helpers.WriteResponse(w, http.StatusNotFound, "user not found")
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(User.Password))
	if err != nil {
		helpers.WriteResponse(w, http.StatusNonAuthoritativeInfo, "Incorrect password ")
		return
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"issuer": strconv.Itoa(int(user.ID)),
		"role":   strconv.Itoa(int(user.Role)),
		"exp":    time.Now().Add(time.Hour * 24).Unix(), // 1 day
	})

	token, err := claims.SignedString([]byte("secret"))
	if err != nil {
		fmt.Println(err)
		helpers.WriteResponse(w, http.StatusInternalServerError, "Login failed")
		return
	}

	err = helpers.WriteResponse(w, http.StatusOK, token)
	if err != nil {
		http.Error(w, "Failed to authenticate user", http.StatusInternalServerError)
	}
}
