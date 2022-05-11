package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/takadev15/mygram-api/controllers"
	"github.com/takadev15/mygram-api/databases"
)

func SocmedRoutes() *gin.Engine {
	r := gin.Default()
	db := databases.GetDB()
	socmedController := &controllers.Handlers{Connect: db}

	socmedRoutes := r.Group("/socmed")
	{
		socmedRoutes.GET("/", socmedController.GetAllSocmed)
		socmedRoutes.POST("/", socmedController.CreateSocmed)
		socmedRoutes.PUT("/:socmed_id", socmedController.UpdateSocmed)
		socmedRoutes.DELETE("/:socmed_id", socmedController.DeleteSocmed)
	}
	return r
}
