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
		// blogRouter := apiRouter.Group("/blogs")
		// {
		// 	blogRouter.GET("/", controller.GetBlogList)
		// }

		// headerRouter := apiRouter.Group("/headers")
		// {
		// 	headerRouter.GET("/", controller.GetHeaderList)
		// }

		authRouter := apiRouter.Group("/auths")
		{
			// authRouter.POST("/register", controller.RegisterUser)
			authRouter.POST("/login", controller.LoginUser)
		}

		cmsRouter := apiRouter.Group("/admin")
		{
			userRouter := cmsRouter.Group("/users")
			{
				userRouter.Use(middleware.Authentication(), middleware.Authorization("userId"))
				// userRouter.GET("/", controller.GetProfile)
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
	}

	return r
}
