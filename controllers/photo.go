package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/takadev15/mygram-api/repository"
)

func (db Handlers) GetAllPhotos(*gin.Context) {
	res, err := repository.GetAllPhotos(db.Connect)
}
