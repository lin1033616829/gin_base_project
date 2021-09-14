package router

import (
	"github.com/gin-gonic/gin"
	v1 "mmrights/api/v1"
)

func InitMindRightsRouter(Router *gin.RouterGroup) {
	ApiRouter := Router.Group("mind_rights").Use()
	{
		ApiRouter.POST("createApi", v1.CheckMind)   // 创建Api
	}
}