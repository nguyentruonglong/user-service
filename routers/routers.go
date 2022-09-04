package routers

import (
	"user_service/versioned/v1/middlewares"
	users "user_service/versioned/v1/users"

	"github.com/gin-gonic/gin"
)

func BindRouters(baseRouter *gin.Engine) {
	superRouter := baseRouter.Group("api")
	superRouter.Use(gin.Logger())
	superRouter.Use(gin.Recovery())

	superRouter.Use(middlewares.AuthMiddleware())

	users.UsersRegister(superRouter)
}
