package repository

import (
	"github.com/takadev15/mygram-api/models"
	"gorm.io/gorm"
)

func FindByID(db *gorm.DB, ID int) (models.Photo, error) {
	photo := models.Photo{}

	err := db.Where("id = ?", ID).Find(&photo).Error

	if err != nil {
		return photo, err
	}

	return photo, nil
}

func FindByUserID(db *gorm.DB, ID int) ([]models.Photo, error) {
	var photos []models.Photo

	err := db.Where("user_id = ?", ID).Find(&photos).Error

	if err != nil {
		return []models.Photo{}, err
	}

	return photos, nil
}

