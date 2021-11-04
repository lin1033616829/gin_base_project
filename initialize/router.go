package initialize

import (
	"github.com/gin-gonic/gin"
	"mmfile/middleware"
	"mmfile/router"
)

var Router = gin.Default()

// 初始化总路由

func InitRouter() *gin.Engine {


	PrivateGroup := Router.Group("")
	PrivateGroup.Use(middleware.RateMiddleware())
	{
		router.InitMindRightsRouter(PrivateGroup)                   // 注册功能api路由
	}

	return Router
}
