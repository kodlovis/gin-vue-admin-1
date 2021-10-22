package position_reward

import (
	"gin-vue-admin/api/v1/internalSystem/position_reward"
	"gin-vue-admin/middleware"

	"github.com/gin-gonic/gin"
)

func InitPositionRewardRuleRouter(Router *gin.RouterGroup) {
	PositionRewardRuleRouter := Router.Group("positionRewardRule").Use(middleware.OperationRecord())
	{
		PositionRewardRuleRouter.POST("createPositionRewardRule", position_reward.CreatePositionRewardRule)             // 新建PositionRewardRule
		PositionRewardRuleRouter.DELETE("deletePositionRewardRule", position_reward.DeletePositionRewardRule)           // 删除PositionRewardRule
		PositionRewardRuleRouter.DELETE("deletePositionRewardRuleByIds", position_reward.DeletePositionRewardRuleByIds) // 批量删除PositionRewardRule
		PositionRewardRuleRouter.PUT("updatePositionRewardRule", position_reward.UpdatePositionRewardRule)              // 更新PositionRewardRule
		PositionRewardRuleRouter.GET("findPositionRewardRule", position_reward.FindPositionRewardRule)                  // 根据ID获取PositionRewardRule
		PositionRewardRuleRouter.GET("getPositionRewardRuleList", position_reward.GetPositionRewardRuleList)            // 获取PositionRewardRule列表
	}
}
