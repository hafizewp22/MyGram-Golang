package router

import (
	"project_final/controllers"
	"project_final/middlewares"

	"github.com/gin-gonic/gin"
)

func New() *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("users")
	{
		userRouter.POST("/register", controllers.RegisterUser)
		userRouter.POST("/login", controllers.LoginUser)
	}

	socialMediaRouter := r.Group("socialmedia")
	{
		socialMediaRouter.Use(middlewares.Authorization())
		socialMediaRouter.POST("/", controllers.CreateSocialMedia)
		socialMediaRouter.GET("/", middlewares.Authorization(), controllers.GetAllSocialMedia)
		socialMediaRouter.GET("/:socialMediaID", middlewares.Authorization(), controllers.GetSocialMedia)
		socialMediaRouter.PUT("/:socialMediaID", middlewares.SosialMediaAuthorization(), controllers.UpdateSocialMedia)
		socialMediaRouter.DELETE("/:socialMediaID", middlewares.SosialMediaAuthorization(), controllers.DeleteSocialMedia)
	}

	photoRouter := r.Group("photo")
	{
		photoRouter.Use(middlewares.Authorization())
		photoRouter.POST("/", controllers.CreatePhoto)
		photoRouter.GET("/", middlewares.Authorization(), controllers.GetAllPhoto)
		photoRouter.GET("/:PhotoID", middlewares.Authorization(), controllers.GetPhoto)
		photoRouter.PUT("/:PhotoID", middlewares.PhotoAuthorization(), controllers.UpdatePhoto)
		photoRouter.DELETE("/:PhotoID", middlewares.PhotoAuthorization(), controllers.DeletePhoto)
	}

	commentRouter := r.Group("comment")
	{
		commentRouter.Use(middlewares.Authorization())
		commentRouter.POST("/", controllers.CreateComment)
		commentRouter.GET("/", middlewares.Authorization(), controllers.GetAllComment)
		commentRouter.GET("/:CommentID", middlewares.Authorization(), controllers.GetComment)
		commentRouter.PUT("/:CommentID", middlewares.CommentAuthorization(), controllers.UpdateComment)
		commentRouter.DELETE("/:CommentID", middlewares.CommentAuthorization(), controllers.DeleteComment)
	}

	return r
}
