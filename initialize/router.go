package initialize

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mmfile/middleware"
	"mmfile/router"
)


// 初始化总路由

func InitRouter(Router *gin.Engine) *gin.Engine {
	fmt.Println("开始注册路由")
	PrivateGroup := Router.Group("")
	PrivateGroup.Use(middleware.RateMiddleware())
	{
		router.InitMindRightsRouter(PrivateGroup)                   // 注册功能api路由
	}

	return Router
}
