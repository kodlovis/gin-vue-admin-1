package rms

import (
	"gin-vue-admin/api/v1/rms"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitMacRouter(Router *gin.RouterGroup) {
	MacRouter := Router.Group("Mac").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler()).Use(middleware.OperationRecord())
	{
		MacRouter.POST("createMac", rms.CreateMac)   // 新建Mac
		MacRouter.DELETE("deleteMac", rms.DeleteMac) // 删除Mac
		MacRouter.DELETE("deleteMacByIds", rms.DeleteMacByIds) // 批量删除Mac
		MacRouter.POST("updateMac", rms.UpdateMac)    // 更新Mac
		MacRouter.GET("findMac", rms.FindMac)        // 根据ID获取Mac
		MacRouter.GET("getMacList", rms.GetMacList)  // 获取Mac列表
	}
}
