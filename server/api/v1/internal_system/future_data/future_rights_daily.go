package future_data

import (
	"gin-vue-admin/global"
	mif "gin-vue-admin/model/internal_system/future_data"
	rif "gin-vue-admin/model/request/internal_system/future_data"
	"gin-vue-admin/model/response"
	sif "gin-vue-admin/service/internal_system/future_data"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// @Tags FutureRightsDaily
// @Summary 创建FutureRightsDaily
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.FutureRightsDaily true "创建FutureRightsDaily"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /futureRightsDaily/createFutureRightsDaily [post]
func CreateFutureRightsDaily(c *gin.Context) {
	var futureRightsDaily mif.FutureRightsDaily
	_ = c.ShouldBindJSON(&futureRightsDaily)
	if err := sif.CreateFutureRightsDaily(futureRightsDaily); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags FutureRightsDaily
// @Summary 删除FutureRightsDaily
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.FutureRightsDaily true "删除FutureRightsDaily"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /futureRightsDaily/deleteFutureRightsDaily [delete]
func DeleteFutureRightsDaily(c *gin.Context) {
	var futureRightsDaily mif.FutureRightsDaily
	_ = c.ShouldBindJSON(&futureRightsDaily)
	if err := sif.DeleteFutureRightsDaily(futureRightsDaily); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags FutureRightsDaily
// @Summary 批量删除FutureRightsDaily
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除FutureRightsDaily"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /futureRightsDaily/deleteFutureRightsDailyByIds [delete]
func DeleteFutureRightsDailyByIds(c *gin.Context) {
	var IDS rif.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := sif.DeleteFutureRightsDailyByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags FutureRightsDaily
// @Summary 更新FutureRightsDaily
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.FutureRightsDaily true "更新FutureRightsDaily"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /futureRightsDaily/updateFutureRightsDaily [put]
func UpdateFutureRightsDaily(c *gin.Context) {
	var futureRightsDaily mif.FutureRightsDaily
	_ = c.ShouldBindJSON(&futureRightsDaily)
	if err := sif.UpdateFutureRightsDaily(futureRightsDaily); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags FutureRightsDaily
// @Summary 用id查询FutureRightsDaily
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.FutureRightsDaily true "用id查询FutureRightsDaily"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /futureRightsDaily/findFutureRightsDaily [get]
func FindFutureRightsDaily(c *gin.Context) {
	var futureRightsDaily mif.FutureRightsDaily
	_ = c.ShouldBindQuery(&futureRightsDaily)
	if err, refutureRightsDaily := sif.GetFutureRightsDaily(futureRightsDaily.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"refutureRightsDaily": refutureRightsDaily}, c)
	}
}

// @Tags FutureRightsDaily
// @Summary 分页获取FutureRightsDaily列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.FutureRightsDailySearch true "分页获取FutureRightsDaily列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /futureRightsDaily/getFutureRightsDailyList [get]
func GetFutureRightsDailyList(c *gin.Context) {
	var pageInfo rif.FutureRightsDailySearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := sif.GetFutureRightsDailyInfoList(pageInfo); err != nil {
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
