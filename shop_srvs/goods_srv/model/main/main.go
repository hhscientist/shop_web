package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"io"
	"log"
	"math/rand"
	"os"
	"shop_srvs/goods_srv/model"
	"time"
)

func genMd5(code string) string {
	Md5 := md5.New()
	_, _ = io.WriteString(Md5, code)
	return hex.EncodeToString(Md5.Sum(nil))
}

// 随机生成手机号
func generateMobile() string {
	prefixes := []string{
		"133", "149", "153", "173", "177", "180", "181", "189", // 运营商号段
		"130", "131", "132", "155", "156", "166", "175", "176", // 其他号段
		"185", "186", "187", "188", "199",
	}
	rand.Seed(time.Now().UnixNano()) // 设定随机种子
	prefix := prefixes[rand.Intn(len(prefixes))]
	suffix := fmt.Sprintf("%08d", rand.Intn(100000000)) // 生成 8 位随机数
	return prefix + suffix
}

func main() {
	dsn := "root:Aa123456@tcp(172.30.70.200:3306)/shop_goods_srv?charset=utf8mb4&parseTime=True&loc=Local"

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // 慢 SQL 阈值
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // 禁用彩色打印
		},
	)

	//全局模式
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}
	_ = db.AutoMigrate(&model.Category{},
		&model.Brands{}, &model.GoodsCategoryBrand{}, &model.Banner{}, &model.Goods{})

}
