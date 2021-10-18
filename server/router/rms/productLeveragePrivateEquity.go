package rms

import (
	"gin-vue-admin/api/v1/rms"
	"gin-vue-admin/middleware"

	"github.com/gin-gonic/gin"
)

func InitProductLeveragePrivateEquityRouter(Router *gin.RouterGroup) {
	ProductLeveragePrivateEquityRouter := Router.Group("productLeveragePrivateEquity").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler()).Use(middleware.OperationRecord())
	{
		ProductLeveragePrivateEquityRouter.POST("createProductLeveragePrivateEquity", rms.CreateProductLeveragePrivateEquity)             // 新建ProductLeveragePrivateEquity
		ProductLeveragePrivateEquityRouter.DELETE("deleteProductLeveragePrivateEquity", rms.DeleteProductLeveragePrivateEquity)           // 删除ProductLeveragePrivateEquity
		ProductLeveragePrivateEquityRouter.DELETE("deleteProductLeveragePrivateEquityByIds", rms.DeleteProductLeveragePrivateEquityByIds) // 批量删除ProductLeveragePrivateEquity
		ProductLeveragePrivateEquityRouter.PUT("updateProductLeveragePrivateEquity", rms.UpdateProductLeveragePrivateEquity)              // 更新ProductLeveragePrivateEquity
		ProductLeveragePrivateEquityRouter.GET("findProductLeveragePrivateEquity", rms.FindProductLeveragePrivateEquity)                  // 根据ID获取ProductLeveragePrivateEquity
		ProductLeveragePrivateEquityRouter.GET("getProductLeveragePrivateEquityList", rms.GetProductLeveragePrivateEquityList)            // 获取ProductLeveragePrivateEquity列表
	}
}
