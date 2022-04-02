package router

import (
	"github.com/gin-gonic/gin"
	"weishop-api/user-web/api"
)

func InitUserRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("user")
	{
		UserRouter.GET("list", api.GetUserList)
		UserRouter.GET("/id/:id", api.GetUserById)
		UserRouter.POST("/login", api.PasswordLogin)
	}
}
