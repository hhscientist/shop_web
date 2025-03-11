package initialize

import (
	"encoding/json"
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"shop_srvs/order_srv/global"
)

func GetEnvInfo(env string) string {
	viper.AutomaticEnv()
	value := viper.GetString(env)
	if value == "" {
		return "true"
	}
	return value
}

func InitConfig() {
	data := GetEnvInfo("Debug")
	var configFileName string
	configFileNamePrefix := "config"
	if data == "true" {
		configFileName = fmt.Sprintf("order_srv/%s-debug.yaml", configFileNamePrefix)
	} else {
		configFileName = fmt.Sprintf("order_srv/%s-pro.yaml", configFileNamePrefix)
	}

	v := viper.New()
	v.SetConfigFile(configFileName)
	err := v.ReadInConfig()
	if err != nil {
		panic(err)
	}
	//这个对象如何在其他文件中使用 - 全局变量
	if err = v.Unmarshal(global.NacosConfig); err != nil {
		panic(err)
	}

	zap.S().Infof("配置信息: &v", global.NacosConfig)

	clientConfig := constant.ClientConfig{
		NamespaceId:         global.NacosConfig.Namespace, // 如果需要支持多namespace，我们可以创建多个client,它们有不同的NamespaceId。当namespace是public时，此处填空字符串。
		TimeoutMs:           1000,
		NotLoadCacheAtStart: false,
		LogDir:              "/tmp/nacos/log",
		CacheDir:            "/tmp/nacos/cache",
		LogLevel:            "warn",
		Username:            global.NacosConfig.User,
		Password:            global.NacosConfig.Password,
	}

	serverConfigs := []constant.ServerConfig{
		{
			IpAddr: global.NacosConfig.Host,
			Port:   global.NacosConfig.Port,
		},
	}

	// create config client
	configClient, err := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": serverConfigs,
		"clientConfig":  clientConfig,
	})
	if err != nil {
		panic(err)
	}
	//配置的内容写道context里面
	content, err := configClient.GetConfig(vo.ConfigParam{
		DataId: global.NacosConfig.DataId,
		Group:  global.NacosConfig.Group})

	if err != nil {
		if content == "" {
			zap.S().Infof("获取配置失败%s", err.Error())
		}
		panic(err)
	}
	//fmt.Println(content) //字符串 - yaml
	//想要将一个json字符串转换成struct，需要去设置这个struct的tag
	err = json.Unmarshal([]byte(content), &global.ServerConfig)
	if err != nil {
		zap.S().Fatalf("读取nacos配置失败： %s", err.Error())
	}
	fmt.Println(&global.ServerConfig)

}
