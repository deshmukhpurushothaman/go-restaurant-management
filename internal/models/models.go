package models

import (
	"time"

	"gorm.io/gorm"
)

type Category struct {
	ID        uint      `gorm:"primary key;autoIncrement" json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func MigrateCategory(db *gorm.DB) error {
	err := db.AutoMigrate(&Category{})
	return err
}
