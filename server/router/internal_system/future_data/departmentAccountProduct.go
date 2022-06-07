package future_data

import (
	v1 "gin-vue-admin/api/v1/internal_system/future_data"
	"gin-vue-admin/middleware"

	"github.com/gin-gonic/gin"
)

func InitDepartmentAccountProductRouter(Router *gin.RouterGroup) {
	DepartmentAccountProductRouter := Router.Group("departmentAccountProduct").Use(middleware.OperationRecord())
	{
		DepartmentAccountProductRouter.POST("createDepartmentAccountProduct", v1.CreateDepartmentAccountProduct)             // 新建DepartmentAccountProduct
		DepartmentAccountProductRouter.DELETE("deleteDepartmentAccountProduct", v1.DeleteDepartmentAccountProduct)           // 删除DepartmentAccountProduct
		DepartmentAccountProductRouter.DELETE("deleteDepartmentAccountProductByIds", v1.DeleteDepartmentAccountProductByIds) // 批量删除DepartmentAccountProduct
		DepartmentAccountProductRouter.PUT("updateDepartmentAccountProduct", v1.UpdateDepartmentAccountProduct)              // 更新DepartmentAccountProduct
		DepartmentAccountProductRouter.GET("findDepartmentAccountProduct", v1.FindDepartmentAccountProduct)                  // 根据ID获取DepartmentAccountProduct
		DepartmentAccountProductRouter.GET("getDepartmentAccountProductList", v1.GetDepartmentAccountProductList)            // 获取DepartmentAccountProduct列表
	}
}
