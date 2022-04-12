package future_data

import (
	"gin-vue-admin/global"
	model "gin-vue-admin/model/internal_system/future_data"
	request "gin-vue-admin/model/request/internal_system/future_data"
	"gin-vue-admin/model/response"
	service "gin-vue-admin/service/internal_system/future_data"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// @Tags FutureDeliveryDetailByInput
// @Summary 创建FutureDeliveryDetailByInput
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.FutureDeliveryDetailByInput true "创建FutureDeliveryDetailByInput"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /futureDeliveryDetailByInput/createFutureDeliveryDetailByInput [post]
func CreateFutureDeliveryDetailByInput(c *gin.Context) {
	var futureDeliveryDetailByInput model.FutureDeliveryDetailByInput
	_ = c.ShouldBindJSON(&futureDeliveryDetailByInput)
	if err := service.CreateFutureDeliveryDetailByInput(futureDeliveryDetailByInput); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags FutureDeliveryDetailByInput
// @Summary 删除FutureDeliveryDetailByInput
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.FutureDeliveryDetailByInput true "删除FutureDeliveryDetailByInput"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /futureDeliveryDetailByInput/deleteFutureDeliveryDetailByInput [delete]
func DeleteFutureDeliveryDetailByInput(c *gin.Context) {
	var futureDeliveryDetailByInput model.FutureDeliveryDetailByInput
	_ = c.ShouldBindJSON(&futureDeliveryDetailByInput)
	if err := service.DeleteFutureDeliveryDetailByInput(futureDeliveryDetailByInput); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags FutureDeliveryDetailByInput
// @Summary 批量删除FutureDeliveryDetailByInput
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除FutureDeliveryDetailByInput"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /futureDeliveryDetailByInput/deleteFutureDeliveryDetailByInputByIds [delete]
func DeleteFutureDeliveryDetailByInputByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteFutureDeliveryDetailByInputByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags FutureDeliveryDetailByInput
// @Summary 更新FutureDeliveryDetailByInput
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.FutureDeliveryDetailByInput true "更新FutureDeliveryDetailByInput"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /futureDeliveryDetailByInput/updateFutureDeliveryDetailByInput [put]
func UpdateFutureDeliveryDetailByInput(c *gin.Context) {
	var futureDeliveryDetailByInput model.FutureDeliveryDetailByInput
	_ = c.ShouldBindJSON(&futureDeliveryDetailByInput)
	if err := service.UpdateFutureDeliveryDetailByInput(futureDeliveryDetailByInput); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags FutureDeliveryDetailByInput
// @Summary 用id查询FutureDeliveryDetailByInput
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.FutureDeliveryDetailByInput true "用id查询FutureDeliveryDetailByInput"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /futureDeliveryDetailByInput/findFutureDeliveryDetailByInput [get]
func FindFutureDeliveryDetailByInput(c *gin.Context) {
	var futureDeliveryDetailByInput model.FutureDeliveryDetailByInput
	_ = c.ShouldBindQuery(&futureDeliveryDetailByInput)
	if err, refutureDeliveryDetailByInput := service.GetFutureDeliveryDetailByInput(futureDeliveryDetailByInput.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"refutureDeliveryDetailByInput": refutureDeliveryDetailByInput}, c)
	}
}

// @Tags FutureDeliveryDetailByInput
// @Summary 分页获取FutureDeliveryDetailByInput列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.FutureDeliveryDetailByInputSearch true "分页获取FutureDeliveryDetailByInput列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /futureDeliveryDetailByInput/getFutureDeliveryDetailByInputList [get]
func GetFutureDeliveryDetailByInputList(c *gin.Context) {
	var pageInfo request.FutureDeliveryDetailByInputSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetFutureDeliveryDetailByInputInfoList(pageInfo); err != nil {
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
