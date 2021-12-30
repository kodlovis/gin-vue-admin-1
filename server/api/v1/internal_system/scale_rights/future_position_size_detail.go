package scale_rights

import (
	"gin-vue-admin/global"
	mp "gin-vue-admin/model/internal_system/scale_rights"
	rp "gin-vue-admin/model/request/internal_system/scale_rights"
	"gin-vue-admin/model/response"
	sp "gin-vue-admin/service/internal_system/scale_rights"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// @Tags FuturePositionSizeDetail
// @Summary 创建FuturePositionSizeDetail
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.FuturePositionSizeDetail true "创建FuturePositionSizeDetail"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /futurePositionSizeDetail/createFuturePositionSizeDetail [post]
func CreateFuturePositionSizeDetail(c *gin.Context) {
	var futurePositionSizeDetail mp.FuturePositionSizeDetail
	_ = c.ShouldBindJSON(&futurePositionSizeDetail)
	if err := sp.CreateFuturePositionSizeDetail(futurePositionSizeDetail); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags FuturePositionSizeDetail
// @Summary 删除FuturePositionSizeDetail
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.FuturePositionSizeDetail true "删除FuturePositionSizeDetail"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /futurePositionSizeDetail/deleteFuturePositionSizeDetail [delete]
func DeleteFuturePositionSizeDetail(c *gin.Context) {
	var futurePositionSizeDetail mp.FuturePositionSizeDetail
	_ = c.ShouldBindJSON(&futurePositionSizeDetail)
	if err := sp.DeleteFuturePositionSizeDetail(futurePositionSizeDetail); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags FuturePositionSizeDetail
// @Summary 批量删除FuturePositionSizeDetail
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除FuturePositionSizeDetail"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /futurePositionSizeDetail/deleteFuturePositionSizeDetailByIds [delete]
func DeleteFuturePositionSizeDetailByIds(c *gin.Context) {
	var IDS rp.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := sp.DeleteFuturePositionSizeDetailByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags FuturePositionSizeDetail
// @Summary 更新FuturePositionSizeDetail
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.FuturePositionSizeDetail true "更新FuturePositionSizeDetail"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /futurePositionSizeDetail/updateFuturePositionSizeDetail [put]
func UpdateFuturePositionSizeDetail(c *gin.Context) {
	var futurePositionSizeDetail mp.FuturePositionSizeDetail
	_ = c.ShouldBindJSON(&futurePositionSizeDetail)
	if err := sp.UpdateFuturePositionSizeDetail(futurePositionSizeDetail); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags FuturePositionSizeDetail
// @Summary 用id查询FuturePositionSizeDetail
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.FuturePositionSizeDetail true "用id查询FuturePositionSizeDetail"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /futurePositionSizeDetail/findFuturePositionSizeDetail [get]
func FindFuturePositionSizeDetail(c *gin.Context) {
	var futurePositionSizeDetail mp.FuturePositionSizeDetail
	_ = c.ShouldBindQuery(&futurePositionSizeDetail)
	if err, refuturePositionSizeDetail := sp.GetFuturePositionSizeDetail(futurePositionSizeDetail.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"refuturePositionSizeDetail": refuturePositionSizeDetail}, c)
	}
}

// @Tags FuturePositionSizeDetail
// @Summary 分页获取FuturePositionSizeDetail列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.FuturePositionSizeDetailSearch true "分页获取FuturePositionSizeDetail列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /futurePositionSizeDetail/getFuturePositionSizeDetailList [get]
func GetFuturePositionSizeDetailList(c *gin.Context) {
	var pageInfo rp.FuturePositionSizeDetailSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := sp.GetFuturePositionSizeDetailInfoList(pageInfo); err != nil {
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
