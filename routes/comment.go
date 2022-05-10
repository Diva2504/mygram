package routes

import (
	"github.com/gin-gonic/gin"
)

func CommentRoutes() *gin.Engine {
	r := gin.Default()
	//db := database.GetDB()
	//commentController := &controllers.CommentController{Connect: db}

	commentRoutes := r.Group("/comments")
	{
		commentRoutes.GET("/", commentController.GetAllComments)
		commentRoutes.POST("/", commentController.CreateComment)
		commentRoutes.PUT("/:comment_id", commentController.UpdateComment)
		commentRoutes.DELETE("/:comment_id", commentController.DeleteComment)
	}

	return r
}
