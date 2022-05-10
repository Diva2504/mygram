package repository

import (
	"fmt"

	"github.com/takadev15/mygram-api/models"
	"gorm.io/gorm"
)

func GetAllPhotos(db *gorm.DB) ([]models.Photo, error) {
  var photos []models.Photo
  res := db.Find(&photos)
  // select * form table photos
  //
  fmt.Print(photos)
  if res.Error != nil {
    return nil, res.Error
  }
  return photos, nil
}
