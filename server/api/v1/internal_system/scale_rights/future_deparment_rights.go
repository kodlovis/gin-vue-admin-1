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

// @Tags FutureDeparmentRights
// @Summary 创建FutureDeparmentRights
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.FutureDeparmentRights true "创建FutureDeparmentRights"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /futureDeparmentRights/createFutureDeparmentRights [post]
func CreateFutureDeparmentRights(c *gin.Context) {
	var futureDeparmentRights mp.FutureDeparmentRights
	_ = c.ShouldBindJSON(&futureDeparmentRights)
	if err := sp.CreateFutureDeparmentRights(futureDeparmentRights); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags FutureDeparmentRights
// @Summary 删除FutureDeparmentRights
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.FutureDeparmentRights true "删除FutureDeparmentRights"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /futureDeparmentRights/deleteFutureDeparmentRights [delete]
func DeleteFutureDeparmentRights(c *gin.Context) {
	var futureDeparmentRights mp.FutureDeparmentRights
	_ = c.ShouldBindJSON(&futureDeparmentRights)
	if err := sp.DeleteFutureDeparmentRights(futureDeparmentRights); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags FutureDeparmentRights
// @Summary 批量删除FutureDeparmentRights
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除FutureDeparmentRights"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /futureDeparmentRights/deleteFutureDeparmentRightsByIds [delete]
func DeleteFutureDeparmentRightsByIds(c *gin.Context) {
	var IDS rp.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := sp.DeleteFutureDeparmentRightsByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags FutureDeparmentRights
// @Summary 更新FutureDeparmentRights
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.FutureDeparmentRights true "更新FutureDeparmentRights"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /futureDeparmentRights/updateFutureDeparmentRights [put]
func UpdateFutureDeparmentRights(c *gin.Context) {
	var futureDeparmentRights mp.FutureDeparmentRights
	_ = c.ShouldBindJSON(&futureDeparmentRights)
	if err := sp.UpdateFutureDeparmentRights(futureDeparmentRights); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags FutureDeparmentRights
// @Summary 用id查询FutureDeparmentRights
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.FutureDeparmentRights true "用id查询FutureDeparmentRights"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /futureDeparmentRights/findFutureDeparmentRights [get]
func FindFutureDeparmentRights(c *gin.Context) {
	var futureDeparmentRights mp.FutureDeparmentRights
	_ = c.ShouldBindQuery(&futureDeparmentRights)
	if err, refutureDeparmentRights := sp.GetFutureDeparmentRights(futureDeparmentRights.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"refutureDeparmentRights": refutureDeparmentRights}, c)
	}
}

// @Tags FutureDeparmentRights
// @Summary 分页获取FutureDeparmentRights列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.FutureDeparmentRightsSearch true "分页获取FutureDeparmentRights列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /futureDeparmentRights/getFutureDeparmentRightsList [get]
func GetFutureDeparmentRightsList(c *gin.Context) {
	var pageInfo rp.FutureDeparmentRightsSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := sp.GetFutureDeparmentRightsInfoList(pageInfo); err != nil {
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
