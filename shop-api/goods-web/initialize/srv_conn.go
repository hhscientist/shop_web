package initialize

import (
	"fmt"
	_ "github.com/mbobakov/grpc-consul-resolver"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"shop-api/goods-web/global"
	"shop-api/goods-web/proto"
)

func InitSrvConn() {
	consulInfo := global.ServerConfig.ConsulConfig

	goodsConn, err := grpc.Dial(
		fmt.Sprintf("consul://%s:%d/%s?wait=14s", consulInfo.Host, consulInfo.Port, global.ServerConfig.GoodsSrvInfo.Name),
		grpc.WithInsecure(),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
	)
	if err != nil {
		zap.S().Error(err)
		zap.S().Fatal(fmt.Sprintf("[InitSrvConn] 连接 【用户服务失败】global.ServerConfig.GoodsSrvInfo.Name %s", global.ServerConfig.GoodsSrvInfo.Name))
		zap.S().Fatal("[InitSrvConn] 连接 【用户服务失败】")
	}

	goodsSrvClient := proto.NewGoodsClient(goodsConn)
	global.GoodsSrvclient = goodsSrvClient
}
