package future_data

import (
	"gin-vue-admin/api/v1/internalSystem/future_data"
	"gin-vue-admin/middleware"

	"github.com/gin-gonic/gin"
)

func InitFutureRightsDailyRouter(Router *gin.RouterGroup) {
	FutureRightsDailyRouter := Router.Group("futureRightsDaily").Use(middleware.OperationRecord())
	{
		FutureRightsDailyRouter.POST("createFutureRightsDaily", future_data.CreateFutureRightsDaily)             // 新建FutureRightsDaily
		FutureRightsDailyRouter.DELETE("deleteFutureRightsDaily", future_data.DeleteFutureRightsDaily)           // 删除FutureRightsDaily
		FutureRightsDailyRouter.DELETE("deleteFutureRightsDailyByIds", future_data.DeleteFutureRightsDailyByIds) // 批量删除FutureRightsDaily
		FutureRightsDailyRouter.PUT("updateFutureRightsDaily", future_data.UpdateFutureRightsDaily)              // 更新FutureRightsDaily
		FutureRightsDailyRouter.GET("findFutureRightsDaily", future_data.FindFutureRightsDaily)                  // 根据ID获取FutureRightsDaily
		FutureRightsDailyRouter.GET("getFutureRightsDailyList", future_data.GetFutureRightsDailyList)            // 获取FutureRightsDaily列表
	}
}
