package future_data

import (
	"gin-vue-admin/api/v1/internalSystem/future_data"
	"gin-vue-admin/middleware"

	"github.com/gin-gonic/gin"
)

func InitRightsDetailRouter(Router *gin.RouterGroup) {
	RightsDetailRouter := Router.Group("rightsDetail").Use(middleware.OperationRecord())
	{
		RightsDetailRouter.POST("createRightsDetail", future_data.CreateRightsDetail)             // 新建RightsDetail
		RightsDetailRouter.DELETE("deleteRightsDetail", future_data.DeleteRightsDetail)           // 删除RightsDetail
		RightsDetailRouter.DELETE("deleteRightsDetailByIds", future_data.DeleteRightsDetailByIds) // 批量删除RightsDetail
		RightsDetailRouter.PUT("updateRightsDetail", future_data.UpdateRightsDetail)              // 更新RightsDetail
		RightsDetailRouter.GET("findRightsDetail", future_data.FindRightsDetail)                  // 根据ID获取RightsDetail
		RightsDetailRouter.GET("getRightsDetailList", future_data.GetRightsDetailList)            // 获取RightsDetail列表
	}
}
