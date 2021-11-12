package scale_rights

import (
	v1 "gin-vue-admin/api/v1/internal_system/scale_rights"
	"gin-vue-admin/middleware"

	"github.com/gin-gonic/gin"
)

func InitFutureDepartmentScaleDetailRouter(Router *gin.RouterGroup) {
	FutureDepartmentScaleDetailRouter := Router.Group("futureDepartmentScaleDetail").Use(middleware.OperationRecord())
	{
		FutureDepartmentScaleDetailRouter.POST("createFutureDepartmentScaleDetail", v1.CreateFutureDepartmentScaleDetail)             // 新建FutureDepartmentScaleDetail
		FutureDepartmentScaleDetailRouter.DELETE("deleteFutureDepartmentScaleDetail", v1.DeleteFutureDepartmentScaleDetail)           // 删除FutureDepartmentScaleDetail
		FutureDepartmentScaleDetailRouter.DELETE("deleteFutureDepartmentScaleDetailByIds", v1.DeleteFutureDepartmentScaleDetailByIds) // 批量删除FutureDepartmentScaleDetail
		FutureDepartmentScaleDetailRouter.PUT("updateFutureDepartmentScaleDetail", v1.UpdateFutureDepartmentScaleDetail)              // 更新FutureDepartmentScaleDetail
		FutureDepartmentScaleDetailRouter.GET("findFutureDepartmentScaleDetail", v1.FindFutureDepartmentScaleDetail)                  // 根据ID获取FutureDepartmentScaleDetail
		FutureDepartmentScaleDetailRouter.GET("getFutureDepartmentScaleDetailList", v1.GetFutureDepartmentScaleDetailList)            // 获取FutureDepartmentScaleDetail列表
	}
}
