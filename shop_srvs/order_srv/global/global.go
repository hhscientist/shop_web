package global

import (
	goredislib "github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"shop_srvs/order_srv/config"
	"shop_srvs/order_srv/proto"
)

var (
	DB                 *gorm.DB
	ServerConfig       config.ServerConfig
	NacosConfig        *config.NacosConfig = &config.NacosConfig{}
	GoodsSrvClient     proto.GoodsClient
	InventorySrvClient proto.InventoryClient
	RedisClient        *goredislib.Client
)

//func init() {
//
//	dsn := "root:Aa123456@tcp(172.30.70.200:3306)/shop_usr_srv?charset=utf8mb4&parseTime=True&loc=Local"
//
//	newLogger := logger.New(
//		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
//		logger.Config{
//			SlowThreshold: time.Second, // 慢 SQL 阈值
//			LogLevel:      logger.Info, // Log level
//			Colorful:      true,        // 禁用彩色打印
//		},
//	)
//
//	//全局模式
//	var err error
//	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
//		NamingStrategy: schema.NamingStrategy{
//			SingularTable: true,
//		},
//		Logger: newLogger,
//	})
//	if err != nil {
//		panic(err)
//	}
//
//}
