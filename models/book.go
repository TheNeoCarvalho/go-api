package models

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	ID          uint           `json:"id" gorm:primaryKey`
	Name        string         `gorm:name`
	Description string         `gorm:description`
	Price       string         `gorm:price`
	Author      string         `gorm:author`
	ImageUrl    string         `gorm:image_url`
	CreatedAt   time.Time      `gorm:created_at`
	UpdatedAt   time.Time      `gorm:updated_at`
	DeletedAt   gorm.DeletedAt `gorm:"index" json "deleted"`
}
