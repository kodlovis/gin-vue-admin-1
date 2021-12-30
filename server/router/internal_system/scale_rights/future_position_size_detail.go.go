package scale_rights

import (
	v1 "gin-vue-admin/api/v1/internal_system/scale_rights"
	"gin-vue-admin/middleware"

	"github.com/gin-gonic/gin"
)

func InitFuturePositionSizeDetailRouter(Router *gin.RouterGroup) {
	FuturePositionSizeDetailRouter := Router.Group("futurePositionSizeDetail").Use(middleware.OperationRecord())
	{
		FuturePositionSizeDetailRouter.POST("createFuturePositionSizeDetail", v1.CreateFuturePositionSizeDetail)             // 新建FuturePositionSizeDetail
		FuturePositionSizeDetailRouter.DELETE("deleteFuturePositionSizeDetail", v1.DeleteFuturePositionSizeDetail)           // 删除FuturePositionSizeDetail
		FuturePositionSizeDetailRouter.DELETE("deleteFuturePositionSizeDetailByIds", v1.DeleteFuturePositionSizeDetailByIds) // 批量删除FuturePositionSizeDetail
		FuturePositionSizeDetailRouter.PUT("updateFuturePositionSizeDetail", v1.UpdateFuturePositionSizeDetail)              // 更新FuturePositionSizeDetail
		FuturePositionSizeDetailRouter.GET("findFuturePositionSizeDetail", v1.FindFuturePositionSizeDetail)                  // 根据ID获取FuturePositionSizeDetail
		FuturePositionSizeDetailRouter.GET("getFuturePositionSizeDetailList", v1.GetFuturePositionSizeDetailList)            // 获取FuturePositionSizeDetail列表
	}
}
