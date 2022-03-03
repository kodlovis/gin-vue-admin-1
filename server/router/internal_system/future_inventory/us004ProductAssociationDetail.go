package future_inventory

import (
	v1 "gin-vue-admin/api/v1/internal_system/future_inventory"
	"gin-vue-admin/middleware"

	"github.com/gin-gonic/gin"
)

func InitUs004ProductAssociationDetailRouter(Router *gin.RouterGroup) {
	Us004ProductAssociationDetailRouter := Router.Group("us004ProductAssociationDetail").Use(middleware.OperationRecord())
	{
		Us004ProductAssociationDetailRouter.POST("createUs004ProductAssociationDetail", v1.CreateUs004ProductAssociationDetail)             // 新建Us004ProductAssociationDetail
		Us004ProductAssociationDetailRouter.DELETE("deleteUs004ProductAssociationDetail", v1.DeleteUs004ProductAssociationDetail)           // 删除Us004ProductAssociationDetail
		Us004ProductAssociationDetailRouter.DELETE("deleteUs004ProductAssociationDetailByIds", v1.DeleteUs004ProductAssociationDetailByIds) // 批量删除Us004ProductAssociationDetail
		Us004ProductAssociationDetailRouter.PUT("updateUs004ProductAssociationDetail", v1.UpdateUs004ProductAssociationDetail)              // 更新Us004ProductAssociationDetail
		Us004ProductAssociationDetailRouter.GET("findUs004ProductAssociationDetail", v1.FindUs004ProductAssociationDetail)                  // 根据ID获取Us004ProductAssociationDetail
		Us004ProductAssociationDetailRouter.GET("getUs004ProductAssociationDetailList", v1.GetUs004ProductAssociationDetailList)            // 获取Us004ProductAssociationDetail列表
	}
}
