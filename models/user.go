package models

import (
	"github.com/asaskevich/govalidator"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

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
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
  salt := 10
  _, err := govalidator.ValidateStruct(u)
  if err != nil {
    return err
  }
  password := []byte(u.Password)
  hash, _ := bcrypt.GenerateFromPassword(password, salt)
  u.Password = string(hash)
  return nil
}
