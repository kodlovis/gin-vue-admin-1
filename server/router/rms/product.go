package rms

import (
	"gin-vue-admin/api/v1/rms"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitProductRouter(Router *gin.RouterGroup) {
	ProductRouter := Router.Group("Product").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler()).Use(middleware.OperationRecord())
	{
		ProductRouter.POST("createProduct", rms.CreateProduct)   // 新建Product
		ProductRouter.DELETE("deleteProduct", rms.DeleteProduct) // 删除Product
		ProductRouter.DELETE("deleteProductByIds", rms.DeleteProductByIds) // 批量删除Product
		ProductRouter.PUT("updateProduct", rms.UpdateProduct)    // 更新Product
		ProductRouter.GET("findProduct", rms.FindProduct)        // 根据ID获取Product
		ProductRouter.GET("getProductList", rms.GetProductList)  // 获取Product列表
	}
}
