package initialize

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	_ "github.com/mbobakov/grpc-consul-resolver"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"shop-api/user-web/global"
	"shop-api/user-web/proto"
)

func InitSrvConn() {
	consulInfo := global.ServerConfig.ConsulConfig

	userConn, err := grpc.Dial(
		fmt.Sprintf("consul://%s:%d/%s?wait=14s", consulInfo.Host, consulInfo.Port, global.ServerConfig.UserSrvInfo.Name),
		grpc.WithInsecure(),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
	)
	if err != nil {
		zap.S().Fatal("[InitSrvConn] 连接 【用户服务失败】")
	}

	userSrvClient := proto.NewUserClient(userConn)
	global.UserSrvClient = userSrvClient
}

func InitSrvConn2() {
	//consulInfo := global.ServerConfig.ConsulConfig
	//userConn, err := grpc.Dial(
	//	fmt.Sprintf("%s:%d/%s?wait=14s", consulInfo.Host, consulInfo.Port, global.ServerConfig.UserSrvInfo.Name),
	//	grpc.WithInsecure(),
	//	grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
	//)
	//
	//if err != nil {
	//	zap.S().Fatal("[InitSrvConn] 连接 【用户服务失败】")
	//}
	//
	//userSrvClient := proto.NewUserClient(userConn)
	//global.UserSrvClient = userSrvClient

	//创建 Consul 客户端
	cfg := api.DefaultConfig()
	cfg.Address = fmt.Sprintf("%s:%d", global.ServerConfig.ConsulConfig.Host, global.ServerConfig.ConsulConfig.Port)
	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}

	//通过 Consul 查询 user-srv
	//data, err := client.Agent().ServicesWithFilter(`Service == "user-srv"`)
	data, err := client.Agent().ServicesWithFilter(fmt.Sprintf(`Service == "%s"`, global.ServerConfig.UserSrvInfo.Name))
	if err != nil {
		panic(err)
	}
	if len(data) == 0 {
		zap.S().Fatal("没有找到名为 user-srv 的服务实例")
	}

	//获取服务地址和端口
	addr, port := "", 0
	for _, value := range data {
		addr = value.Address
		port = value.Port
		break
	}

	zap.S().Info("服务端ip:", addr)
	zap.S().Info("服务端:", port)

	//拨号连接用户grpc服务
	userConn, err := grpc.Dial(fmt.Sprintf("%s:%d", addr, port), grpc.WithInsecure())
	if err != nil {
		zap.S().Errorw("[GetUserList]连接[用户服务失败]",
			"msg", err.Error())
	}

	//调用接口 创建grpc客户端
	userSrvClient := proto.NewUserClient(userConn)
	global.UserSrvClient = userSrvClient

}
