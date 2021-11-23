package internal_system

import (
	"gin-vue-admin/global"
	mp "gin-vue-admin/model/internal_system"
	rp "gin-vue-admin/model/request/internal_system"
	"gin-vue-admin/model/response"
	sp "gin-vue-admin/service/internal_system"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// @Tags ExchangeProductInfo
// @Summary 创建ExchangeProductInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ExchangeProductInfo true "创建ExchangeProductInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /exchangeProductInfo/createExchangeProductInfo [post]
func CreateExchangeProductInfo(c *gin.Context) {
	var exchangeProductInfo mp.ExchangeProductInfo
	_ = c.ShouldBindJSON(&exchangeProductInfo)
	if err := sp.CreateExchangeProductInfo(exchangeProductInfo); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags ExchangeProductInfo
// @Summary 删除ExchangeProductInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ExchangeProductInfo true "删除ExchangeProductInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /exchangeProductInfo/deleteExchangeProductInfo [delete]
func DeleteExchangeProductInfo(c *gin.Context) {
	var exchangeProductInfo mp.ExchangeProductInfo
	_ = c.ShouldBindJSON(&exchangeProductInfo)
	if err := sp.DeleteExchangeProductInfo(exchangeProductInfo); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags ExchangeProductInfo
// @Summary 批量删除ExchangeProductInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ExchangeProductInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /exchangeProductInfo/deleteExchangeProductInfoByIds [delete]
func DeleteExchangeProductInfoByIds(c *gin.Context) {
	var IDS rp.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := sp.DeleteExchangeProductInfoByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags ExchangeProductInfo
// @Summary 更新ExchangeProductInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ExchangeProductInfo true "更新ExchangeProductInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /exchangeProductInfo/updateExchangeProductInfo [put]
func UpdateExchangeProductInfo(c *gin.Context) {
	var exchangeProductInfo mp.ExchangeProductInfo
	_ = c.ShouldBindJSON(&exchangeProductInfo)
	if err := sp.UpdateExchangeProductInfo(exchangeProductInfo); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags ExchangeProductInfo
// @Summary 用id查询ExchangeProductInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ExchangeProductInfo true "用id查询ExchangeProductInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /exchangeProductInfo/findExchangeProductInfo [get]
func FindExchangeProductInfo(c *gin.Context) {
	var exchangeProductInfo mp.ExchangeProductInfo
	_ = c.ShouldBindQuery(&exchangeProductInfo)
	if err, reexchangeProductInfo := sp.GetExchangeProductInfo(exchangeProductInfo.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reexchangeProductInfo": reexchangeProductInfo}, c)
	}
}

// @Tags ExchangeProductInfo
// @Summary 分页获取ExchangeProductInfo列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.ExchangeProductInfoSearch true "分页获取ExchangeProductInfo列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /exchangeProductInfo/getExchangeProductInfoList [get]
func GetExchangeProductInfoList(c *gin.Context) {
	var pageInfo rp.ExchangeProductInfoSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := sp.GetExchangeProductInfoInfoList(pageInfo); err != nil {
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
