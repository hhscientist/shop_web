package main

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"shop-api/order-web/global"
	"shop-api/order-web/initialize"
	"shop-api/order-web/utils/register/consul"
	"syscall"
)

func main() {
	//1.初始化logger
	initialize.InitLogger()
	//2.初始化配置文件
	initialize.InitConfig()
	//3.初始化routers
	Router := initialize.Router()
	//4. 初始化翻译
	if err := initialize.InitTrans("zh"); err != nil {
		panic(err)
	}
	////5. 初始化srv的连接
	initialize.InitSrvConn()

	//6端口获取
	//viper.AutomaticEnv()
	////本地开发环境端口号固定，线上环境启动获取端口号
	//debug := viper.GetBool("SHOP_DEBUG")
	//if !debug {
	//	port, err := utils.GetFreePort()
	//	if err != nil {
	//		panic(err)
	//	}
	//	global.ServerConfig.Port = port
	//}

	//端口获取，从conf.yaml里面获取
	//viper.SetConfigName("conf")
	//viper.SetConfigType("yaml")
	//viper.AddConfigPath("./order-web/config")
	//
	//err := viper.ReadInConfig()
	//if err != nil {
	//	panic(err)
	//}
	//debug := viper.GetBool("SHOP_DEBUG")
	//if !debug {
	//	port, err := utils.GetFreePort()
	//	if err != nil {
	//		panic(err)
	//	}
	//	global.ServerConfig.Port = port
	//	//}

	registerClient := consul.NewRegistryClient(global.ServerConfig.ConsulConfig.Host, global.ServerConfig.ConsulConfig.Port)
	serviceId := uuid.NewV4().String()
	registerClient.Register(global.ServerConfig.Host, global.ServerConfig.Port, global.ServerConfig.Name, global.ServerConfig.Tags, serviceId)

	//运行
	zap.S().Infof("启动服务器,端口:%d", global.ServerConfig.Port)
	if err := Router.Run(fmt.Sprintf("0.0.0.0:%d", global.ServerConfig.Port)); err != nil {
		zap.S().Panic("启动失败", err.Error())
	}

	//接受终止信号
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	if err := registerClient.DeRegister(serviceId); err != nil {
		zap.S().Info("注销失败", err.Error())
	} else {
		zap.S().Info("注销成功")

	}
}
