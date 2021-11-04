package initialize

import (
	"fmt"
	"mmfile/global"
)

func init() {
	fmt.Println("初始化工作开始......")
	global.ServerViper = InitConf()
	global.ServerLog = InitZap()
	InitRouter()
	InitDb()
	fmt.Println("初始化工作结束......")
}