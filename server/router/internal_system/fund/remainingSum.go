package fund

import (
	v1 "gin-vue-admin/api/v1/internal_system/fund"
	"gin-vue-admin/middleware"

	"github.com/gin-gonic/gin"
)

func InitRemainingSumRouter(Router *gin.RouterGroup) {
	RemainingSumRouter := Router.Group("remainingSum").Use(middleware.OperationRecord())
	{
		RemainingSumRouter.POST("createRemainingSum", v1.CreateRemainingSum)             // 新建RemainingSum
		RemainingSumRouter.DELETE("deleteRemainingSum", v1.DeleteRemainingSum)           // 删除RemainingSum
		RemainingSumRouter.DELETE("deleteRemainingSumByIds", v1.DeleteRemainingSumByIds) // 批量删除RemainingSum
		RemainingSumRouter.PUT("updateRemainingSum", v1.UpdateRemainingSum)              // 更新RemainingSum
		RemainingSumRouter.GET("findRemainingSum", v1.FindRemainingSum)                  // 根据ID获取RemainingSum
		RemainingSumRouter.GET("getRemainingSumList", v1.GetRemainingSumList)            // 获取RemainingSum列表

		RemainingSumRouter.POST("getOneRemainingSum", v1.GetOneRemainingSum) // 获取该日
		RemainingSumRouter.POST("getwarehouse", v1.Getwarehouse)             // 获取该日
	}
}
