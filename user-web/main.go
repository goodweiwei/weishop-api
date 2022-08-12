package main

import (
	"fmt"
	"github.com/gin-gonic/gin/binding"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/satori/go.uuid"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"
	"weishop-api/user-web/global"
	"weishop-api/user-web/initialize"
	"weishop-api/user-web/utils/register/consul"
	myvalidator "weishop-api/user-web/validator"
)

func main() {
	//1、初始化配置文件

	//2.初始化logger
	if err := initialize.InitLogger(global.ServerConfig); err != nil {
		panic(err)
	}
	initialize.InitConfig()
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

	//服务注册
	registerClient := consul.NewRegistryClient(global.ServerConfig.ConsulInfo.Host, global.ServerConfig.ConsulInfo.Port)
	serviceId := fmt.Sprintf("%s", uuid.NewV4())
	err := registerClient.Register(global.ServerConfig.Host, global.ServerConfig.Port, global.ServerConfig.Name, global.ServerConfig.Tags, serviceId)
	if err != nil {
		zap.S().Panic("服务注册失败：", err.Error())
	}

	zap.S().Debugf("启动服务器，端口：%d", global.ServerConfig.Port)
	if err := Router.Run(fmt.Sprintf(":%d", global.ServerConfig.Port)); err != nil {
		zap.S().Panic("启动失败", err.Error())
	}

	//接收终止信号
	quit := make(chan os.Signal)
	done := make(chan bool, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-quit
		if err:= registerClient.DeRegister(serviceId);err != nil{
			zap.S().Info("注销失败：", err.Error())
		}else {
			done <- true
			zap.S().Info("注销成功")
		}
	}()
	fmt.Println(serviceId)
	<-done
	fmt.Println("进程被终止")
}
