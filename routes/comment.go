package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/takadev15/mygram-api/controllers"
	"github.com/takadev15/mygram-api/databases"
)

func CommentRoutes() *gin.Engine {
	r := gin.Default()
	db := databases.GetDB()
	commentController := &controllers.Handlers{Connect: db}

	commentRoutes := r.Group("/comments")
	{
		commentRoutes.GET("/", commentController.GetAllComments)
		commentRoutes.POST("/", commentController.CreateComment)
		commentRoutes.PUT("/:comment_id", commentController.UpdateComment)
		commentRoutes.DELETE("/:comment_id", commentController.DeleteComment)
	}

	return r
}
