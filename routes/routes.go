package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/takadev15/mygram-api/controllers"
	"github.com/takadev15/mygram-api/databases"
	middleware "github.com/takadev15/mygram-api/middlewares"
)

func RouteList() *gin.Engine {
  r := gin.Default()
  db := databases.GetDB()
  handler := controllers.Handlers{ Connect: db}

  userRoutes := r.Group("/user")
  {
    userRoutes.POST("/register", handler.UserRegister)
    userRoutes.POST("/login", handler.UserLogin)
    userRoutes.PUT("/", middleware.Authentication(), handler.UserRegister)
    userRoutes.DELETE("/", middleware.Authentication(), handler.UserLogin)
  }

  photoRoutes := r.Group("/photos")
  photoRoutes.Use(middleware.Authentication())
  {
    photoRoutes.GET("/", handler.GetAllPhotos)
    photoRoutes.POST("/", handler.UploadPhoto)
    photoRoutes.GET("/:id", handler.GetPhoto)
    photoRoutes.PUT("/:id", handler.UpdatePhoto)
    photoRoutes.DELETE("/:id", handler.DeletePhoto)
  }

	commentRoutes := r.Group("/comments")
  commentRoutes.Use(middleware.Authentication())
	{
		commentRoutes.GET("/", handler.GetAllComments)
		commentRoutes.POST("/", handler.CreateComment)
		commentRoutes.PUT("/:comment_id", handler.UpdateComment)
		commentRoutes.DELETE("/:comment_id", handler.DeleteComment)
	}

	socmedRoutes := r.Group("/socmed")
  socmedRoutes.Use(middleware.Authentication())
	{
		socmedRoutes.GET("/", handler.GetAllSocmed)
		socmedRoutes.POST("/", handler.CreateSocmed)
		socmedRoutes.PUT("/:id", handler.UpdateSocmed)
		socmedRoutes.DELETE("/:id", handler.DeleteSocmed)
	}

  return r
}
