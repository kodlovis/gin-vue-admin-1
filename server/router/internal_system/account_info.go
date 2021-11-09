package internal_system

import (
	v1 "gin-vue-admin/api/v1/internal_system"
	"gin-vue-admin/middleware"

	"github.com/gin-gonic/gin"
)

func InitAccountInfoRouter(Router *gin.RouterGroup) {
	AccountInfoRouter := Router.Group("accountInfo").Use(middleware.OperationRecord())
	{
		AccountInfoRouter.POST("createAccountInfo", v1.CreateAccountInfo)             // 新建AccountInfo
		AccountInfoRouter.DELETE("deleteAccountInfo", v1.DeleteAccountInfo)           // 删除AccountInfo
		AccountInfoRouter.DELETE("deleteAccountInfoByIds", v1.DeleteAccountInfoByIds) // 批量删除AccountInfo
		AccountInfoRouter.PUT("updateAccountInfo", v1.UpdateAccountInfo)              // 更新AccountInfo
		AccountInfoRouter.GET("findAccountInfo", v1.FindAccountInfo)                  // 根据ID获取AccountInfo
		AccountInfoRouter.GET("getAccountInfoList", v1.GetAccountInfoList)            // 获取AccountInfo列表
	}
}
