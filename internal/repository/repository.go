package repository

import "github.com/deshmukhpurushothaman/go-restaurant-management/internal/models"

type DatabaseRepo interface {
	AllUsers() bool
	GetAllCategory() ([]models.Category, error)
	CreateCategory(data *models.Category) (*models.Category, error)
	GetCategoryById(id int) (*models.Category, error)
	DeleteCatogory(id int) (*models.Category, error)
	UpdateCategory(data *models.Category) (*models.Category, error)

	GetAllFoods() ([]models.Food, error)
	GetFoodById(id int) (*models.Food, error)
	CreateFood(data *models.Food) (*models.Food, error)
	UpdateFood(data *models.Food) (*models.Food, error)
	DeleteFoodById(id int) (*models.Food, error)
}
