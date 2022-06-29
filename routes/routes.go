package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/takadev15/mygram-api/controllers"
	"github.com/takadev15/mygram-api/databases"
)

func RouteList() *gin.Engine {
  r := gin.Default()
  db := databases.GetDB()
  handler := controllers.Handlers{ Connect: db}

  userRoutes := r.Group("/user")
  {
    userRoutes.POST("/register", handler.UserRegister)
    userRoutes.POST("/login", handler.UserLogin)
  }

  photoRoutes := r.Group("/photos")
  {
    photoRoutes.GET("/")
    photoRoutes.POST("/")
    photoRoutes.GET("/:id")
    photoRoutes.PUT("/:id")
    photoRoutes.DELETE("/:id")
  }

	// commentRoutes := r.Group("/comments")
	{
		// commentRoutes.GET("/", handler.GetAllComments)
		// commentRoutes.POST("/", handler.CreateComment)
		// commentRoutes.PUT("/:comment_id", handler.UpdateComment)
		// commentRoutes.DELETE("/:comment_id", handler.DeleteComment)
	}

	// socmedRoutes := r.Group("/socmed")
	// {
	// 	socmedRoutes.GET("/", handler.GetAllSocmeds)
	// 	socmedRoutes.POST("/", handler.CreateSocmed)
	// 	socmedRoutes.PUT("/:id", handler.UpdateSocmed)
	// 	socmedRoutes.DELETE("/:id", handler.DeleteSocmed)
	// }

  return r
}
