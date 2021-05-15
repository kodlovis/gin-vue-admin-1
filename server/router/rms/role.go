package rms

import (
	"gin-vue-admin/api/v1/rms"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitRoleRouter(Router *gin.RouterGroup) {
	RoleRouter := Router.Group("Role").Use(middleware.JWTAuth()).Use(middleware.CasbinHandler()).Use(middleware.OperationRecord())
	{
		RoleRouter.POST("createRole", rms.CreateRole)   // 新建Role
		RoleRouter.DELETE("deleteRole", rms.DeleteRole) // 删除Role
		RoleRouter.DELETE("deleteRoleByIds", rms.DeleteRoleByIds) // 批量删除Role
		RoleRouter.PUT("updateRole", rms.UpdateRole)    // 更新Role
		RoleRouter.GET("findRole", rms.FindRole)        // 根据ID获取Role
		RoleRouter.GET("getRoleList", rms.GetRoleList)  // 获取Role列表
	}
}
