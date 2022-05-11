package repository

import (
	"github.com/takadev15/mygram-api/models"
	"gorm.io/gorm"
)

func GetAllComments(db *gorm.DB) ([]models.Comment, error) {
	var comments []models.Comment
	result := db.Find(&comments)

	if result.Error != nil {
		return nil, result.Error
	} else {
		if result.RowsAffected <= 0 {
			return nil, result.Error
		} else {
			return comments, result.Error
		}
	}
}

func CreateComment(input *models.Comment, db *gorm.DB) error {
	result := db.Debug().Create(&input)

	if result.Error != nil {
		return result.Error
	}
	return nil
}

func UpdateComment(id int, data *models.Comment, db *gorm.DB) (models.Comment, error) {
	var comment models.Comment
	err := db.Model(&comment).Where("id = ?", id).Updates(&data).Error
	if err != nil {
		return models.Comment{}, err
	}
	return comment, err
}

func DeleteComment(id int, db *gorm.DB) error {
	var comment models.Comment

	del := db.Delete(&comment, id)

	if del.Error != nil {
		return del.Error
	} else {
		return nil
	}

}
