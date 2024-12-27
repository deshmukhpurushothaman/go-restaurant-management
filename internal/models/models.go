package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint      `gorm:"primary key:autoIncrement" json:"id"`
	Name      string    `json:"name"`
	Email     string    `gorm:"unique" json:"email"`
	Password  string    `json:"password"`
	Role      uint      `gorm:"not null" json:"role"`
	Hotel     Hotel     `gorm:"foreignKey:UserID" json:"hotel"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

type Hotel struct {
	ID     uint   `gorm:"primary key;autoIncrement" json:"id"`
	Name   string `json:"name"`
	UserID uint   `gorm:"not null" json:"user_id"`
}

type Category struct {
	ID        uint      `gorm:"primary key;autoIncrement" json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	Foods     []Food    `gorm:"foreignKey:CategoryID" json:"foods"` // Reverse relation
}

type Food struct {
	ID         uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Name       string    `json:"name"`
	CategoryID uint      `gorm:"not null" json:"category_id"`                                    // Foreign key
	Category   Category  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"category"` // Association
	CreatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func Migrate(db *gorm.DB) error {
	models := []interface{}{
		&Category{},
		&Food{},
		&User{},
		&Hotel{},
	}
	return db.AutoMigrate(models...)
}
