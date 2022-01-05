package position_reward

import (
	"gin-vue-admin/middleware"

	"gin-vue-admin/api/v1/internal_system/position_reward"

	"github.com/gin-gonic/gin"
)

func InitExchangeTradingfeeRewardRouter(Router *gin.RouterGroup) {
	ExchangeTradingfeeRewardRouter := Router.Group("exchangeTradingfeeReward").Use(middleware.OperationRecord())
	{
		ExchangeTradingfeeRewardRouter.POST("createExchangeTradingfeeReward", position_reward.CreateExchangeTradingfeeReward)             // 新建ExchangeTradingfeeReward
		ExchangeTradingfeeRewardRouter.DELETE("deleteExchangeTradingfeeReward", position_reward.DeleteExchangeTradingfeeReward)           // 删除ExchangeTradingfeeReward
		ExchangeTradingfeeRewardRouter.DELETE("deleteExchangeTradingfeeRewardByIds", position_reward.DeleteExchangeTradingfeeRewardByIds) // 批量删除ExchangeTradingfeeReward
		ExchangeTradingfeeRewardRouter.PUT("updateExchangeTradingfeeReward", position_reward.UpdateExchangeTradingfeeReward)              // 更新ExchangeTradingfeeReward
		ExchangeTradingfeeRewardRouter.GET("findExchangeTradingfeeReward", position_reward.FindExchangeTradingfeeReward)                  // 根据ID获取ExchangeTradingfeeReward
		ExchangeTradingfeeRewardRouter.GET("getExchangeTradingfeeRewardList", position_reward.GetExchangeTradingfeeRewardList)            // 获取ExchangeTradingfeeReward列表
	}
}
