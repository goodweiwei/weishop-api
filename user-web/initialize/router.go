package initialize

import (
	"github.com/gin-gonic/gin"
	"mxshop-api/user-web/router"
	"net/http"
)

func Routers() *gin.Engine {
	gin.SetMode("debug")
	Router := gin.Default()
	Router.Use(GinLogger(),GinRecovery(true))
	Router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"success": true,
		})
	})

	ApiGroup := Router.Group("/u/v1")
	router.InitUserRouter(ApiGroup)
	router.InitBaseRouter(ApiGroup)
	return Router
}
