package future_data

import (
	v1 "gin-vue-admin/api/v1/internal_system/future_data"
	"gin-vue-admin/middleware"

	"github.com/gin-gonic/gin"
)

func InitUs002ForeignExchangeProfitRouter(Router *gin.RouterGroup) {
	Us002ForeignExchangeProfitRouter := Router.Group("us002ForeignExchangeProfit").Use(middleware.OperationRecord())
	{
		Us002ForeignExchangeProfitRouter.POST("createUs002ForeignExchangeProfit", v1.CreateUs002ForeignExchangeProfit)             // 新建Us002ForeignExchangeProfit
		Us002ForeignExchangeProfitRouter.DELETE("deleteUs002ForeignExchangeProfit", v1.DeleteUs002ForeignExchangeProfit)           // 删除Us002ForeignExchangeProfit
		Us002ForeignExchangeProfitRouter.DELETE("deleteUs002ForeignExchangeProfitByIds", v1.DeleteUs002ForeignExchangeProfitByIds) // 批量删除Us002ForeignExchangeProfit
		Us002ForeignExchangeProfitRouter.PUT("updateUs002ForeignExchangeProfit", v1.UpdateUs002ForeignExchangeProfit)              // 更新Us002ForeignExchangeProfit
		Us002ForeignExchangeProfitRouter.GET("findUs002ForeignExchangeProfit", v1.FindUs002ForeignExchangeProfit)                  // 根据ID获取Us002ForeignExchangeProfit
		Us002ForeignExchangeProfitRouter.GET("getUs002ForeignExchangeProfitList", v1.GetUs002ForeignExchangeProfitList)            // 获取Us002ForeignExchangeProfit列表
	}
}
