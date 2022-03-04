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

// @Tags Us004FutureStockLme
// @Summary 创建Us004FutureStockLme
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Us004FutureStockLme true "创建Us004FutureStockLme"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /us004FutureStockLme/createUs004FutureStockLme [post]
func CreateUs004FutureStockLme(c *gin.Context) {
	var us004FutureStockLme mif.Us004FutureStockLme
	_ = c.ShouldBindJSON(&us004FutureStockLme)
	if err := sif.CreateUs004FutureStockLme(us004FutureStockLme); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags Us004FutureStockLme
// @Summary 删除Us004FutureStockLme
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Us004FutureStockLme true "删除Us004FutureStockLme"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /us004FutureStockLme/deleteUs004FutureStockLme [delete]
func DeleteUs004FutureStockLme(c *gin.Context) {
	var us004FutureStockLme mif.Us004FutureStockLme
	_ = c.ShouldBindJSON(&us004FutureStockLme)
	if err := sif.DeleteUs004FutureStockLme(us004FutureStockLme); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Us004FutureStockLme
// @Summary 批量删除Us004FutureStockLme
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Us004FutureStockLme"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /us004FutureStockLme/deleteUs004FutureStockLmeByIds [delete]
func DeleteUs004FutureStockLmeByIds(c *gin.Context) {
	var IDS rif.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := sif.DeleteUs004FutureStockLmeByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags Us004FutureStockLme
// @Summary 更新Us004FutureStockLme
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Us004FutureStockLme true "更新Us004FutureStockLme"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /us004FutureStockLme/updateUs004FutureStockLme [put]
func UpdateUs004FutureStockLme(c *gin.Context) {
	var us004FutureStockLme mif.Us004FutureStockLme
	_ = c.ShouldBindJSON(&us004FutureStockLme)
	if err := sif.UpdateUs004FutureStockLme(us004FutureStockLme); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags Us004FutureStockLme
// @Summary 用id查询Us004FutureStockLme
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Us004FutureStockLme true "用id查询Us004FutureStockLme"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /us004FutureStockLme/findUs004FutureStockLme [get]
func FindUs004FutureStockLme(c *gin.Context) {
	var us004FutureStockLme mif.Us004FutureStockLme
	_ = c.ShouldBindQuery(&us004FutureStockLme)
	if err, reus004FutureStockLme := sif.GetUs004FutureStockLme(us004FutureStockLme.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reus004FutureStockLme": reus004FutureStockLme}, c)
	}
}

// @Tags Us004FutureStockLme
// @Summary 分页获取Us004FutureStockLme列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.Us004FutureStockLmeSearch true "分页获取Us004FutureStockLme列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /us004FutureStockLme/getUs004FutureStockLmeList [get]
func GetUs004FutureStockLmeList(c *gin.Context) {
	var pageInfo rif.Us004FutureStockLmeSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := sif.GetUs004FutureStockLmeInfoList(pageInfo); err != nil {
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
