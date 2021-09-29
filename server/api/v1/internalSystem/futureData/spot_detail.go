package futureData

import (
	"gin-vue-admin/global"
	mif "gin-vue-admin/model/internalSystem/futureData"
	rif "gin-vue-admin/model/request/internalSystem/futureData"
	"gin-vue-admin/model/response"
	sif "gin-vue-admin/service/internalSystem/futureData"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// @Tags SpotDetail
// @Summary 创建SpotDetail
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SpotDetail true "创建SpotDetail"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /spotDetail/createSpotDetail [post]
func CreateSpotDetail(c *gin.Context) {
	var spotDetail mif.SpotDetail
	_ = c.ShouldBindJSON(&spotDetail)
	if err := sif.CreateSpotDetail(spotDetail); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags SpotDetail
// @Summary 删除SpotDetail
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SpotDetail true "删除SpotDetail"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /spotDetail/deleteSpotDetail [delete]
func DeleteSpotDetail(c *gin.Context) {
	var spotDetail mif.SpotDetail
	_ = c.ShouldBindJSON(&spotDetail)
	if err := sif.DeleteSpotDetail(spotDetail); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags SpotDetail
// @Summary 批量删除SpotDetail
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除SpotDetail"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /spotDetail/deleteSpotDetailByIds [delete]
func DeleteSpotDetailByIds(c *gin.Context) {
	var IDS rif.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := sif.DeleteSpotDetailByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags SpotDetail
// @Summary 更新SpotDetail
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SpotDetail true "更新SpotDetail"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /spotDetail/updateSpotDetail [put]
func UpdateSpotDetail(c *gin.Context) {
	var spotDetail mif.SpotDetail
	_ = c.ShouldBindJSON(&spotDetail)
	if err := sif.UpdateSpotDetail(spotDetail); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags SpotDetail
// @Summary 用id查询SpotDetail
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SpotDetail true "用id查询SpotDetail"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /spotDetail/findSpotDetail [get]
func FindSpotDetail(c *gin.Context) {
	var spotDetail mif.SpotDetail
	_ = c.ShouldBindQuery(&spotDetail)
	if err, respotDetail := sif.GetSpotDetail(spotDetail.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"respotDetail": respotDetail}, c)
	}
}

// @Tags SpotDetail
// @Summary 分页获取SpotDetail列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.SpotDetailSearch true "分页获取SpotDetail列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /spotDetail/getSpotDetailList [get]
func GetSpotDetailList(c *gin.Context) {
	var pageInfo rif.SpotDetailSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := sif.GetSpotDetailInfoList(pageInfo); err != nil {
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
