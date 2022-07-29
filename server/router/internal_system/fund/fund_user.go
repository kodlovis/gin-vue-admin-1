package fund

import (
	v1 "gin-vue-admin/api/v1/internal_system/fund"
	"gin-vue-admin/middleware"

	"github.com/gin-gonic/gin"
)

func InitFundUserRouter(Router *gin.RouterGroup) {
	FundUserRouter := Router.Group("getfund").Use(middleware.OperationRecord())
	{
		FundUserRouter.POST("role", v1.GetFundRole) // 查询角色
	}
}
