package initialize

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mmfile/global"
)

func init() {
	fmt.Println("初始化工作开始......")

	global.ServerViper = InitConf()
	global.ServerLog = InitZap()
	InitDb()
	newRouter := gin.Default()
	InitRouter(newRouter)
	fmt.Println("初始化工作结束......")
	newRouter.Run(global.ServerConf.System.Port)
}