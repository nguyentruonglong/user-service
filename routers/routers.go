package routers

import (
	auth "user_service/auth"
	users "user_service/versioned/v1/users"

	"github.com/gin-gonic/gin"
)

func BindRouters(baseRouter *gin.Engine) {
	userController := new(users.UserController)
	superRouter := baseRouter.Group("api")
	superRouter.Use(gin.Logger())
	superRouter.Use(gin.Recovery())
	superRouter.POST("auth/register", userController.CreateUser)
	superRouter.POST("auth/login", auth.AuthGenerateJWTMiddleware())

	superRouter.Use(auth.AuthValidateTokenMiddleware())
	// superRouter.GET("/auth/logout", auth.AuthLogoutMiddleware())
	// superRouter.POST("/auth/refresh", auth.AuthRefreshTokenMiddleware())
	users.UsersRegister(superRouter)
}
