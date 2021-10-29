package future_data

import (
	"gin-vue-admin/api/v1/internalSystem/future_data"
	"gin-vue-admin/middleware"

	"github.com/gin-gonic/gin"
)

func InitFutureAccountRouter(Router *gin.RouterGroup) {
	FutureAccountRouter := Router.Group("futureAccount").Use(middleware.OperationRecord())
	{
		FutureAccountRouter.POST("createFutureAccount", future_data.CreateFutureAccount)             // 新建FutureAccount
		FutureAccountRouter.DELETE("deleteFutureAccount", future_data.DeleteFutureAccount)           // 删除FutureAccount
		FutureAccountRouter.DELETE("deleteFutureAccountByIds", future_data.DeleteFutureAccountByIds) // 批量删除FutureAccount
		FutureAccountRouter.PUT("updateFutureAccount", future_data.UpdateFutureAccount)              // 更新FutureAccount
		FutureAccountRouter.GET("findFutureAccount", future_data.FindFutureAccount)                  // 根据ID获取FutureAccount
		FutureAccountRouter.GET("getFutureAccountList", future_data.GetFutureAccountList)            // 获取FutureAccount列表
	}
}
