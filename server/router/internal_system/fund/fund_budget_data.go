package fund

import (
	v1 "gin-vue-admin/api/v1/internal_system/fund"
	"gin-vue-admin/middleware"

	"github.com/gin-gonic/gin"
)

func InitBudgetDataRouter(Router *gin.RouterGroup) {
	BudgetDataRouter := Router.Group("budget_data").Use(middleware.OperationRecord())
	{
		BudgetDataRouter.POST("createBudgetData", v1.CreateBudgetData)             // 新建BudgetData
		BudgetDataRouter.DELETE("deleteBudgetData", v1.DeleteBudgetData)           // 删除BudgetData
		BudgetDataRouter.DELETE("deleteBudgetDataByIds", v1.DeleteBudgetDataByIds) // 批量删除BudgetData
		BudgetDataRouter.PUT("updateBudgetData", v1.UpdateBudgetData)              // 更新BudgetData
		BudgetDataRouter.GET("findBudgetData", v1.FindBudgetData)                  // 根据ID获取BudgetData
		BudgetDataRouter.GET("getBudgetDataList", v1.GetBudgetDataList)            // 获取BudgetData列表

		BudgetDataRouter.POST("clickBudgetData", v1.ClickBudgetData) // 一天数据查询BudgetData

		BudgetDataRouter.POST("searchConfirmBudgetDate", v1.SearchConfirmBudgetDate) // 返回待确认数据
		BudgetDataRouter.POST("createConfirmBudgetDate", v1.CreateConfirmBudgetDate) // 录入确认数据
		BudgetDataRouter.POST("getConfirmBudgetDate", v1.GetConfirmBudgetDate)       // 获取录入情况

		BudgetDataRouter.POST("getCycleBudgetData", v1.CycleBudgetData) // 当前页面数据查询BudgetData，加合计算，绑定日历

		BudgetDataRouter.POST("createBudgetDatas", v1.CreateBudgetDatas) // 一条数据录入BudgetData

		BudgetDataRouter.POST("getReport", v1.GetAll) //获取当前报表总数据

	}
}
