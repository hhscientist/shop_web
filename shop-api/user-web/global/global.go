package global

import (
	ut "github.com/go-playground/universal-translator"
	"shop-api/user-web/config"
	"shop-api/user-web/proto"
)

var (
	Trans         ut.Translator
	ServerConfig  *config.ServerConfig = &config.ServerConfig{}
	UserSrvClient proto.UserClient
	NacosConfig   *config.NacosConfig = &config.NacosConfig{}
)
