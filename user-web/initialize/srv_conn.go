package initialize

import (
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"weishop-api/user-web/global"
	"weishop-api/user-web/proto"
)

func InitSrvConn()  {
	userConn, err := grpc.Dial(":50051", grpc.WithInsecure())
	if err != nil{
		zap.S().Fatal("[InitSrvConn] 连接 【用户服务失败】")
	}
	userSrcClient := proto.NewUserClient(userConn)
	global.UserSrvClient = userSrcClient
}
