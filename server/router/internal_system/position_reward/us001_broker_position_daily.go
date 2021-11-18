package position_reward

import (
	"gin-vue-admin/api/v1/internal_system/position_reward"
	"gin-vue-admin/middleware"

	"github.com/gin-gonic/gin"
)

func InitBrokerPositionDailyRouter(Router *gin.RouterGroup) {
	BrokerPositionDailyRouter := Router.Group("brokerPositionDaily").Use(middleware.OperationRecord())
	{
		BrokerPositionDailyRouter.POST("createBrokerPositionDaily", position_reward.CreateBrokerPositionDaily)             // 新建BrokerPositionDaily
		BrokerPositionDailyRouter.DELETE("deleteBrokerPositionDaily", position_reward.DeleteBrokerPositionDaily)           // 删除BrokerPositionDaily
		BrokerPositionDailyRouter.DELETE("deleteBrokerPositionDailyByIds", position_reward.DeleteBrokerPositionDailyByIds) // 批量删除BrokerPositionDaily
		BrokerPositionDailyRouter.PUT("updateBrokerPositionDaily", position_reward.UpdateBrokerPositionDaily)              // 更新BrokerPositionDaily
		BrokerPositionDailyRouter.GET("findBrokerPositionDaily", position_reward.FindBrokerPositionDaily)                  // 根据ID获取BrokerPositionDaily
		BrokerPositionDailyRouter.GET("getBrokerPositionDailyList", position_reward.GetBrokerPositionDailyList)            // 获取BrokerPositionDaily列表

		BrokerPositionDailyRouter.GET("loadBrokerPositionExcel", position_reward.LoadBrokerPositionExcelData)
	}
}
