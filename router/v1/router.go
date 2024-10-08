package router

import (
	"github.com/risdatamamal/api-javaprojects/controller"
	"github.com/risdatamamal/api-javaprojects/middleware"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	apiRouter := r.Group("/api/v1")
	{
		userRouter := apiRouter.Group("/users")
		{
			userRouter.POST("/register", controller.RegisterUser)
			userRouter.POST("/login", controller.LoginUser)
			userRouter.Use(middleware.Authentication(), middleware.Authorization("userId"))
			// userRouter.PUT("/:userId", controller.UpdateUser)
		}

		// photoRouter := apiRouter.Group("/photos")
		// {
		// 	photoRouter.Use(middleware.Authentication())
		// 	photoRouter.POST("/", controller.PostPhoto)
		// 	photoRouter.GET("/", controller.GetPhotos)
		// 	photoRouter.Use(middleware.Authorization("photoId"))
		// 	photoRouter.PUT("/:photoId", controller.UpdatePhoto)
		// 	photoRouter.DELETE("/:photoId", controller.DeletePhoto)
		// }
	}

	return r
}
