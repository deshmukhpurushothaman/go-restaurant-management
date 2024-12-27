package dbrepo

import (
	"context"
	"fmt"
	"time"

	"github.com/deshmukhpurushothaman/go-restaurant-management/internal/models"
)

func (m *postgresDBRepo) GetUserByID(id int) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var user models.User
	err := m.DB.WithContext(ctx).Where("ID=?", id).Find(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (m *postgresDBRepo) GetUserByEmail(email string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var user models.User
	err := m.DB.WithContext(ctx).Where("Email=?", email).Find(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (m *postgresDBRepo) CreateUser(data *models.User) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.WithContext(ctx).Create(data).Error
	if err != nil {
		fmt.Println("Failed to create a food")
		return nil, err
	}

	return data, err
}
