package czce

import (
	"gin-vue-admin/middleware"

	v1 "gin-vue-admin/api/v1/internal_system/position_reward/czce"

	"github.com/gin-gonic/gin"
)

func InitUs001CzceBrokerPositionDailyRouter(Router *gin.RouterGroup) {
	Us001CzceBrokerPositionDailyRouter := Router.Group("us001CzceBrokerPositionDaily").Use(middleware.OperationRecord())
	{
		Us001CzceBrokerPositionDailyRouter.POST("createUs001CzceBrokerPositionDaily", v1.CreateUs001CzceBrokerPositionDaily)             // 新建Us001CzceBrokerPositionDaily
		Us001CzceBrokerPositionDailyRouter.DELETE("deleteUs001CzceBrokerPositionDaily", v1.DeleteUs001CzceBrokerPositionDaily)           // 删除Us001CzceBrokerPositionDaily
		Us001CzceBrokerPositionDailyRouter.DELETE("deleteUs001CzceBrokerPositionDailyByIds", v1.DeleteUs001CzceBrokerPositionDailyByIds) // 批量删除Us001CzceBrokerPositionDaily
		Us001CzceBrokerPositionDailyRouter.PUT("updateUs001CzceBrokerPositionDaily", v1.UpdateUs001CzceBrokerPositionDaily)              // 更新Us001CzceBrokerPositionDaily
		Us001CzceBrokerPositionDailyRouter.GET("findUs001CzceBrokerPositionDaily", v1.FindUs001CzceBrokerPositionDaily)                  // 根据ID获取Us001CzceBrokerPositionDaily
		Us001CzceBrokerPositionDailyRouter.GET("getUs001CzceBrokerPositionDailyList", v1.GetUs001CzceBrokerPositionDailyList)            // 获取Us001CzceBrokerPositionDaily列表
		Us001CzceBrokerPositionDailyRouter.GET("loadUs001CzceBrokerPositionExcel", v1.LoadUs001CzceBrokerPositionExcelData)
	}
}
