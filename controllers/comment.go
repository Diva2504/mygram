package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/takadev15/mygram-api/models"
	"github.com/takadev15/mygram-api/repository"
)

type ResponsePhoto struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	PhotoUrl string `json:"photo_url"`
	UserID   string `json:"user_id"`
}

type ResponseComment struct {
	ID         int           `json:"id"`
	Message    string        `json:"message"`
	Photo_id   string        `json:"photo_id"`
	Updated_at time.Time     `json:"updated_at"`
	Created_at time.Time     `json:"created_at"`
	Photos     ResponsePhoto `json:"photos"`
}

func (db Handlers) GetAllComments(c *gin.Context) {
	var commentRes []ResponseComment
	res, err := repository.GetAllComments(db.Connect)

	for i := range res {
		commentRes[i].ID = int(res[i].ID)
		commentRes[i].Message = res[i].Content
		commentRes[i].Photo_id = string(res[i].PhotoID)
		// commentRes[i].Updated_at = res[i]
		// commentRes[i].Created_at = res[i].Created_at
		// commentRes[i].Photos.ID = res[i].Photo.ID
		// commentRes[i].Photos.Title = res[i].Photo.Title
		// commentRes[i].Photos.Caption = res[i].Photo.Caption
		// commentRes[i].Photos.PhotoUrl = res[i].Photo.PhotoUrl
		//commentRes[i].Photos.UserID = res[i].Photo.User
	}
	var result gin.H

	if err != nil {
		result = gin.H{
			"message": err.Error(),
		}
	}
	result = gin.H{
		"data": commentRes,
	}
	c.JSON(http.StatusOK, result)
}

func (db Handlers) CreateComment(c *gin.Context) {
	var (
		comment models.Comment
		result  gin.H
	)
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}
	err := repository.CreateComment(&comment, db.Connect)
	if err != nil {
		result = gin.H{
			"message": err,
		}
	}
	result = gin.H{
		"result": comment,
	}
	c.JSON(http.StatusOK, result)
}

func (db Handlers) DeleteComment(c *gin.Context) {
	var result gin.H
	requestId := c.Param("id")
	id, _ := strconv.Atoi(requestId)
	err := repository.DeleteComment(id, db.Connect)
	if err != nil {
		result = gin.H{
			"message": err,
		}
	}
	result = gin.H{
		"message": "Comment has been deleted",
	}
	c.JSON(http.StatusOK, result)
}

func (db Handlers) UpdateComment(c *gin.Context) {
	var (
		comment models.Comment
		result  gin.H
	)
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}
	commentId := int(comment.ID)
	_, err := repository.UpdateComment(commentId, &comment, db.Connect)
	if err != nil {
		result = gin.H{
			"message": err,
		}
	}
	result = gin.H{
		"result": comment,
	}
	c.JSON(http.StatusOK, result)
}
