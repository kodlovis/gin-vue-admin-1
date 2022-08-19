package public

import (
	v1 "gin-vue-admin/api/v1/internal_system/public"
	"gin-vue-admin/middleware"

	"github.com/gin-gonic/gin"
)

func InitUserDepartmentRouter(Router *gin.RouterGroup) {
	UserDepartmentRouter := Router.Group("userDepartment").Use(middleware.OperationRecord())
	{
		UserDepartmentRouter.POST("createUserDepartment", v1.CreateUserDepartment)             // 新建UserDepartment
		UserDepartmentRouter.DELETE("deleteUserDepartment", v1.DeleteUserDepartment)           // 删除UserDepartment
		UserDepartmentRouter.DELETE("deleteUserDepartmentByIds", v1.DeleteUserDepartmentByIds) // 批量删除UserDepartment
		UserDepartmentRouter.PUT("updateUserDepartment", v1.UpdateUserDepartment)              // 更新UserDepartment
		UserDepartmentRouter.GET("findUserDepartment", v1.FindUserDepartment)                  // 根据ID获取UserDepartment
		UserDepartmentRouter.GET("getUserDepartmentList", v1.GetUserDepartmentList)            // 获取UserDepartment列表
		UserDepartmentRouter.GET("getUserDepartmentListByID", v1.GetUserDepartmentListByID)    // 获取UserDepartment列表
	}
}
