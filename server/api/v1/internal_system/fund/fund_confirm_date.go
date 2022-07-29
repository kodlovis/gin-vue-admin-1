package fund

import (
	"fmt"
	"gin-vue-admin/global"
	model "gin-vue-admin/model/internal_system/fund"
	request "gin-vue-admin/model/request/internal_system/fund"
	"gin-vue-admin/model/response"
	service "gin-vue-admin/service/internal_system/fund"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// @Tags FundConfirm
// @Summary 创建FundConfirm
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.FundConfirm true "创建FundConfirm"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /fundConfirmDate/createFundConfirm [post]
func CreateFundConfirm(c *gin.Context) {
	var fundConfirmDate model.FundConfirm
	_ = c.ShouldBindJSON(&fundConfirmDate)
	if err := service.CreateFundConfirm(fundConfirmDate); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags FundConfirm
// @Summary 删除FundConfirm
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.FundConfirm true "删除FundConfirm"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /fundConfirmDate/deleteFundConfirm [delete]
func DeleteFundConfirm(c *gin.Context) {
	var fundConfirmDate model.FundConfirm
	_ = c.ShouldBindJSON(&fundConfirmDate)
	if err := service.DeleteFundConfirm(fundConfirmDate); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags FundConfirm
// @Summary 批量删除FundConfirm
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除FundConfirm"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /fundConfirmDate/deleteFundConfirmByIds [delete]
func DeleteFundConfirmByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteFundConfirmByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags FundConfirm
// @Summary 更新FundConfirm
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.FundConfirm true "更新FundConfirm"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /fundConfirmDate/updateFundConfirm [put]
func UpdateFundConfirm(c *gin.Context) {
	var fundConfirmDate model.FundConfirm
	_ = c.ShouldBindJSON(&fundConfirmDate)
	if err := service.UpdateFundConfirm(fundConfirmDate); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags FundConfirm
// @Summary 用id查询FundConfirm
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.FundConfirm true "用id查询FundConfirm"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /fundConfirmDate/findFundConfirm [get]
func FindFundConfirm(c *gin.Context) {
	var fundConfirmDate model.FundConfirm
	_ = c.ShouldBindQuery(&fundConfirmDate)
	if err, refundConfirmDate := service.GetFundConfirm(fundConfirmDate.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"refundConfirmDate": refundConfirmDate}, c)
	}
}

// @Tags FundConfirm
// @Summary 分页获取FundConfirm列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.FundConfirmSearch true "分页获取FundConfirm列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /fundConfirmDate/getFundConfirmList [get]
func GetFundConfirmList(c *gin.Context) {
	var pageInfo request.FundConfirmSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetFundConfirmInfoList(pageInfo); err != nil {
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
func GetLastConfirmDate(c *gin.Context) {
	if err, list := service.GetLast(); err != nil {
		global.GVA_LOG.Error("获取失败", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List: list,
		}, "获取成功", c)
	}
}

//根据选择日期返回指定日期得预算截止日期
func GetCycleLast(c *gin.Context) {
	var start_date model.SearchFundConfirm
	_ = c.ShouldBindJSON(&start_date)
	if err, list := service.GetCycle(start_date.Start_date); err != nil {
		global.GVA_LOG.Error("获取失败", zap.Any("err", err))
		fmt.Println(err)
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List: list,
		}, "获取成功", c)
	}
}

//根据选择日期返回指定日期得预算截止日期
func GetConfirmSetting(c *gin.Context) {
	var start_date model.SearchConfirmSetting
	_ = c.ShouldBindJSON(&start_date)
	if err, list := service.GetSetting(start_date.Confirm_date); err != nil {
		global.GVA_LOG.Error("获取失败", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List: list,
		}, "获取成功", c)
	}
}
func UpConfirmSetting(c *gin.Context) {
	var confirm_data model.ConfirmSetting
	_ = c.ShouldBindJSON(&confirm_data)
	if err := service.UpSetting(confirm_data); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}
