package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/takadev15/mygram-api/models"
	"github.com/takadev15/mygram-api/repository"
)

type PhotosResponse struct {
	ID        uint 
	CreatedAt time.Time
	Title    string
	Caption  string
	PhotoUrl string
	User     *models.User
}

func (db Handlers) GetAllPhotos(c *gin.Context) {
  var result gin.H
  photos, err := repository.GetAllPhotos(db.Connect)
  if err != nil {
    result = gin.H{
      "message" : err,
    }
  } 
  result = gin.H{
    "data" : photos,
  }
  c.JSON(http.StatusOK, result)
}
