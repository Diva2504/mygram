package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/takadev15/mygram-api/controllers"
	"github.com/takadev15/mygram-api/databases"
)

func PhotoRoutes() *gin.Engine {
  routes := gin.Default()
  db := databases.GetDB()
  photoHandler := &controllers.Handlers{Connect: db}
  photoRoutes := routes.Group("/photos")
  photoRoutes.GET("/", photoHandler.GetAllPhotos)
  return nil
}
