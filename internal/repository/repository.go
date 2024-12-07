package repository

import "github.com/deshmukhpurushothaman/go-restaurant-management/internal/models"

type DatabaseRepo interface {
	AllUsers() bool
	GetAllCategory() ([]models.Category, error)
	CreateCategory(data *models.Category) (*models.Category, error)
}
