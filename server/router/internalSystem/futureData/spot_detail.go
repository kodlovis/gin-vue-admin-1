package futureData

import (
	"gin-vue-admin/api/v1/internalSystem/futureData"
	"gin-vue-admin/middleware"

	"github.com/gin-gonic/gin"
)

func InitSpotDetailRouter(Router *gin.RouterGroup) {
	SpotDetailRouter := Router.Group("spotDetail").Use(middleware.OperationRecord())
	{
		SpotDetailRouter.POST("createSpotDetail", futureData.CreateSpotDetail)             // 新建SpotDetail
		SpotDetailRouter.DELETE("deleteSpotDetail", futureData.DeleteSpotDetail)           // 删除SpotDetail
		SpotDetailRouter.DELETE("deleteSpotDetailByIds", futureData.DeleteSpotDetailByIds) // 批量删除SpotDetail
		SpotDetailRouter.PUT("updateSpotDetail", futureData.UpdateSpotDetail)              // 更新SpotDetail
		SpotDetailRouter.GET("findSpotDetail", futureData.FindSpotDetail)                  // 根据ID获取SpotDetail
		SpotDetailRouter.GET("getSpotDetailList", futureData.GetSpotDetailList)            // 获取SpotDetail列表
	}
}
