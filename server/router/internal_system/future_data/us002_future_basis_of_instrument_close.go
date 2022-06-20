package future_data

import (
	v1 "gin-vue-admin/api/v1/internal_system/future_data"
	"gin-vue-admin/middleware"

	"github.com/gin-gonic/gin"
)

func InitUs002FutureBasisOfInstrumentCloseRouter(Router *gin.RouterGroup) {
	Us002FutureBasisOfInstrumentCloseRouter := Router.Group("us002FutureBasisOfInstrumentClose").Use(middleware.OperationRecord())
	{
		Us002FutureBasisOfInstrumentCloseRouter.POST("createUs002FutureBasisOfInstrumentClose", v1.CreateUs002FutureBasisOfInstrumentClose)             // 新建Us002FutureBasisOfInstrumentClose
		Us002FutureBasisOfInstrumentCloseRouter.DELETE("deleteUs002FutureBasisOfInstrumentClose", v1.DeleteUs002FutureBasisOfInstrumentClose)           // 删除Us002FutureBasisOfInstrumentClose
		Us002FutureBasisOfInstrumentCloseRouter.DELETE("deleteUs002FutureBasisOfInstrumentCloseByIds", v1.DeleteUs002FutureBasisOfInstrumentCloseByIds) // 批量删除Us002FutureBasisOfInstrumentClose
		Us002FutureBasisOfInstrumentCloseRouter.PUT("updateUs002FutureBasisOfInstrumentClose", v1.UpdateUs002FutureBasisOfInstrumentClose)              // 更新Us002FutureBasisOfInstrumentClose
		Us002FutureBasisOfInstrumentCloseRouter.GET("findUs002FutureBasisOfInstrumentClose", v1.FindUs002FutureBasisOfInstrumentClose)                  // 根据ID获取Us002FutureBasisOfInstrumentClose
		Us002FutureBasisOfInstrumentCloseRouter.GET("getUs002FutureBasisOfInstrumentCloseList", v1.GetUs002FutureBasisOfInstrumentCloseList)            // 获取Us002FutureBasisOfInstrumentClose列表
		Us002FutureBasisOfInstrumentCloseRouter.GET("getPositionDeliveryMonthInstrumentList", v1.GetPositionDeliveryMonthInstrumentList)                // 获取近月合约信息
		Us002FutureBasisOfInstrumentCloseRouter.GET("getBenchmarkInstrumentList", v1.GetBenchmarkInstrumentList)                                        // 更新Us002FutureBasisOfInstrumentClose
		Us002FutureBasisOfInstrumentCloseRouter.GET("getNoRiskValue", v1.GetNoRiskValue)                                                                // 获取基准合约
		Us002FutureBasisOfInstrumentCloseRouter.POST("createBasisOfInstrumentCloseList", v1.CreateBasisOfInstrumentCloseList)                           // 新建Us002FutureBasisOfInstrumentClose
	}
}
