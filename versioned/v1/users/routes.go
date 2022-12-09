package users

import (
	"github.com/gin-gonic/gin"
)

func UsersRegister(baseRouter *gin.RouterGroup) *gin.RouterGroup {

	v1 := baseRouter.Group("v1")
	{
		userGroup := v1.Group("users")
		{
			userController := new(UserController)
			userGroup.GET("/search", userController.GetUsers)
			userGroup.GET("/detail/:id", userController.GetUser)
		}
	}

	return v1
}
