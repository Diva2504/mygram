package controllers

import (
	"net/http"
	"strconv"
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

type PhotoRequest struct {
  Title     string
  Caption   string
  PhotoUrl  string
}

func (db Handlers) GetAllPhotos(c *gin.Context) {
  var result gin.H
  userData := c.MustGet("id")
	userId := uint(userData.(float64))

  photos, err := repository.FindByUserID(db.Connect, userId)
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

func (db Handlers) GetPhoto(c *gin.Context) {
  inputId := c.Param("id")
  photoID, _ := strconv.Atoi(inputId)
  res, err := repository.FindByID(db.Connect, photoID)

  if err != nil {
    c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
      "message" : err,
    })
  }
  c.JSON(http.StatusOK, gin.H{
    "data": res,
  })
}

func (db Handlers) UploadPhoto(c *gin.Context) {
	var (
		photo   models.Photo
		result gin.H
	)
	if err := c.ShouldBindJSON(&photo); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}

  userData := c.MustGet("id")
	userId := uint(userData.(float64))
  photo.UserID = userId

	err := repository.SavePhotos(photo, db.Connect)
	if err != nil {
		result = gin.H{
			"message": err,
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, result)
	}
	result = gin.H{
		"id":          photo.ID,
		"title":       photo.Title,
    "caption":     photo.Caption,
    "photo_url":   photo.PhotoUrl,
	}
	c.JSON(http.StatusOK, result)
}

func (db Handlers) UpdatePhoto(c *gin.Context) {
	var (
		reqPhoto PhotoRequest
		result  gin.H
	)
	if err := c.ShouldBindJSON(&reqPhoto); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}

  userData := c.MustGet("id")
  userId := uint(userData.(float64))

	photoId := c.Param("id")
	id, _ := strconv.Atoi(photoId)

	task := models.Photo{
		Title:       reqPhoto.Title,
		Caption: reqPhoto.Caption,
    PhotoUrl: reqPhoto.PhotoUrl,
    UserID: userId,
	}
	//task.Description = reqTask.Description


	_, err := repository.UpdatePhoto(*db.Connect, task, id)
	if err != nil {
		result = gin.H{
			"message": err,
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, result)
	}
	result = gin.H{
		"data": task,
	}
	c.JSON(http.StatusCreated, result)
}

func (db Handlers) DeletePhoto(c *gin.Context) {
	photoId := c.Param("id")
	id, _ := strconv.Atoi(photoId)
	err := repository.DeletePhotos(db.Connect, id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "task has been succesfully deleted",
	})
}
