package initialize

import (
	"github.com/gin-gonic/gin"
	"shop-api/user-web/middlewares"
	router2 "shop-api/user-web/router"
)

func Router() *gin.Engine {
	Router := gin.Default()
	Router.Use(middlewares.Cors())
	ApiGroup := Router.Group("/u/v1")
	router2.InitUserRouter(ApiGroup)
	router2.InitBaseRouter(ApiGroup)
	return Router
}
