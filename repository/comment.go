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

func CreateComment(db *gorm.DB, input *models.Comment) error {
	result := db.Create(&input)

	if result.Error != nil {
		return result.Error
	}
	return nil
}

//func UpdateComment(db *gorm.DB, id uint, data *models.Comment) (models.Comment, error) {
//var comment models.Comment

//}

func DeleteComment(id uint, db *gorm.DB) error {
	var comment models.Comment

	del := db.Delete(&comment, id)

	if del.Error != nil {
		return del.Error
	} else {
		return nil
	}

}
