package public

import (
	v1 "gin-vue-admin/api/v1/internal_system/public"
	"gin-vue-admin/middleware"

	"github.com/gin-gonic/gin"
)

func InitUs100CurrencyRouter(Router *gin.RouterGroup) {
	Us100CurrencyRouter := Router.Group("us100Currency").Use(middleware.OperationRecord())
	{
		Us100CurrencyRouter.POST("createUs100Currency", v1.CreateUs100Currency)             // 新建Us100Currency
		Us100CurrencyRouter.DELETE("deleteUs100Currency", v1.DeleteUs100Currency)           // 删除Us100Currency
		Us100CurrencyRouter.DELETE("deleteUs100CurrencyByIds", v1.DeleteUs100CurrencyByIds) // 批量删除Us100Currency
		Us100CurrencyRouter.PUT("updateUs100Currency", v1.UpdateUs100Currency)              // 更新Us100Currency
		Us100CurrencyRouter.GET("findUs100Currency", v1.FindUs100Currency)                  // 根据ID获取Us100Currency
		Us100CurrencyRouter.GET("getUs100CurrencyList", v1.GetUs100CurrencyList)            // 获取Us100Currency列表
	}
}
