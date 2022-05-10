package models

import (
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	UserID  uint   `gorm:"not null" json:"user_id"`
	PhotoID uint   `gorm:"not null" json:"photo_id"`
	Content string `gorm:"type:text" json:"message" binding:"required"`
}
