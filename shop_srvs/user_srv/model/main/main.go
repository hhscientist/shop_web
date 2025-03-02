package main

import (
	"crypto/md5"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"github.com/anaskhan96/go-password-encoder"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"io"
	"log"
	"math/rand"
	"os"
	"shop_srvs/user_srv/model"
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
	dsn := "root:Aa123456@tcp(172.30.70.200:3306)/shop_usr_srv?charset=utf8mb4&parseTime=True&loc=Local"

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

	//
	//_ = db.AutoMigrate(&model.User{})

	//fmt.Println(genMd5("112233"))

	//// Using the default options
	//salt, encodedPwd := password.Encode("generic password", nil)
	//fmt.Println(salt, encodedPwd)
	//check := password.Verify("generic password", salt, encodedPwd, nil)
	//fmt.Println(check) // true

	//db 摆着看的
	//db.First(&model.User{})

	// Using custom options
	//加密
	options := &password.Options{16, 10000, 32, sha512.New}
	salt, encodedPwd := password.Encode("admin123", options)
	newPassword := fmt.Sprintf("$pbkdf2-sha512$%s$%s", salt, encodedPwd)
	fmt.Println(newPassword)

	for i := 0; i < 10; i++ {
		user := model.User{
			NickName: fmt.Sprintf("bobby%d", i),
			Mobile:   fmt.Sprintf("1878222222%d", i),
			Password: newPassword,
		}
		db.Save(&user)
	}
	////fmt.Println(newPassword)
	////fmt.Println(len(newPassword))
	//
	////解析密码
	//passwordInfo := strings.Split(newPassword, "$")
	//check := password.Verify("admin", passwordInfo[2], passwordInfo[3], options)
	//fmt.Println(check) // true

}
