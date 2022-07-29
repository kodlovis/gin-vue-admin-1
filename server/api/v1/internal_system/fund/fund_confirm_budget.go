package fund

import (
	"gin-vue-admin/global"
	model "gin-vue-admin/model/internal_system/fund"

	// request "gin-vue-admin/model/request/internal_system/fund"
	"gin-vue-admin/model/response"
	service "gin-vue-admin/service/internal_system/fund"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

//创建确认数据
func CreateConfirmBudgetDate(c *gin.Context) {
	var confirmBudget model.ConfirmBudget
	var confirm_setting model.ConfirmSetting
	_ = c.ShouldBindJSON(&confirmBudget)
	confirm_date := confirmBudget.Confirm_date

	db := global.GVA_DB.Model(&model.ConfirmSetting{})
	_ = db.Where("confirm_date = ?", confirm_date).Find(&confirm_setting).Error
	if confirm_setting.Is_ok == 0 {
		if err := service.CreateconfirmBudget(confirmBudget); err != nil {
			global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
			response.FailWithMessage("创建失败", c)
		} else {
			response.OkWithMessage("创建成功", c)
		}
	} else {
		response.FailWithMessage("录入关闭", c)
	}

}

//获取当日录入数据
func SearchConfirmBudgetDate(c *gin.Context) {
	var searchConfirmBudgetDate model.SearchConfirm
	_ = c.ShouldBindJSON(&searchConfirmBudgetDate)
	if err, lists := service.FindConfirmBudget(searchConfirmBudgetDate); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List: lists,
		}, "获取成功", c)

	}
}

func GetConfirmBudgetDate(c *gin.Context) {
	var choseDate model.LimitData
	_ = c.ShouldBindJSON(&choseDate)
	if err, lists := service.GetConfirmBudget(choseDate); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List: lists,
		}, "获取成功", c)
	}
}
