package api

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"math/rand"
	"shop-api/user-web/forms"
	"shop-api/user-web/global"
	"strings"
	"time"
)

func GenerateSmsCode(width int) string {
	numeric := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	r := len(numeric)
	//rand.Seed(time.Now().UnixNano())

	var sb strings.Builder
	for i := 0; i < width; i++ {
		fmt.Fprintf(&sb, "%d", numeric[rand.Intn(r)])
	}
	return sb.String()
}

func SendSms(c *gin.Context) {
	sendSmsForm := forms.SendSmsForm{}
	if err := c.ShouldBind(&sendSmsForm); err != nil {
		HandleValidatorError(c, err)
		return
	}

	mobile := "13876544567"

	smsCode := GenerateSmsCode(6)

	fmt.Println("短信验证码", smsCode)
	fmt.Println("认证信息", fmt.Sprintf("%s:%d", global.ServerConfig.RedisInfo.Host, global.ServerConfig.RedisInfo.Port))
	fmt.Println("host：", global.ServerConfig.RedisInfo.Host)
	fmt.Println("port: ", global.ServerConfig.RedisInfo.Port)
	//redis初始化
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", global.ServerConfig.RedisInfo.Host, global.ServerConfig.RedisInfo.Port),
		Password: "", // 没有密码，默认值
		DB:       0,  // 默认DB 0
	})

	rdb.Set(context.Background(), mobile, smsCode, time.Duration(global.ServerConfig.RedisInfo.Expire)*time.Second)
	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		fmt.Println("连接Redis失败:", err)
		c.JSON(500, gin.H{
			"msg": "连接Redis失败",
		})
		return
	}
	c.JSON(200, gin.H{
		"msg":      "短信发送成功",
		"mobile":   mobile,
		"sms_code": smsCode,
	})

}
