package futureData

import (
	"gin-vue-admin/api/v1/internalSystem/futureData"
	"gin-vue-admin/middleware"

	"github.com/gin-gonic/gin"
)

func InitForexFutureDetailRouter(Router *gin.RouterGroup) {
	ForexFutureDetailRouter := Router.Group("forexFutureDetail").Use(middleware.OperationRecord())
	{
		ForexFutureDetailRouter.POST("createForexFutureDetail", futureData.CreateForexFutureDetail)             // 新建ForexFutureDetail
		ForexFutureDetailRouter.DELETE("deleteForexFutureDetail", futureData.DeleteForexFutureDetail)           // 删除ForexFutureDetail
		ForexFutureDetailRouter.DELETE("deleteForexFutureDetailByIds", futureData.DeleteForexFutureDetailByIds) // 批量删除ForexFutureDetail
		ForexFutureDetailRouter.PUT("updateForexFutureDetail", futureData.UpdateForexFutureDetail)              // 更新ForexFutureDetail
		ForexFutureDetailRouter.GET("findForexFutureDetail", futureData.FindForexFutureDetail)                  // 根据ID获取ForexFutureDetail
		ForexFutureDetailRouter.GET("getForexFutureDetailList", futureData.GetForexFutureDetailList)            // 获取ForexFutureDetail列表
	}
}
