package controllers

import (
	// "net/http"
	"time"

	// "github.com/gin-gonic/gin"
	// "github.com/takadev15/mygram-api/repository"
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

type ResponseTest struct {
  ID int
  Message string
}

// func (db Handlers) GetAllComments(c *gin.Context) {
// 	var (
//     commentRes []ResponseComment
//     testResponse ResponseTest
//     )
// 	res, err := repository.GetAllComments(db.Connect)
//
// 	for i := range res {
// 		commentRes[i].ID = int(res[i].ID)
// 		commentRes[i].Message = res[i].Content
// 		commentRes[i].Photo_id = string(res[i].PhotoID)
// 		commentRes[i].Updated_at = res[i].Updated_at
// 		commentRes[i].Created_at = res[i].Created_at
// 		commentRes[i].Photos.ID = res[i].Photo.ID
// 		commentRes[i].Photos.Title = res[i].Photo.Title
// 		commentRes[i].Photos.Caption = res[i].Photo.Caption
// 		commentRes[i].Photos.PhotoUrl = res[i].Photo.PhotoUrl
// 		//commentRes[i].Photos.UserID = res[i].Photo.User
// 	}
// 	var result gin.H
//
// 	if err != nil {
// 		result = gin.H{
// 			"message": err.Error(),
// 		}
// 	}
// 	result = gin.H{
// 		"data": testResponse.ID,
// 	}
// 	c.JSON(http.StatusOK, result)
// }
