package dbrepo

import (
	"context"
	"fmt"
	"time"

	"github.com/deshmukhpurushothaman/go-restaurant-management/internal/models"
)

func (m *postgresDBRepo) AllUsers() bool {
	return true
}

func (m *postgresDBRepo) GetAllCategory() ([]models.Category, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var Categories []models.Category
	err := m.DB.WithContext(ctx).Find(&Categories).Error
	if err != nil {
		return nil, fmt.Errorf("failed to fetch categories: %w", err)
	}

	return Categories, nil
}

func (m *postgresDBRepo) CreateCategory(data *models.Category) (*models.Category, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.WithContext(ctx).Create(&data).Error
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (m *postgresDBRepo) UpdateCategory(data *models.Category) (*models.Category, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.WithContext(ctx).Save(&data).Error
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (m *postgresDBRepo) DeleteCatogory(id int) (*models.Category, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var category models.Category
	err := m.DB.WithContext(ctx).Where("ID=?", id).Delete(category).Error
	if err != nil {
		return nil, err
	}

	return &category, nil
}
