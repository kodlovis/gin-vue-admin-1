package future_data

import (
	v1 "gin-vue-admin/api/v1/internal_system/future_data"
	"gin-vue-admin/middleware"

	"github.com/gin-gonic/gin"
)

func InitDepartmentRouter(Router *gin.RouterGroup) {
	DepartmentRouter := Router.Group("department").Use(middleware.OperationRecord())
	{
		DepartmentRouter.POST("createDepartment", v1.CreateDepartment)             // 新建Department
		DepartmentRouter.DELETE("deleteDepartment", v1.DeleteDepartment)           // 删除Department
		DepartmentRouter.DELETE("deleteDepartmentByIds", v1.DeleteDepartmentByIds) // 批量删除Department
		DepartmentRouter.PUT("updateDepartment", v1.UpdateDepartment)              // 更新Department
		DepartmentRouter.GET("findDepartment", v1.FindDepartment)                  // 根据ID获取Department
		DepartmentRouter.GET("getDepartmentList", v1.GetDepartmentList)            // 获取Department列表
	}
}
