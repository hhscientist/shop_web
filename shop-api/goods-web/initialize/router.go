package initialize

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shop-api/goods-web/middlewares"
	"shop-api/goods-web/router"
)

func Router() *gin.Engine {
	Router := gin.Default()
	Router.Use(middlewares.Cors())
	Router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"code":    http.StatusOK,
			"success": true,
		})
	})

	ApiGroup := Router.Group("/g/v1")

	router.InitGoodsRouter(ApiGroup)
	router.InitCategoryRouter(ApiGroup)
	router.InitBannerRouter(ApiGroup)
	router.InitBrandRouter(ApiGroup)

	return Router
}
