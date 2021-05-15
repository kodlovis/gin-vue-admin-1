package rms

import (
	"gin-vue-admin/api/v1/rms"
	"github.com/gin-gonic/gin"
)

func InitExternalRiskSettingRouter(Router *gin.RouterGroup) {
	//RiskSettingRouter := Router.Group("ExternalRiskSetting").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler()).Use(middleware.OperationRecord())
	ExternalRiskSettingRouter := Router.Group("ExternalRiskSetting")
	{
		ExternalRiskSettingRouter.GET("getRSIList", rms.GetRSIList)  // 获取RiskSetting列表
	}
}
