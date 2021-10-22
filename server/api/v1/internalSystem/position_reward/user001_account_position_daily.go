package position_reward

import (
	"gin-vue-admin/global"
	mp "gin-vue-admin/model/internalSystem/position_reward"
	rp "gin-vue-admin/model/request/internalSystem/position_reward"
	"gin-vue-admin/model/response"
	sp "gin-vue-admin/service/internalSystem/position_reward"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// @Tags AccountPositionDaily
// @Summary 创建AccountPositionDaily
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AccountPositionDaily true "创建AccountPositionDaily"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /accountPositionDaily/createAccountPositionDaily [post]
func CreateAccountPositionDaily(c *gin.Context) {
	var accountPositionDaily mp.AccountPositionDaily
	_ = c.ShouldBindJSON(&accountPositionDaily)
	if err := sp.CreateAccountPositionDaily(accountPositionDaily); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags AccountPositionDaily
// @Summary 删除AccountPositionDaily
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AccountPositionDaily true "删除AccountPositionDaily"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /accountPositionDaily/deleteAccountPositionDaily [delete]
func DeleteAccountPositionDaily(c *gin.Context) {
	var accountPositionDaily mp.AccountPositionDaily
	_ = c.ShouldBindJSON(&accountPositionDaily)
	if err := sp.DeleteAccountPositionDaily(accountPositionDaily); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags AccountPositionDaily
// @Summary 批量删除AccountPositionDaily
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除AccountPositionDaily"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /accountPositionDaily/deleteAccountPositionDailyByIds [delete]
func DeleteAccountPositionDailyByIds(c *gin.Context) {
	var IDS rp.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := sp.DeleteAccountPositionDailyByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags AccountPositionDaily
// @Summary 更新AccountPositionDaily
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AccountPositionDaily true "更新AccountPositionDaily"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /accountPositionDaily/updateAccountPositionDaily [put]
func UpdateAccountPositionDaily(c *gin.Context) {
	var accountPositionDaily mp.AccountPositionDaily
	_ = c.ShouldBindJSON(&accountPositionDaily)
	if err := sp.UpdateAccountPositionDaily(accountPositionDaily); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags AccountPositionDaily
// @Summary 用id查询AccountPositionDaily
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AccountPositionDaily true "用id查询AccountPositionDaily"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /accountPositionDaily/findAccountPositionDaily [get]
func FindAccountPositionDaily(c *gin.Context) {
	var accountPositionDaily mp.AccountPositionDaily
	_ = c.ShouldBindQuery(&accountPositionDaily)
	if err, reus001AccountPositionDaily := sp.GetAccountPositionDaily(accountPositionDaily.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reus001AccountPositionDaily": reus001AccountPositionDaily}, c)
	}
}

// @Tags AccountPositionDaily
// @Summary 分页获取AccountPositionDaily列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.AccountPositionDailySearch true "分页获取AccountPositionDaily列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /accountPositionDaily/getAccountPositionDailyList [get]
func GetAccountPositionDailyList(c *gin.Context) {
	var pageInfo rp.AccountPositionDailySearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := sp.GetAccountPositionDailyInfoList(pageInfo); err != nil {
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
