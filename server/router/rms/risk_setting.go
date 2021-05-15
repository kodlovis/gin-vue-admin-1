package rms

import (
	"gin-vue-admin/api/v1/rms"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitRiskSettingRouter(Router *gin.RouterGroup) {
	RiskSettingRouter := Router.Group("RiskSetting").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler()).Use(middleware.OperationRecord())
	{
		RiskSettingRouter.POST("createRiskSetting", rms.CreateRiskSetting)   // 新建RiskSetting
		RiskSettingRouter.DELETE("deleteRiskSetting", rms.DeleteRiskSetting) // 删除RiskSetting
		RiskSettingRouter.DELETE("deleteRiskSettingByIds", rms.DeleteRiskSettingByIds) // 批量删除RiskSetting
		RiskSettingRouter.PUT("updateRiskSetting", rms.UpdateRiskSetting)    // 更新RiskSetting
		RiskSettingRouter.GET("findRiskSetting", rms.FindRiskSetting)        // 根据ID获取RiskSetting
		RiskSettingRouter.GET("getRiskSettingList", rms.GetRiskSettingList)  // 获取RiskSetting列表
	}
}
