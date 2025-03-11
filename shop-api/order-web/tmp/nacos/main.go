package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	// 添加预检：测试 TCP 连接
	conn, err := net.DialTimeout("tcp", "172.30.70.200:8848", 5*time.Second)
	if err != nil {
		panic(fmt.Sprintf("无法连接到 Nacos 服务器: %v", err))
	}
	defer conn.Close()
	fmt.Println("Nacos 服务器连接正常")

	// 原有代码...
}
