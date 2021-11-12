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

// @Tags FutureDepartmentScaleDetail
// @Summary 创建FutureDepartmentScaleDetail
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.FutureDepartmentScaleDetail true "创建FutureDepartmentScaleDetail"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /futureDepartmentScaleDetail/createFutureDepartmentScaleDetail [post]
func CreateFutureDepartmentScaleDetail(c *gin.Context) {
	var futureDepartmentScaleDetail mp.FutureDepartmentScaleDetail
	_ = c.ShouldBindJSON(&futureDepartmentScaleDetail)
	if err := sp.CreateFutureDepartmentScaleDetail(futureDepartmentScaleDetail); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags FutureDepartmentScaleDetail
// @Summary 删除FutureDepartmentScaleDetail
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.FutureDepartmentScaleDetail true "删除FutureDepartmentScaleDetail"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /futureDepartmentScaleDetail/deleteFutureDepartmentScaleDetail [delete]
func DeleteFutureDepartmentScaleDetail(c *gin.Context) {
	var futureDepartmentScaleDetail mp.FutureDepartmentScaleDetail
	_ = c.ShouldBindJSON(&futureDepartmentScaleDetail)
	if err := sp.DeleteFutureDepartmentScaleDetail(futureDepartmentScaleDetail); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags FutureDepartmentScaleDetail
// @Summary 批量删除FutureDepartmentScaleDetail
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除FutureDepartmentScaleDetail"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /futureDepartmentScaleDetail/deleteFutureDepartmentScaleDetailByIds [delete]
func DeleteFutureDepartmentScaleDetailByIds(c *gin.Context) {
	var IDS rp.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := sp.DeleteFutureDepartmentScaleDetailByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags FutureDepartmentScaleDetail
// @Summary 更新FutureDepartmentScaleDetail
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.FutureDepartmentScaleDetail true "更新FutureDepartmentScaleDetail"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /futureDepartmentScaleDetail/updateFutureDepartmentScaleDetail [put]
func UpdateFutureDepartmentScaleDetail(c *gin.Context) {
	var futureDepartmentScaleDetail mp.FutureDepartmentScaleDetail
	_ = c.ShouldBindJSON(&futureDepartmentScaleDetail)
	if err := sp.UpdateFutureDepartmentScaleDetail(futureDepartmentScaleDetail); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags FutureDepartmentScaleDetail
// @Summary 用id查询FutureDepartmentScaleDetail
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.FutureDepartmentScaleDetail true "用id查询FutureDepartmentScaleDetail"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /futureDepartmentScaleDetail/findFutureDepartmentScaleDetail [get]
func FindFutureDepartmentScaleDetail(c *gin.Context) {
	var futureDepartmentScaleDetail mp.FutureDepartmentScaleDetail
	_ = c.ShouldBindQuery(&futureDepartmentScaleDetail)
	if err, refutureDepartmentScaleDetail := sp.GetFutureDepartmentScaleDetail(futureDepartmentScaleDetail.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"refutureDepartmentScaleDetail": refutureDepartmentScaleDetail}, c)
	}
}

// @Tags FutureDepartmentScaleDetail
// @Summary 分页获取FutureDepartmentScaleDetail列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.FutureDepartmentScaleDetailSearch true "分页获取FutureDepartmentScaleDetail列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /futureDepartmentScaleDetail/getFutureDepartmentScaleDetailList [get]
func GetFutureDepartmentScaleDetailList(c *gin.Context) {
	var pageInfo rp.FutureDepartmentScaleDetailSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := sp.GetFutureDepartmentScaleDetailInfoList(pageInfo); err != nil {
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
