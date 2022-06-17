package rms

import (
	"gin-vue-admin/api/v1/rms"
	"github.com/gin-gonic/gin"
)

func InitExternalUsersRouter(Router *gin.RouterGroup) {
	ExternalUsersRouter := Router.Group("ExternalUsers")
	{
		ExternalUsersRouter.GET("getUIList", rms.GetUIList)  // 获取Users列表
		ExternalUsersRouter.GET("getUIListByServer", rms.GetUIListByServer)  // 获取Users列表
		
	}
}
