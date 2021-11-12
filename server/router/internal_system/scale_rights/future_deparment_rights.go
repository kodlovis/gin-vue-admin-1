package scale_rights

import (
	v1 "gin-vue-admin/api/v1/internal_system/scale_rights"
	"gin-vue-admin/middleware"

	"github.com/gin-gonic/gin"
)

func InitFutureDeparmentRightsRouter(Router *gin.RouterGroup) {
	FutureDeparmentRightsRouter := Router.Group("futureDeparmentRights").Use(middleware.OperationRecord())
	{
		FutureDeparmentRightsRouter.POST("createFutureDeparmentRights", v1.CreateFutureDeparmentRights)             // 新建FutureDeparmentRights
		FutureDeparmentRightsRouter.DELETE("deleteFutureDeparmentRights", v1.DeleteFutureDeparmentRights)           // 删除FutureDeparmentRights
		FutureDeparmentRightsRouter.DELETE("deleteFutureDeparmentRightsByIds", v1.DeleteFutureDeparmentRightsByIds) // 批量删除FutureDeparmentRights
		FutureDeparmentRightsRouter.PUT("updateFutureDeparmentRights", v1.UpdateFutureDeparmentRights)              // 更新FutureDeparmentRights
		FutureDeparmentRightsRouter.GET("findFutureDeparmentRights", v1.FindFutureDeparmentRights)                  // 根据ID获取FutureDeparmentRights
		FutureDeparmentRightsRouter.GET("getFutureDeparmentRightsList", v1.GetFutureDeparmentRightsList)            // 获取FutureDeparmentRights列表
	}
}
