package repository

import (
	"github.com/takadev15/mygram-api/models"
	"gorm.io/gorm"
)

func GetAllSocmed(db *gorm.DB) ([]models.SocialMedia, error) {
	var socmed []models.SocialMedia
	result := db.Find(&socmed)

	if result.Error != nil {
		return nil, result.Error
	} else {
		if result.RowsAffected <= 0 {
			return nil, result.Error
		} else {
			return socmed, result.Error
		}
	}
}

func CreateSocmed(req *models.SocialMedia, db *gorm.DB) error {
	result := db.Create(&req)

	if result.Error != nil {
		return result.Error
	}
	return nil
}

func UpdateSocmed(id int, data *models.SocialMedia, db *gorm.DB) (models.SocialMedia, error) {
	var socmed models.SocialMedia
	err := db.Model(&socmed).Where("id = ?", id).Updates(&data).Error
	if err != nil {
		return models.SocialMedia{}, err
	}
	return socmed, err
}

func DeleteSocmed(id int, db *gorm.DB) error {
	var socmed models.SocialMedia

	del := db.Delete(&socmed, id)

	if del.Error != nil {
		return del.Error
	} else {
		return nil
	}

}
