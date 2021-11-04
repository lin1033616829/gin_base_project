package initialize

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"mmfile/global"
)

func RunServer() {
	fmt.Println("初始化工作开始......")

	global.ServerViper = initConf()
	global.ServerLog = InitZap()
	initDb()

	newRouter := gin.Default()
	initRouter(newRouter)
	fmt.Println("初始化工作结束......")
	err := newRouter.Run(fmt.Sprintf(":%v", global.ServerConf.System.Port))
	if err != nil {
		global.ServerLog.Error("routerErr", zap.Error(err))
	}
}