package fund

import (
	"gin-vue-admin/global"
	model "gin-vue-admin/model/internal_system/fund"
	request "gin-vue-admin/model/request/internal_system/fund"
	"gin-vue-admin/model/response"
	service "gin-vue-admin/service/internal_system/fund"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// @Tags BudgetData
// @Summary 创建BudgetData
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BudgetData true "创建BudgetData"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /budget_data/createBudgetData [post]
func CreateBudgetData(c *gin.Context) {
	var budget_data model.BudgetData
	_ = c.ShouldBindJSON(&budget_data)
	if err := service.CreateBudgetData(budget_data); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}
func CreateBudgetDatas(c *gin.Context) {

	var budgets []model.BudgetData
	_ = c.ShouldBindJSON(&budgets)
	var confirm_setting model.ConfirmSetting
	confirm_date := budgets[0].Up_date
	create_role := budgets[0].Create_role
	department_code := budgets[0].Department_code
	db := global.GVA_DB.Model(&model.ConfirmSetting{})
	_ = db.Where("confirm_date = ?", confirm_date).Find(&confirm_setting).Error

	var is_confirm model.ConfirmBudget
	confirm_db := global.GVA_DB.Model(&model.ConfirmBudget{})
	_ = confirm_db.Where("confirm_date = ? and create_role = ? and department_code = ?", confirm_date, create_role, department_code).Delete(&is_confirm)

	if confirm_setting.Is_ok == 0 {
		if err := service.SaveBudgetData(budgets); err != nil {
			global.GVA_LOG.Error("创建失败!", zap.Error(err))
			response.FailWithMessage("创建失败", c)
		} else {
			response.OkWithMessage("创建成功", c)
		}
	} else {
		response.FailWithMessage("录入关闭", c)
	}

}

func ClickBudgetData(c *gin.Context) {
	var searchInfo model.SearchDay
	_ = c.ShouldBindJSON(&searchInfo)
	if err, lists := service.FindDayData(searchInfo); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List: lists,
		}, "获取成功", c)

	}
}
func CycleBudgetData(c *gin.Context) {

	var searchInfo model.SearchDay
	_ = c.ShouldBindJSON(&searchInfo)
	if err, lists := service.GetDayData(searchInfo); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List: lists,
		}, "获取成功", c)

	}
}

func GetAll(c *gin.Context) {
	var searchInfo model.SearchDay
	_ = c.ShouldBindJSON(&searchInfo)
	if err, lists := service.GetReport(searchInfo); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List: lists,
		}, "获取成功", c)

	}
}

// @Tags BudgetData
// @Summary 删除BudgetData
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BudgetData true "删除BudgetData"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /budget_data/deleteBudgetData [delete]
func DeleteBudgetData(c *gin.Context) {
	var budget_data model.BudgetData
	_ = c.ShouldBindJSON(&budget_data)
	if err := service.DeleteBudgetData(budget_data); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags BudgetData
// @Summary 批量删除BudgetData
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除BudgetData"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /budget_data/deleteBudgetDataByIds [delete]
func DeleteBudgetDataByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteBudgetDataByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags BudgetData
// @Summary 更新BudgetData
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BudgetData true "更新BudgetData"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /budget_data/updateBudgetData [put]
func UpdateBudgetData(c *gin.Context) {
	var budget_data model.BudgetData
	_ = c.ShouldBindJSON(&budget_data)
	if err := service.UpdateBudgetData(budget_data); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags BudgetData
// @Summary 用id查询BudgetData
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BudgetData true "用id查询BudgetData"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /budget_data/findBudgetData [get]
func FindBudgetData(c *gin.Context) {
	var budget_data model.BudgetData
	_ = c.ShouldBindQuery(&budget_data)
	if err, rebudget_data := service.GetBudgetData(budget_data.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rebudget_data": rebudget_data}, c)
	}
}

// @Tags BudgetData
// @Summary 分页获取BudgetData列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.BudgetDataSearch true "分页获取BudgetData列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /budget_data/getBudgetDataList [get]
func GetBudgetDataList(c *gin.Context) {
	var pageInfo request.BudgetDataSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetBudgetDataInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}
