package main

import (
	"flag"
	"fmt"
	"github.com/hashicorp/consul/api"
	uuid "github.com/satori/go.uuid"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"net"
	"os"
	"os/signal"
	"shop_srvs/user_srv/global"
	"shop_srvs/user_srv/handler"
	"shop_srvs/user_srv/initialize"
	"shop_srvs/user_srv/proto"
	"shop_srvs/user_srv/utils"
	"syscall"
)

func main() {
	//参数绑定
	IP := flag.String("ip", "0.0.0.0", "ip地址")
	Port := flag.Int("port", 0, "端口号")

	initialize.InitLogger()
	initialize.InitConfig()
	initialize.InitDB()

	//zap.S().Info("ServerConfig的信息", global.ServerConfig)
	flag.Parse()
	zap.S().Info("ip: ", *IP)
	//分配端口
	if *Port == 0 {
		*Port, _ = utils.GetFreePort()
	}
	zap.S().Info("port: ", *Port)

	// 创建一个 gRPC 服务器
	server := grpc.NewServer()
	// 注册服务
	proto.RegisterUserServer(server, &handler.UserServer{})

	// 创建一个监听端口的 listener
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", *IP, *Port))
	if err != nil {
		panic("failed to listen:" + err.Error())
	}

	//注册健康检查
	grpc_health_v1.RegisterHealthServer(server, health.NewServer())

	//在consul里面注册服务
	cfg := api.DefaultConfig()
	cfg.Address = fmt.Sprintf("http://%s:%d", global.ServerConfig.ConsulInfo.Host, global.ServerConfig.ConsulInfo.Port)

	client, err := api.NewClient(cfg)
	if err != nil {
		panic("consul连接失败")
	}

	//生成对应的检查对象
	check := &api.AgentServiceCheck{
		GRPC:                           fmt.Sprintf("%s:%d", "172.30.64.1", *Port),
		Timeout:                        "5s",
		Interval:                       "10s",
		DeregisterCriticalServiceAfter: "3600s",
	}

	registration := new(api.AgentServiceRegistration)
	registration.Name = global.ServerConfig.Name
	serverID := fmt.Sprintf("%s", uuid.NewV4())
	registration.ID = serverID
	registration.Port = *Port
	registration.Tags = []string{"tmp", "bobby", "user", "srv"}
	registration.Address = "172.30.64.1"
	//registration.Address = "10.205.6.4"
	registration.Check = check

	err = client.Agent().ServiceRegister(registration)
	if err != nil {
		panic(err)
	}

	//启动grpc服务
	go func() {
		err = server.Serve(lis)
		if err != nil {
			panic("failed to start grpc:" + err.Error())
		}
	}()

	//接收终止信号
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	if err = client.Agent().ServiceDeregister(serverID); err != nil {
		zap.S().Info("注销失败")
	}
	zap.S().Info("注销成功")
}
