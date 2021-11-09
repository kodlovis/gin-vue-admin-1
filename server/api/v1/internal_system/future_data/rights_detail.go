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

// @Tags RightsDetail
// @Summary 创建RightsDetail
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.RightsDetail true "创建RightsDetail"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /rightsDetail/createRightsDetail [post]
func CreateRightsDetail(c *gin.Context) {
	var rightsDetail mif.RightsDetail
	_ = c.ShouldBindJSON(&rightsDetail)
	if err := sif.CreateRightsDetail(rightsDetail); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags RightsDetail
// @Summary 删除RightsDetail
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.RightsDetail true "删除RightsDetail"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /rightsDetail/deleteRightsDetail [delete]
func DeleteRightsDetail(c *gin.Context) {
	var rightsDetail mif.RightsDetail
	_ = c.ShouldBindJSON(&rightsDetail)
	if err := sif.DeleteRightsDetail(rightsDetail); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags RightsDetail
// @Summary 批量删除RightsDetail
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除RightsDetail"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /rightsDetail/deleteRightsDetailByIds [delete]
func DeleteRightsDetailByIds(c *gin.Context) {
	var IDS rif.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := sif.DeleteRightsDetailByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags RightsDetail
// @Summary 更新RightsDetail
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.RightsDetail true "更新RightsDetail"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /rightsDetail/updateRightsDetail [put]
func UpdateRightsDetail(c *gin.Context) {
	var rightsDetail mif.RightsDetail
	_ = c.ShouldBindJSON(&rightsDetail)
	if err := sif.UpdateRightsDetail(rightsDetail); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags RightsDetail
// @Summary 用id查询RightsDetail
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.RightsDetail true "用id查询RightsDetail"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /rightsDetail/findRightsDetail [get]
func FindRightsDetail(c *gin.Context) {
	var rightsDetail mif.RightsDetail
	_ = c.ShouldBindQuery(&rightsDetail)
	if err, rerightsDetail := sif.GetRightsDetail(rightsDetail.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rerightsDetail": rerightsDetail}, c)
	}
}

// @Tags RightsDetail
// @Summary 分页获取RightsDetail列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.RightsDetailSearch true "分页获取RightsDetail列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /rightsDetail/getRightsDetailList [get]
func GetRightsDetailList(c *gin.Context) {
	var pageInfo rif.RightsDetailSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := sif.GetRightsDetailInfoList(pageInfo); err != nil {
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
