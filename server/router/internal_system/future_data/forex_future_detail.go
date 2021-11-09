package future_data

import (
	"gin-vue-admin/api/v1/internal_system/future_data"
	"gin-vue-admin/middleware"

	"github.com/gin-gonic/gin"
)

func InitForexFutureDetailRouter(Router *gin.RouterGroup) {
	ForexFutureDetailRouter := Router.Group("forexFutureDetail").Use(middleware.OperationRecord())
	{
		ForexFutureDetailRouter.POST("createForexFutureDetail", future_data.CreateForexFutureDetail)             // 新建ForexFutureDetail
		ForexFutureDetailRouter.DELETE("deleteForexFutureDetail", future_data.DeleteForexFutureDetail)           // 删除ForexFutureDetail
		ForexFutureDetailRouter.DELETE("deleteForexFutureDetailByIds", future_data.DeleteForexFutureDetailByIds) // 批量删除ForexFutureDetail
		ForexFutureDetailRouter.PUT("updateForexFutureDetail", future_data.UpdateForexFutureDetail)              // 更新ForexFutureDetail
		ForexFutureDetailRouter.GET("findForexFutureDetail", future_data.FindForexFutureDetail)                  // 根据ID获取ForexFutureDetail
		ForexFutureDetailRouter.GET("getForexFutureDetailList", future_data.GetForexFutureDetailList)            // 获取ForexFutureDetail列表
	}
}
