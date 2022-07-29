package fund

import (
	v1 "gin-vue-admin/api/v1/internal_system/fund"
	"gin-vue-admin/middleware"

	"github.com/gin-gonic/gin"
)

func InitFundConfirmRouter(Router *gin.RouterGroup) {
	FundConfirmRouter := Router.Group("fundConfirmDate").Use(middleware.OperationRecord())
	{
		FundConfirmRouter.POST("createFundConfirm", v1.CreateFundConfirm)             // 新建FundConfirm
		FundConfirmRouter.DELETE("deleteFundConfirm", v1.DeleteFundConfirm)           // 删除FundConfirm
		FundConfirmRouter.DELETE("deleteFundConfirmByIds", v1.DeleteFundConfirmByIds) // 批量删除FundConfirm
		FundConfirmRouter.PUT("updateFundConfirm", v1.UpdateFundConfirm)              // 更新FundConfirm
		FundConfirmRouter.GET("findFundConfirm", v1.FindFundConfirm)                  // 根据ID获取FundConfirm
		FundConfirmRouter.GET("getFundConfirmList", v1.GetFundConfirmList)            // 获取FundConfirm列表
		FundConfirmRouter.GET("getLastConfirmDate", v1.GetLastConfirmDate)            // 获取当日最近预算日期
		FundConfirmRouter.POST("getCycleConfirmDate", v1.GetCycleLast)                // 根据选中日期获取预算终止日期

		FundConfirmRouter.POST("getConfirmSetting", v1.GetConfirmSetting)  // 根据选中天数获取是否允许修改的配置信息
		FundConfirmRouter.PUT("updateConfirmSetting", v1.UpConfirmSetting) // 更改该天的是否允许修改的配置

	}
}
