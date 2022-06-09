package dce

import (
	"gin-vue-admin/middleware"

	v1 "gin-vue-admin/api/v1/internal_system/position_reward/dce"

	"github.com/gin-gonic/gin"
)

func InitUs001DceBrokerPositionDailyRouter(Router *gin.RouterGroup) {
	Us001DceBrokerPositionDailyRouter := Router.Group("us001DceBrokerPositionDaily").Use(middleware.OperationRecord())
	{
		Us001DceBrokerPositionDailyRouter.POST("createUs001DceBrokerPositionDaily", v1.CreateUs001DceBrokerPositionDaily)             // 新建Us001DceBrokerPositionDaily
		Us001DceBrokerPositionDailyRouter.DELETE("deleteUs001DceBrokerPositionDaily", v1.DeleteUs001DceBrokerPositionDaily)           // 删除Us001DceBrokerPositionDaily
		Us001DceBrokerPositionDailyRouter.DELETE("deleteUs001DceBrokerPositionDailyByIds", v1.DeleteUs001DceBrokerPositionDailyByIds) // 批量删除Us001DceBrokerPositionDaily
		Us001DceBrokerPositionDailyRouter.PUT("updateUs001DceBrokerPositionDaily", v1.UpdateUs001DceBrokerPositionDaily)              // 更新Us001DceBrokerPositionDaily
		Us001DceBrokerPositionDailyRouter.GET("findUs001DceBrokerPositionDaily", v1.FindUs001DceBrokerPositionDaily)                  // 根据ID获取Us001DceBrokerPositionDaily
		Us001DceBrokerPositionDailyRouter.GET("getUs001DceBrokerPositionDailyList", v1.GetUs001DceBrokerPositionDailyList)            // 获取Us001DceBrokerPositionDaily列表
		Us001DceBrokerPositionDailyRouter.GET("loadUs001DceBrokerPositionExcel", v1.LoadUs001DceBrokerPositionExcelData)
	}
}
