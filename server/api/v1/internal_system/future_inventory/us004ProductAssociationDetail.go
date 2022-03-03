package future_inventory

import (
	"gin-vue-admin/global"
	mif "gin-vue-admin/model/internal_system/future_inventory"
	rif "gin-vue-admin/model/request/internal_system/future_inventory"
	"gin-vue-admin/model/response"
	sif "gin-vue-admin/service/internal_system/future_inventory"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// @Tags Us004ProductAssociationDetail
// @Summary 创建Us004ProductAssociationDetail
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Us004ProductAssociationDetail true "创建Us004ProductAssociationDetail"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /us004ProductAssociationDetail/createUs004ProductAssociationDetail [post]
func CreateUs004ProductAssociationDetail(c *gin.Context) {
	var us004ProductAssociationDetail mif.Us004ProductAssociationDetail
	_ = c.ShouldBindJSON(&us004ProductAssociationDetail)
	if err := sif.CreateUs004ProductAssociationDetail(us004ProductAssociationDetail); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags Us004ProductAssociationDetail
// @Summary 删除Us004ProductAssociationDetail
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Us004ProductAssociationDetail true "删除Us004ProductAssociationDetail"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /us004ProductAssociationDetail/deleteUs004ProductAssociationDetail [delete]
func DeleteUs004ProductAssociationDetail(c *gin.Context) {
	var us004ProductAssociationDetail mif.Us004ProductAssociationDetail
	_ = c.ShouldBindJSON(&us004ProductAssociationDetail)
	if err := sif.DeleteUs004ProductAssociationDetail(us004ProductAssociationDetail); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Us004ProductAssociationDetail
// @Summary 批量删除Us004ProductAssociationDetail
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Us004ProductAssociationDetail"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /us004ProductAssociationDetail/deleteUs004ProductAssociationDetailByIds [delete]
func DeleteUs004ProductAssociationDetailByIds(c *gin.Context) {
	var IDS rif.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := sif.DeleteUs004ProductAssociationDetailByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags Us004ProductAssociationDetail
// @Summary 更新Us004ProductAssociationDetail
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Us004ProductAssociationDetail true "更新Us004ProductAssociationDetail"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /us004ProductAssociationDetail/updateUs004ProductAssociationDetail [put]
func UpdateUs004ProductAssociationDetail(c *gin.Context) {
	var us004ProductAssociationDetail mif.Us004ProductAssociationDetail
	_ = c.ShouldBindJSON(&us004ProductAssociationDetail)
	if err := sif.UpdateUs004ProductAssociationDetail(us004ProductAssociationDetail); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags Us004ProductAssociationDetail
// @Summary 用id查询Us004ProductAssociationDetail
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Us004ProductAssociationDetail true "用id查询Us004ProductAssociationDetail"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /us004ProductAssociationDetail/findUs004ProductAssociationDetail [get]
func FindUs004ProductAssociationDetail(c *gin.Context) {
	var us004ProductAssociationDetail mif.Us004ProductAssociationDetail
	_ = c.ShouldBindQuery(&us004ProductAssociationDetail)
	if err, reus004ProductAssociationDetail := sif.GetUs004ProductAssociationDetail(us004ProductAssociationDetail.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reus004ProductAssociationDetail": reus004ProductAssociationDetail}, c)
	}
}

// @Tags Us004ProductAssociationDetail
// @Summary 分页获取Us004ProductAssociationDetail列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.Us004ProductAssociationDetailSearch true "分页获取Us004ProductAssociationDetail列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /us004ProductAssociationDetail/getUs004ProductAssociationDetailList [get]
func GetUs004ProductAssociationDetailList(c *gin.Context) {
	var pageInfo rif.Us004ProductAssociationDetailSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := sif.GetUs004ProductAssociationDetailInfoList(pageInfo); err != nil {
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
