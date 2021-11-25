package position_reward

import (
	"gin-vue-admin/api/v1/internal_system/position_reward"
	"gin-vue-admin/middleware"

	"github.com/gin-gonic/gin"
)

func InitAccountPositionDailyRouter(Router *gin.RouterGroup) {
	AccountPositionDailyRouter := Router.Group("accountPositionDaily").Use(middleware.OperationRecord())
	{
		AccountPositionDailyRouter.POST("createAccountPositionDaily", position_reward.CreateAccountPositionDaily)             // 新建AccountPositionDaily
		AccountPositionDailyRouter.DELETE("deleteAccountPositionDaily", position_reward.DeleteAccountPositionDaily)           // 删除AccountPositionDaily
		AccountPositionDailyRouter.DELETE("deleteAccountPositionDailyByIds", position_reward.DeleteAccountPositionDailyByIds) // 批量删除AccountPositionDaily
		AccountPositionDailyRouter.PUT("updateAccountPositionDaily", position_reward.UpdateAccountPositionDaily)              // 更新AccountPositionDaily
		AccountPositionDailyRouter.GET("findAccountPositionDaily", position_reward.FindAccountPositionDaily)                  // 根据ID获取AccountPositionDaily
		AccountPositionDailyRouter.GET("getAccountPositionDailyList", position_reward.GetAccountPositionDailyList)            // 获取AccountPositionDaily列表

		AccountPositionDailyRouter.GET("loadAccountPositionExcel", position_reward.LoadAccountPositionExcelData)
	}
}
