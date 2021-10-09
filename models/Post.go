package models

import (
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	
	Title		string `gorm:"not null" json:"title"`
	Content string `gorm:"not null" json:"content"`
	UserID  uint   `json:"userId"`
}

type CreatePostInput struct {
	Title		 string `json:"title" binding:"required,min=1"`
	Content  string `json:"content" binding:"required,min=1"`
}
type UpdatePostInput struct {
	Title		 string `json:"title" binding:"min=1"`
	Content  string `json:"content" binding:"min=1"`
}