package future_data

import (
	v1 "gin-vue-admin/api/v1/internal_system/future_data"
	"gin-vue-admin/middleware"

	"github.com/gin-gonic/gin"
)

func InitTradeDetailRouter(Router *gin.RouterGroup) {
	TradeDetailRouter := Router.Group("tradeDetail").Use(middleware.OperationRecord())
	{
		TradeDetailRouter.POST("createTradeDetail", v1.CreateTradeDetail)             // 新建TradeDetail
		TradeDetailRouter.DELETE("deleteTradeDetail", v1.DeleteTradeDetail)           // 删除TradeDetail
		TradeDetailRouter.DELETE("deleteTradeDetailByIds", v1.DeleteTradeDetailByIds) // 批量删除TradeDetail
		TradeDetailRouter.PUT("updateTradeDetail", v1.UpdateTradeDetail)              // 更新TradeDetail
		TradeDetailRouter.GET("findTradeDetail", v1.FindTradeDetail)                  // 根据ID获取TradeDetail
		TradeDetailRouter.GET("getTradeDetailList", v1.GetTradeDetailList)            // 获取TradeDetail列表
	}
}
