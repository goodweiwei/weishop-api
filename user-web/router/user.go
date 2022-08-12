package router

import (
	"github.com/gin-gonic/gin"
	"weishop-api/user-web/api"
	"weishop-api/user-web/middlewares"
)

func InitUserRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("user")
	{
		UserRouter.GET("list", middlewares.JWTAuth(), middlewares.IsAdminAuth(),api.GetUserList)
		UserRouter.GET("/id/:id",middlewares.JWTAuth(), middlewares.IsAdminAuth(), api.GetUserById)
		UserRouter.POST("/login", api.PasswordLogin)
		UserRouter.POST("/register", api.RegisterUser)
	}
}
