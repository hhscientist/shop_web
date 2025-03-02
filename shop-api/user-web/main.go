package main

import (
	"fmt"
	"github.com/gin-gonic/gin/binding"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"shop-api/user-web/global"
	"shop-api/user-web/initialize"
	"shop-api/user-web/utils"
	myvalidator "shop-api/user-web/validator"
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
	////如果是本地开发环境端口号固定，线上环境启动获取端口号
	//debug := viper.GetBool("SHOP_DEBUG")
	//if !debug {
	//	port, err := utils.GetFreePort()
	//	if err != nil {
	//		panic(err)
	//	}
	//	global.ServerConfig.Port = port
	//}

	//端口获取，从conf.yaml里面获取
	viper.SetConfigName("conf")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./user-web/config")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	debug := viper.GetBool("SHOP_DEBUG")
	if !debug {
		port, err := utils.GetFreePort()
		if err != nil {
			panic(err)
		}
		global.ServerConfig.Port = port
	}

	//7注册验证器
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		_ = v.RegisterValidation("mobile", myvalidator.ValidateMobile)
		_ = v.RegisterTranslation("mobile", global.Trans, func(ut ut.Translator) error {
			return ut.Add("mobile", "{0} 非法的手机号码!", true) // see universal-translator for details
		}, func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("mobile", fe.Field())
			return t
		})
	}

	//运行
	zap.S().Infof("启动服务器,端口:%d", global.ServerConfig.Port)
	if err := Router.Run(fmt.Sprintf(":%d", global.ServerConfig.Port)); err != nil {
		zap.S().Panic("启动失败", err.Error())
	}
}
