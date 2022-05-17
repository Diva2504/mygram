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

type InputSocmed struct {
	Name           string `json:"name"`
	SocialMediaUrl string `json:"social_media_url"`
}

func (db Handlers) GetAllSocmed(c *gin.Context) {
	var socmedRes []ResponseSocmed
	res, err := repository.GetAllSocmed(db.Connect)

	for i := range res {
		socmedRes[i].ID = int(res[i].ID)
		socmedRes[i].Name = res[i].Name
		socmedRes[i].SocialMediaUrl = res[i].SocialMediaUrl
		socmedRes[i].UserID = int(res[i].UserID)
		socmedRes[i].Created_at = res[i].CreatedAt
		socmedRes[i].User.ID = int(res[i].User.ID)
		socmedRes[i].User.Username = res[i].User.Username
		socmedRes[i].User.Email = res[i].User.Email
	}
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
		socmed    models.SocialMedia
		result    gin.H
		inpSocmed InputSocmed
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
	counter++
	{
		socmed.ID = counter
		socmed.Name = inpSocmed.Name
		socmed.SocialMediaUrl = inpSocmed.SocialMediaUrl
	}
	result = gin.H{
		"id":               socmed.ID,
		"name":             socmed.Name,
		"social_media_url": socmed.SocialMediaUrl,
		"user_id":          socmed.SocialMediaUrl,
		"created_at":       socmed.CreatedAt,
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
		"id":               socmed.ID,
		"name":             socmed.Name,
		"social_media_url": socmed.SocialMediaUrl,
		"user_id":          socmed.SocialMediaUrl,
		"updated_at":       socmed.UpdatedAt,
	}
	c.JSON(http.StatusOK, result)
}
