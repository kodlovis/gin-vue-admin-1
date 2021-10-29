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

// @Tags ForexFutureDetail
// @Summary 创建ForexFutureDetail
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ForexFutureDetail true "创建ForexFutureDetail"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /forexFutureDetail/createForexFutureDetail [post]
func CreateForexFutureDetail(c *gin.Context) {
	var forexFutureDetail mif.ForexFutureDetail
	_ = c.ShouldBindJSON(&forexFutureDetail)
	if err := sif.CreateForexFutureDetail(forexFutureDetail); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags ForexFutureDetail
// @Summary 删除ForexFutureDetail
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ForexFutureDetail true "删除ForexFutureDetail"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /forexFutureDetail/deleteForexFutureDetail [delete]
func DeleteForexFutureDetail(c *gin.Context) {
	var forexFutureDetail mif.ForexFutureDetail
	_ = c.ShouldBindJSON(&forexFutureDetail)
	if err := sif.DeleteForexFutureDetail(forexFutureDetail); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags ForexFutureDetail
// @Summary 批量删除ForexFutureDetail
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ForexFutureDetail"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /forexFutureDetail/deleteForexFutureDetailByIds [delete]
func DeleteForexFutureDetailByIds(c *gin.Context) {
	var IDS rif.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := sif.DeleteForexFutureDetailByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags ForexFutureDetail
// @Summary 更新ForexFutureDetail
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ForexFutureDetail true "更新ForexFutureDetail"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /forexFutureDetail/updateForexFutureDetail [put]
func UpdateForexFutureDetail(c *gin.Context) {
	var forexFutureDetail mif.ForexFutureDetail
	_ = c.ShouldBindJSON(&forexFutureDetail)
	if err := sif.UpdateForexFutureDetail(forexFutureDetail); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags ForexFutureDetail
// @Summary 用id查询ForexFutureDetail
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ForexFutureDetail true "用id查询ForexFutureDetail"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /forexFutureDetail/findForexFutureDetail [get]
func FindForexFutureDetail(c *gin.Context) {
	var forexFutureDetail mif.ForexFutureDetail
	_ = c.ShouldBindQuery(&forexFutureDetail)
	if err, reforexFutureDetail := sif.GetForexFutureDetail(forexFutureDetail.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reforexFutureDetail": reforexFutureDetail}, c)
	}
}

// @Tags ForexFutureDetail
// @Summary 分页获取ForexFutureDetail列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.ForexFutureDetailSearch true "分页获取ForexFutureDetail列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /forexFutureDetail/getForexFutureDetailList [get]
func GetForexFutureDetailList(c *gin.Context) {
	var pageInfo rif.ForexFutureDetailSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := sif.GetForexFutureDetailInfoList(pageInfo); err != nil {
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
