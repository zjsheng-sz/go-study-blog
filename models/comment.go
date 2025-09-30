package models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	Content string `gorm:"type:text;not null" json:"content"`
	PostID  uint   `gorm:"not null" json:"post_id"`
	UserID  uint   `gorm:"not null" json:"user_id"`
}