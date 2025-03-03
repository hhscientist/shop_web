package initialize

import (
	"github.com/gin-gonic/gin"
	"shop-api/goods-web/middlewares"
	router2 "shop-api/goods-web/router"
)

func Router() *gin.Engine {
	Router := gin.Default()
	Router.Use(middlewares.Cors())
	Router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	ApiGroup := Router.Group("/g/v1")
	router2.InitUserRouter(ApiGroup)
	router2.InitBaseRouter(ApiGroup)
	return Router
}
