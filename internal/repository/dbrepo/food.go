package dbrepo

import (
	"context"
	"fmt"
	"time"

	"github.com/deshmukhpurushothaman/go-restaurant-management/internal/models"
)

func (m *postgresDBRepo) GetAllFoods() ([]models.Food, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var foods []models.Food
	err := m.DB.WithContext(ctx).Find(foods).Error
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return foods, nil
}

func (m *postgresDBRepo) GetFoodById(id int) (*models.Food, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var food models.Food
	err := m.DB.WithContext(ctx).Where("ID=?", id).Find(&food).Error
	if err != nil {
		return nil, err
	}

	return &food, nil
}

func (m *postgresDBRepo) CreateFood(data *models.Food) (*models.Food, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.WithContext(ctx).Create(data).Error
	if err != nil {
		fmt.Println("Failed to create a food")
		return nil, err
	}

	return data, err
}

func (m *postgresDBRepo) UpdateFood(data *models.Food) (*models.Food, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.WithContext(ctx).Save(&data).Error
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (m *postgresDBRepo) DeleteFoodById(id int) (*models.Food, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var food models.Food
	err := m.DB.WithContext(ctx).Where("ID=?", id).Delete(&food).Error
	if err != nil {
		fmt.Println("Failed to delete food")
		return nil, err
	}

	return &food, nil
}
