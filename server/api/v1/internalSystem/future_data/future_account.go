package future_data

import (
	"gin-vue-admin/global"
	mif "gin-vue-admin/model/internalSystem/future_data"
	rif "gin-vue-admin/model/request/internalSystem/future_data"
	"gin-vue-admin/model/response"
	sif "gin-vue-admin/service/internalSystem/future_data"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// @Tags FutureAccount
// @Summary 创建FutureAccount
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.FutureAccount true "创建FutureAccount"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /futureAccount/createFutureAccount [post]
func CreateFutureAccount(c *gin.Context) {
	var futureAccount mif.FutureAccount
	_ = c.ShouldBindJSON(&futureAccount)
	if err := sif.CreateFutureAccount(futureAccount); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags FutureAccount
// @Summary 删除FutureAccount
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.FutureAccount true "删除FutureAccount"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /futureAccount/deleteFutureAccount [delete]
func DeleteFutureAccount(c *gin.Context) {
	var futureAccount mif.FutureAccount
	_ = c.ShouldBindJSON(&futureAccount)
	if err := sif.DeleteFutureAccount(futureAccount); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags FutureAccount
// @Summary 批量删除FutureAccount
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除FutureAccount"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /futureAccount/deleteFutureAccountByIds [delete]
func DeleteFutureAccountByIds(c *gin.Context) {
	var IDS rif.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := sif.DeleteFutureAccountByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags FutureAccount
// @Summary 更新FutureAccount
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.FutureAccount true "更新FutureAccount"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /futureAccount/updateFutureAccount [put]
func UpdateFutureAccount(c *gin.Context) {
	var futureAccount mif.FutureAccount
	_ = c.ShouldBindJSON(&futureAccount)
	if err := sif.UpdateFutureAccount(futureAccount); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags FutureAccount
// @Summary 用id查询FutureAccount
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.FutureAccount true "用id查询FutureAccount"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /futureAccount/findFutureAccount [get]
func FindFutureAccount(c *gin.Context) {
	var futureAccount mif.FutureAccount
	_ = c.ShouldBindQuery(&futureAccount)
	if err, refutureAccount := sif.GetFutureAccount(futureAccount.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"refutureAccount": refutureAccount}, c)
	}
}

// @Tags FutureAccount
// @Summary 分页获取FutureAccount列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.FutureAccountSearch true "分页获取FutureAccount列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /futureAccount/getFutureAccountList [get]
func GetFutureAccountList(c *gin.Context) {
	var pageInfo rif.FutureAccountSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := sif.GetFutureAccountInfoList(pageInfo); err != nil {
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
