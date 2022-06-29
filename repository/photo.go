package repository

import (
	"github.com/takadev15/mygram-api/models"
	"gorm.io/gorm"
)

func FindByID(db *gorm.DB, ID int) (models.Photo, error) {
	photo := models.Photo{}

	err := db.Preload("Comment").Where("id = ?", ID).Find(&photo).Error

	if err != nil {
		return photo, err
	}

	return photo, nil
}

func FindByUserID(db *gorm.DB, ID uint) ([]models.Photo, error) {
	var photos []models.Photo

	err := db.Preload("Comment").Where("user_id = ?", ID).Find(&photos).Error

	if err != nil {
		return []models.Photo{}, err
	}

	return photos, nil
}


func SavePhotos(data models.Photo, db *gorm.DB) error {
	err := db.Create(&data).Error
	if err != nil {
		return err
	}
	return nil
}

func DeletePhotos(db *gorm.DB, id int)  error {
	var photoDeleted models.Photo

	err := db.Delete(&photoDeleted, id).Error

	if err != nil {
		return err
	}

	return err
}

func UpdatePhoto(db gorm.DB, photo models.Photo, id int) (models.Photo, error) {
	err := db.Where("id = ?", id).Updates(&photo).Error

	if err != nil {
		return models.Photo{}, err
	}

	return photo, nil
}

