package router

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"shop-api/goods-web/api/goods"
	"shop-api/user-web/middlewares"
)

func InitGoodsRouter(Router *gin.RouterGroup) {
	GoodsRouter := Router.Group("goods")

	zap.S().Infof("设置用户相关的url")
	{
		GoodsRouter.GET("", goods.List)
		GoodsRouter.POST("", middlewares.JWTAuth(), middlewares.IsAdminAuth(), goods.New)
		GoodsRouter.GET("/:id", goods.Detail)
		GoodsRouter.DELETE("/:id", middlewares.JWTAuth(), middlewares.IsAdminAuth(), goods.Delete)
		GoodsRouter.GET("/:id/stocks", goods.Stocks)
		GoodsRouter.PATCH("/:id", goods.Update)
	}
}
