package router

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"shop-api/goods-web/api/goods"
)

func InitUserRouter(Router *gin.RouterGroup) {
	GoodsRouter := Router.Group("goods")

	zap.S().Infof("设置用户相关的url")
	{
		GoodsRouter.GET("", goods.List)
		GoodsRouter.POST("", goods.New)
	}
}
