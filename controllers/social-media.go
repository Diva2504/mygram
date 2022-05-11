package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/takadev15/mygram-api/models"
	"github.com/takadev15/mygram-api/repository"
)

type ResponseUser struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type ResponseSocmed struct {
	ID             int          `json:"id"`
	Name           string       `json:"name"`
	SocialMediaUrl string       `json:"social_media_url"`
	UserID         int          `json:"user_id"`
	Created_at     time.Time    `json:"created_at"`
	User           ResponseUser `json:"user"`
}

func (db Handlers) GetAllSocmed(c *gin.Context) {
	var socmedRes []ResponseSocmed
	_, err := repository.GetAllSocmed(db.Connect)

	// for  := range res {
	// 	// socmedRes[i].ID = int(res[i].ID)
	// 	// socmedRes[i].Name = res[i].
	// 	// socmedRes[i].Photo_id = string(res[i].PhotoID)
	// 	// commentRes[i].Updated_at = res[i]
	// 	// commentRes[i].Created_at = res[i].Created_at
	// 	// commentRes[i].Photos.ID = res[i].Photo.ID
	// 	// commentRes[i].Photos.Title = res[i].Photo.Title
	// 	// commentRes[i].Photos.Caption = res[i].Photo.Caption
	// 	// commentRes[i].Photos.PhotoUrl = res[i].Photo.PhotoUrl
	// 	//commentRes[i].Photos.UserID = res[i].Photo.User
	// }
	var result gin.H

	if err != nil {
		result = gin.H{
			"message": err.Error(),
		}
	}
	result = gin.H{
		"data": socmedRes,
	}
	c.JSON(http.StatusOK, result)
}

func (db Handlers) CreateSocmed(c *gin.Context) {
	var (
		socmed models.SocialMedia
		result gin.H
	)
	if err := c.ShouldBindJSON(&socmed); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}
	err := repository.CreateSocmed(&socmed, db.Connect)
	if err != nil {
		result = gin.H{
			"message": err,
		}
	}
	result = gin.H{
		"result": socmed,
	}
	c.JSON(http.StatusOK, result)
}

func (db Handlers) DeleteSocmed(c *gin.Context) {
	var result gin.H
	requestId := c.Param("id")
	id, _ := strconv.Atoi(requestId)
	err := repository.DeleteSocmed(id, db.Connect)
	if err != nil {
		result = gin.H{
			"message": err,
		}
	}
	result = gin.H{
		"message": "Social media has been deleted",
	}
	c.JSON(http.StatusOK, result)
}

func (db Handlers) UpdateSocmed(c *gin.Context) {
	var (
		socmed models.SocialMedia
		result gin.H
	)
	if err := c.ShouldBindJSON(&socmed); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}
	commentId := int(socmed.ID)
	_, err := repository.UpdateSocmed(commentId, &socmed, db.Connect)
	if err != nil {
		result = gin.H{
			"message": err,
		}
	}
	result = gin.H{
		"result": socmed,
	}
	c.JSON(http.StatusOK, result)
}
