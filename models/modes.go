package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"unique;not null" json:"username"`
	Email    string `gorm:"unique;not null" json:"email"`
	Password string `gorm:"not null" json:"password"`
}

type Post struct {
	gorm.Model
	Title   string `gorm:"not null" json:"title"`
	Content string `gorm:"type:text;not null" json:"content"`
	UserID  uint   `gorm:"not null" json:"user_id"`
}

type Comment struct {
	gorm.Model
	Content string `gorm:"type:text;not null" json:"content"`
	PostID  uint   `gorm:"not null" json:"post_id"`
	UserID  uint   `gorm:"not null" json:"user_id"`
}
