package main

import (
	"fmt"
	"github.com/gin-gonic/gin/binding"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"weishop-api/user-web/global"
	"weishop-api/user-web/initialize"
	myvalidator "weishop-api/user-web/validator"
)

func main() {
	initialize.InitConfig()
	//1.初始化logger
	if err:=initialize.InitLogger(global.ServerConfig);err!=nil{
		panic(err)
	}
	//2.初始化配置文件

	//3.初始化routers
	Router := initialize.Routers()
	//4。初始化翻译器
	if err := initialize.InitTrans(); err != nil {
		panic(err)
	}
	//5.初始化srv的连接
	initialize.InitSrvConn()

	//注册验证器
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		_ = v.RegisterValidation("mobile", myvalidator.ValidateMobile)
		_ = v.RegisterTranslation("mobile", global.Trans, func(ut ut.Translator) error {
			return ut.Add("mobile", "{0} 非法的手机号码!", true) // see universal-translator for details
		}, func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("mobile", fe.Field())
			return t
		})
	}

	zap.S().Debugf("启动服务器，端口：%d", global.ServerConfig.Port)
	if err := Router.Run(fmt.Sprintf(":%d", global.ServerConfig.Port)); err != nil {
		zap.S().Panic("启动失败", err.Error())
	}

}
