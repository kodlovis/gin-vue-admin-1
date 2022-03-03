package future_inventory

import (
	"gin-vue-admin/api/v1/internal_system/future_inventory"
	"gin-vue-admin/middleware"

	"github.com/gin-gonic/gin"
)

func InitUs004FutureInventoryDailyRouter(Router *gin.RouterGroup) {
	Us004FutureInventoryDailyRouter := Router.Group("us004FutureInventoryDaily").Use(middleware.OperationRecord())
	{
		Us004FutureInventoryDailyRouter.POST("createUs004FutureInventoryDaily", future_inventory.CreateUs004FutureInventoryDaily)             // 新建Us004FutureInventoryDaily
		Us004FutureInventoryDailyRouter.DELETE("deleteUs004FutureInventoryDaily", future_inventory.DeleteUs004FutureInventoryDaily)           // 删除Us004FutureInventoryDaily
		Us004FutureInventoryDailyRouter.DELETE("deleteUs004FutureInventoryDailyByIds", future_inventory.DeleteUs004FutureInventoryDailyByIds) // 批量删除Us004FutureInventoryDaily
		Us004FutureInventoryDailyRouter.PUT("updateUs004FutureInventoryDaily", future_inventory.UpdateUs004FutureInventoryDaily)              // 更新Us004FutureInventoryDaily
		Us004FutureInventoryDailyRouter.GET("findUs004FutureInventoryDaily", future_inventory.FindUs004FutureInventoryDaily)                  // 根据ID获取Us004FutureInventoryDaily
		Us004FutureInventoryDailyRouter.GET("getUs004FutureInventoryDailyList", future_inventory.GetUs004FutureInventoryDailyList)            // 根据ID获取Us004FutureInventoryDaily
		Us004FutureInventoryDailyRouter.GET("getUs004FutureInventoryDailyType", future_inventory.GetUs004FutureInventoryDailyType)
		Us004FutureInventoryDailyRouter.POST("/importInventoryExcel", future_inventory.ImportInventoryExcel) // 导入Excel           // 获取Us004FutureInventoryDaily列表

		Us004FutureInventoryDailyRouter.GET("loadInventoryExcelData", future_inventory.LoadInventoryExcelData)
	}
}
