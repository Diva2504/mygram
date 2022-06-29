package controllers

import (
	"net/http"
	"strconv"

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

type InputComment struct {
	Message string `json:"message"`
	PhotoID uint    `json:"photo_id"`
}

var counter uint

func (db Handlers) GetAllComments(c *gin.Context) {
  userData := c.MustGet("id")
	userId := uint(userData.(float64))

	res, err := repository.GetAllComments(db.Connect, userId)
	var result gin.H

	if err != nil {
		result = gin.H{
			"message": err.Error(),
		}
	}
	result = gin.H{
		"data": res,
	}
	c.JSON(http.StatusOK, result)
}

func (db Handlers) CreateComment(c *gin.Context) {
	var (
		comment      models.Comment
		inputComment InputComment
	)
  
  userData := c.MustGet("id")
	userId := uint(userData.(float64))

	if err := c.ShouldBindJSON(&inputComment); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}
  {
    comment.UserID  = userId
    comment.Content = inputComment.Message
    comment.PhotoID = inputComment.PhotoID
  }
	err := repository.CreateComment(&comment, db.Connect)
	if err != nil {
    c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
      "message" : err,
    })
	}
	c.JSON(http.StatusOK, gin.H{
    "message" : "comment uploaded",})
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
