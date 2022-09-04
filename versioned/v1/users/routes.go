package users

import (
	"github.com/gin-gonic/gin"
)

func UsersRegister(baseRouter *gin.RouterGroup) *gin.RouterGroup {

	v1 := baseRouter.Group("v1")
	{
		userGroup := v1.Group("user")
		{
			userController := new(UserController)
			userGroup.GET("/", userController.GetUsers)
			userGroup.GET("/:id", userController.GetUser)
			userGroup.POST("/", userController.CreateUser)
		}
	}

	return v1
}
