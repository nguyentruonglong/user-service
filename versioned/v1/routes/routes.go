package v1_routes

import (
	"user_service/versioned/v1/controllers"

	"github.com/gin-gonic/gin"
)

func BindRoutes(baseApiRouter *gin.RouterGroup) *gin.RouterGroup {

	v1 := baseApiRouter.Group("v1")
	{
		userGroup := v1.Group("user")
		{
			user := new(controllers.UserController)
			userGroup.GET("/:id", user.Retrieve)
		}
	}

	return v1
}
