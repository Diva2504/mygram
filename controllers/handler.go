package controllers

import "gorm.io/gorm"

type Handlers struct {
  Connect *gorm.DB
}
