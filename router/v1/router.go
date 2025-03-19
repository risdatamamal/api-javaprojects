package router

import (
	"log"

	"github.com/risdatamamal/api-javaprojects/controller"
	"github.com/risdatamamal/api-javaprojects/middleware"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	err := r.SetTrustedProxies(nil)
	if err != nil {
		log.Fatal(err)
	}

	apiRouter := r.Group("/api/v1")
	{
		// blogRouter := apiRouter.Group("/blogs")
		// {
		// 	blogRouter.GET("/", controller.GetBlogList)
		// }

		headerRouter := apiRouter.Group("/header")
		{
			headerRouter.GET("/", controller.GetHeader)
		}

		authRouter := apiRouter.Group("/auth")
		{
			authRouter.POST("/register", controller.RegisterUser)
			authRouter.POST("/login", controller.LoginUser)
		}

		userRouter := apiRouter.Group("/user")
		{
			userRouter.Use(middleware.Authentication())
			userRouter.Use(middleware.AuthMiddleware("Admin")) // only admin can access this route (ERROR)
			userRouter.GET("/get-profile", controller.GetProfile)
			// userRouter.Use(middleware.Authorization("userId"))
			// userRouter.PUT("/update/:userId", controller.UpdateUser)
			// userRouter.DELETE("/delete/:userId", controller.DeleteUser)
		}

		cmsRouter := apiRouter.Group("/admin")
		{
			cmsRouter.POST("/auth/login", controller.LoginAdminUser)

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
