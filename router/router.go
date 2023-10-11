package router

import (
	"github.com/gin-gonic/gin"
	"github.com/mcsans/go-jwt/controllers"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("/users")

	{
		userRouter.POST("/register", controllers.UserRegister)
		userRouter.POST("/login", controllers.UserLogin)
	}

	return r
}