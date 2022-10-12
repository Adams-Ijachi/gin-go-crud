package models

import (
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Name string `json:"name" binding:"required"`
	Body string `json:"body" binding:"required,min=5"`
}
