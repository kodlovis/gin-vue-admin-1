package future_data

import (
	"gin-vue-admin/api/v1/internal_system/future_data"
	"gin-vue-admin/middleware"

	"github.com/gin-gonic/gin"
)

func InitSpotDetailRouter(Router *gin.RouterGroup) {
	SpotDetailRouter := Router.Group("spotDetail").Use(middleware.OperationRecord())
	{
		SpotDetailRouter.POST("createSpotDetail", future_data.CreateSpotDetail)             // 新建SpotDetail
		SpotDetailRouter.DELETE("deleteSpotDetail", future_data.DeleteSpotDetail)           // 删除SpotDetail
		SpotDetailRouter.DELETE("deleteSpotDetailByIds", future_data.DeleteSpotDetailByIds) // 批量删除SpotDetail
		SpotDetailRouter.PUT("updateSpotDetail", future_data.UpdateSpotDetail)              // 更新SpotDetail
		SpotDetailRouter.GET("findSpotDetail", future_data.FindSpotDetail)                  // 根据ID获取SpotDetail
		SpotDetailRouter.GET("getSpotDetailList", future_data.GetSpotDetailList)            // 获取SpotDetail列表
	}
}
