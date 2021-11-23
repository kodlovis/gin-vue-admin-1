package internal_system

import (
	v1 "gin-vue-admin/api/v1/internal_system"
	"gin-vue-admin/middleware"

	"github.com/gin-gonic/gin"
)

func InitExchangeProductInfoRouter(Router *gin.RouterGroup) {
	ExchangeProductInfoRouter := Router.Group("exchangeProductInfo").Use(middleware.OperationRecord())
	{
		ExchangeProductInfoRouter.POST("createExchangeProductInfo", v1.CreateExchangeProductInfo)             // 新建ExchangeProductInfo
		ExchangeProductInfoRouter.DELETE("deleteExchangeProductInfo", v1.DeleteExchangeProductInfo)           // 删除ExchangeProductInfo
		ExchangeProductInfoRouter.DELETE("deleteExchangeProductInfoByIds", v1.DeleteExchangeProductInfoByIds) // 批量删除ExchangeProductInfo
		ExchangeProductInfoRouter.PUT("updateExchangeProductInfo", v1.UpdateExchangeProductInfo)              // 更新ExchangeProductInfo
		ExchangeProductInfoRouter.GET("findExchangeProductInfo", v1.FindExchangeProductInfo)                  // 根据ID获取ExchangeProductInfo
		ExchangeProductInfoRouter.GET("getExchangeProductInfoList", v1.GetExchangeProductInfoList)            // 获取ExchangeProductInfo列表
	}
}
