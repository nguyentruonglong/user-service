package routers

import (
	auth "user_service/auth"
	users "user_service/versioned/v1/users"

	"github.com/gin-gonic/gin"
)

func BindRouters(baseRouter *gin.Engine) {
	superRouter := baseRouter.Group("api")
	superRouter.Use(gin.Logger())
	superRouter.Use(gin.Recovery())
	superRouter.POST("token/", auth.AuthGenerateJWTMiddleware())
	superRouter.Use(auth.AuthValidateTokenMiddleware())
	users.UsersRegister(superRouter)
}
