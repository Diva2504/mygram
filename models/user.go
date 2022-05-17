package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string
	Email    string
	Password string
	Age      uint
}

type SocialMedia struct {
	gorm.Model
	Name           string `gorm:"type:varchar(255)" json:"name" binding:"required"`
	SocialMediaUrl string `gorm:"type:text" json:"socmed_url" binding:"required"`
	UserID         uint   `gorm:"not null" json:"user_id"`
	User           *User
}
