package routes

import "github.com/gin-gonic/gin"

func SocmedRoutes() *gin.Engine {
	r := gin.Default()
	//db := database.GetDB()
	//commentController := &controllers.CommentController{Connect: db}

	socmedRoutes := r.Group("/socmed")
	{
		socmedRoutes.GET("/", socmedController.GetAllSocmeds)
		socmedRoutes.POST("/", socmedController.CreateSocmed)
		socmedRoutes.PUT("/:socmed_id", socmedController.UpdateSocmed)
		socmedRoutes.DELETE("/:socmed_id", socmedController.DeleteSocmed)
	}
}
