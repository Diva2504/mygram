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
	UserID   int    `json:"user_id"`
}

type ResponseComment struct {
	ID         int           `json:"id"`
	Message    string        `json:"message"`
	Photo_id   string        `json:"photo_id"`
	Updated_at time.Time     `json:"updated_at"`
	Created_at time.Time     `json:"created_at"`
	Photos     ResponsePhoto `json:"photos"`
}
type InputComment struct {
	Message string `json:"message"`
	PhotoID int    `json:"photo_id"`
}

var counter uint

func (db Handlers) GetAllComments(c *gin.Context) {
	var commentRes []ResponseComment
	res, err := repository.GetAllComments(db.Connect)

	for i := range res {
		commentRes[i].ID = int(res[i].ID)
		commentRes[i].Message = res[i].Content
		commentRes[i].Photo_id = string(res[i].PhotoID)
		commentRes[i].Updated_at = res[i].UpdatedAt
		commentRes[i].Created_at = res[i].CreatedAt
		commentRes[i].Photos.ID = int(res[i].PhotoID)
		commentRes[i].Photos.Title = res[i].Photo.Title
		commentRes[i].Photos.Caption = res[i].Photo.Caption
		commentRes[i].Photos.PhotoUrl = res[i].Photo.PhotoUrl
		commentRes[i].Photos.UserID = res[i].Photo.UserID
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
		comment      models.Comment
		result       gin.H
		inputComment InputComment
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
	counter++
	{
		comment.ID = counter
		comment.Content = inputComment.Message
		comment.PhotoID = uint(inputComment.PhotoID)
	}
	result = gin.H{
		"id":         comment.ID,
		"message":    comment.Content,
		"photo_id":   comment.PhotoID,
		"user_id":    comment.UserID,
		"created_at": comment.CreatedAt,
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
	commentId, _ := strconv.Atoi(c.Param("id"))
	_, err := repository.UpdateComment(commentId, &comment, db.Connect)
	if err != nil {
		result = gin.H{
			"message": err,
		}
	}
	result = gin.H{
		"id":         comment.ID,
		"message":    comment.Content,
		"photo_id":   comment.PhotoID,
		"user_id":    comment.UserID,
		"updated_at": comment.UpdatedAt,
	}
	c.JSON(http.StatusOK, result)
}
