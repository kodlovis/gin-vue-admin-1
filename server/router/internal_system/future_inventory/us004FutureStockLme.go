package future_inventory

import (
	v1 "gin-vue-admin/api/v1/internal_system/future_inventory"
	"gin-vue-admin/middleware"

	"github.com/gin-gonic/gin"
)

func InitUs004FutureStockLmeRouter(Router *gin.RouterGroup) {
	Us004FutureStockLmeRouter := Router.Group("us004FutureStockLme").Use(middleware.OperationRecord())
	{
		Us004FutureStockLmeRouter.POST("createUs004FutureStockLme", v1.CreateUs004FutureStockLme)             // 新建Us004FutureStockLme
		Us004FutureStockLmeRouter.DELETE("deleteUs004FutureStockLme", v1.DeleteUs004FutureStockLme)           // 删除Us004FutureStockLme
		Us004FutureStockLmeRouter.DELETE("deleteUs004FutureStockLmeByIds", v1.DeleteUs004FutureStockLmeByIds) // 批量删除Us004FutureStockLme
		Us004FutureStockLmeRouter.PUT("updateUs004FutureStockLme", v1.UpdateUs004FutureStockLme)              // 更新Us004FutureStockLme
		Us004FutureStockLmeRouter.GET("findUs004FutureStockLme", v1.FindUs004FutureStockLme)                  // 根据ID获取Us004FutureStockLme
		Us004FutureStockLmeRouter.GET("getUs004FutureStockLmeList", v1.GetUs004FutureStockLmeList)            // 获取Us004FutureStockLme列表
	}
}
