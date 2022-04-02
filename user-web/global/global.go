package global

import (
	ut "github.com/go-playground/universal-translator"
	"weishop-api/user-web/config"
	"weishop-api/user-web/proto"
	"github.com/go-playground/validator/v10"
)

var (
	//Trans ut.Translator

	NacosConfig *config.NacosConfig = &config.NacosConfig{}
	ServerConfig *config.ServerConfig = &config.ServerConfig{}
	UserSrvClient proto.UserClient
	//LogConfig *config.LogConfig = &config.LogConfig{}
)

var (
	Uni      *ut.UniversalTranslator
	Validate *validator.Validate
	Trans ut.Translator
)