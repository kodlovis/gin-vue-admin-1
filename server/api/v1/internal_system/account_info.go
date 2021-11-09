package internal_system

import (
	"gin-vue-admin/global"
	mp "gin-vue-admin/model/internal_system"
	rp "gin-vue-admin/model/request/internal_system"
	"gin-vue-admin/model/response"
	sp "gin-vue-admin/service/internal_system"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// @Tags AccountInfo
// @Summary 创建AccountInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AccountInfo true "创建AccountInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /accountInfo/createAccountInfo [post]
func CreateAccountInfo(c *gin.Context) {
	var accountInfo mp.AccountInfo
	_ = c.ShouldBindJSON(&accountInfo)
	if err := sp.CreateAccountInfo(accountInfo); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags AccountInfo
// @Summary 删除AccountInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AccountInfo true "删除AccountInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /accountInfo/deleteAccountInfo [delete]
func DeleteAccountInfo(c *gin.Context) {
	var accountInfo mp.AccountInfo
	_ = c.ShouldBindJSON(&accountInfo)
	if err := sp.DeleteAccountInfo(accountInfo); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags AccountInfo
// @Summary 批量删除AccountInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除AccountInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /accountInfo/deleteAccountInfoByIds [delete]
func DeleteAccountInfoByIds(c *gin.Context) {
	var IDS rp.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := sp.DeleteAccountInfoByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags AccountInfo
// @Summary 更新AccountInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AccountInfo true "更新AccountInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /accountInfo/updateAccountInfo [put]
func UpdateAccountInfo(c *gin.Context) {
	var accountInfo mp.AccountInfo
	_ = c.ShouldBindJSON(&accountInfo)
	if err := sp.UpdateAccountInfo(accountInfo); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags AccountInfo
// @Summary 用id查询AccountInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AccountInfo true "用id查询AccountInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /accountInfo/findAccountInfo [get]
func FindAccountInfo(c *gin.Context) {
	var accountInfo mp.AccountInfo
	_ = c.ShouldBindQuery(&accountInfo)
	if err, reaccountInfo := sp.GetAccountInfo(accountInfo.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reaccountInfo": reaccountInfo}, c)
	}
}

// @Tags AccountInfo
// @Summary 分页获取AccountInfo列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.AccountInfoSearch true "分页获取AccountInfo列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /accountInfo/getAccountInfoList [get]
func GetAccountInfoList(c *gin.Context) {
	var pageInfo rp.AccountInfoSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := sp.GetAccountInfoInfoList(pageInfo); err != nil {
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
