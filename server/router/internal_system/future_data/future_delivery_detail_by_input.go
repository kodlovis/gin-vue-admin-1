package future_data

import (
	v1 "gin-vue-admin/api/v1/internal_system/future_data"
	"gin-vue-admin/middleware"

	"github.com/gin-gonic/gin"
)

func InitFutureDeliveryDetailByInputRouter(Router *gin.RouterGroup) {
	FutureDeliveryDetailByInputRouter := Router.Group("futureDeliveryDetailByInput").Use(middleware.OperationRecord())
	{
		FutureDeliveryDetailByInputRouter.POST("createFutureDeliveryDetailByInput", v1.CreateFutureDeliveryDetailByInput)             // 新建FutureDeliveryDetailByInput
		FutureDeliveryDetailByInputRouter.DELETE("deleteFutureDeliveryDetailByInput", v1.DeleteFutureDeliveryDetailByInput)           // 删除FutureDeliveryDetailByInput
		FutureDeliveryDetailByInputRouter.DELETE("deleteFutureDeliveryDetailByInputByIds", v1.DeleteFutureDeliveryDetailByInputByIds) // 批量删除FutureDeliveryDetailByInput
		FutureDeliveryDetailByInputRouter.PUT("updateFutureDeliveryDetailByInput", v1.UpdateFutureDeliveryDetailByInput)              // 更新FutureDeliveryDetailByInput
		FutureDeliveryDetailByInputRouter.GET("findFutureDeliveryDetailByInput", v1.FindFutureDeliveryDetailByInput)                  // 根据ID获取FutureDeliveryDetailByInput
		FutureDeliveryDetailByInputRouter.GET("getFutureDeliveryDetailByInputList", v1.GetFutureDeliveryDetailByInputList)            // 获取FutureDeliveryDetailByInput列表
	}
}
