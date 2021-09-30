package futureData

import (
	"gin-vue-admin/api/v1/internalSystem/futureData"
	"gin-vue-admin/middleware"

	"github.com/gin-gonic/gin"
)

func InitRightsDetailRouter(Router *gin.RouterGroup) {
	RightsDetailRouter := Router.Group("rightsDetail").Use(middleware.OperationRecord())
	{
		RightsDetailRouter.POST("createRightsDetail", futureData.CreateRightsDetail)             // 新建RightsDetail
		RightsDetailRouter.DELETE("deleteRightsDetail", futureData.DeleteRightsDetail)           // 删除RightsDetail
		RightsDetailRouter.DELETE("deleteRightsDetailByIds", futureData.DeleteRightsDetailByIds) // 批量删除RightsDetail
		RightsDetailRouter.PUT("updateRightsDetail", futureData.UpdateRightsDetail)              // 更新RightsDetail
		RightsDetailRouter.GET("findRightsDetail", futureData.FindRightsDetail)                  // 根据ID获取RightsDetail
		RightsDetailRouter.GET("getRightsDetailList", futureData.GetRightsDetailList)            // 获取RightsDetail列表
	}
}
