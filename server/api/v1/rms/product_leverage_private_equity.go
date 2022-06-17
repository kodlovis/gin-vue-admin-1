package rms

import (
	"gin-vue-admin/global"
	rp "gin-vue-admin/model/request/rms"
	"gin-vue-admin/model/response"
	mp "gin-vue-admin/model/rms"
	sp "gin-vue-admin/service/rms"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// @Tags ProductLeveragePrivateEquity
// @Summary 创建ProductLeveragePrivateEquity
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ProductLeveragePrivateEquity true "创建ProductLeveragePrivateEquity"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /productLeveragePrivateEquity/createProductLeveragePrivateEquity [post]
func CreateProductLeveragePrivateEquity(c *gin.Context) {
	var productLeveragePrivateEquity mp.ProductLeveragePrivateEquity
	_ = c.ShouldBindJSON(&productLeveragePrivateEquity)
	if err := sp.CreateProductLeveragePrivateEquity(productLeveragePrivateEquity); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags ProductLeveragePrivateEquity
// @Summary 删除ProductLeveragePrivateEquity
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ProductLeveragePrivateEquity true "删除ProductLeveragePrivateEquity"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /productLeveragePrivateEquity/deleteProductLeveragePrivateEquity [delete]
func DeleteProductLeveragePrivateEquity(c *gin.Context) {
	var productLeveragePrivateEquity mp.ProductLeveragePrivateEquity
	_ = c.ShouldBindJSON(&productLeveragePrivateEquity)
	if err := sp.DeleteProductLeveragePrivateEquity(productLeveragePrivateEquity); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags ProductLeveragePrivateEquity
// @Summary 批量删除ProductLeveragePrivateEquity
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ProductLeveragePrivateEquity"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /productLeveragePrivateEquity/deleteProductLeveragePrivateEquityByIds [delete]
func DeleteProductLeveragePrivateEquityByIds(c *gin.Context) {
	var IDS rp.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := sp.DeleteProductLeveragePrivateEquityByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags ProductLeveragePrivateEquity
// @Summary 更新ProductLeveragePrivateEquity
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ProductLeveragePrivateEquity true "更新ProductLeveragePrivateEquity"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /productLeveragePrivateEquity/updateProductLeveragePrivateEquity [put]
func UpdateProductLeveragePrivateEquity(c *gin.Context) {
	var productLeveragePrivateEquity mp.ProductLeveragePrivateEquity
	_ = c.ShouldBindJSON(&productLeveragePrivateEquity)
	if err := sp.UpdateProductLeveragePrivateEquity(productLeveragePrivateEquity); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags ProductLeveragePrivateEquity
// @Summary 用id查询ProductLeveragePrivateEquity
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ProductLeveragePrivateEquity true "用id查询ProductLeveragePrivateEquity"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /productLeveragePrivateEquity/findProductLeveragePrivateEquity [get]
func FindProductLeveragePrivateEquity(c *gin.Context) {
	var productLeveragePrivateEquity mp.ProductLeveragePrivateEquity
	_ = c.ShouldBindQuery(&productLeveragePrivateEquity)
	if err, reproductLeveragePrivateEquity := sp.GetProductLeveragePrivateEquity(productLeveragePrivateEquity.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reproductLeveragePrivateEquity": reproductLeveragePrivateEquity}, c)
	}
}

// @Tags ProductLeveragePrivateEquity
// @Summary 分页获取ProductLeveragePrivateEquity列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.ProductLeveragePrivateEquitySearch true "分页获取ProductLeveragePrivateEquity列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /productLeveragePrivateEquity/getProductLeveragePrivateEquityList [get]
func GetProductLeveragePrivateEquityList(c *gin.Context) {
	var pageInfo rp.ProductLeveragePrivateEquitySearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := sp.GetProductLeveragePrivateEquityInfoList(pageInfo); err != nil {
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
