package routers

import (
	"user_service/versioned/v1/middlewares"
	routes_v1 "user_service/versioned/v1/routes"

	"github.com/gin-gonic/gin"
)

func BindRouters(baseRouter *gin.Engine) {
	apiRouterGroup := baseRouter.Group("api")
	apiRouterGroup.Use(gin.Logger())
	apiRouterGroup.Use(gin.Recovery())

	apiRouterGroup.Use(middlewares.AuthMiddleware())

	routes_v1.BindRoutes(apiRouterGroup)
}
