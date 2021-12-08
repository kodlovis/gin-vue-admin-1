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

// @Tags TradeDetail
// @Summary 创建TradeDetail
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TradeDetail true "创建TradeDetail"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /tradeDetail/createTradeDetail [post]
func CreateTradeDetail(c *gin.Context) {
	var tradeDetail mif.TradeDetail
	_ = c.ShouldBindJSON(&tradeDetail)
	if err := sif.CreateTradeDetail(tradeDetail); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags TradeDetail
// @Summary 删除TradeDetail
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TradeDetail true "删除TradeDetail"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /tradeDetail/deleteTradeDetail [delete]
func DeleteTradeDetail(c *gin.Context) {
	var tradeDetail mif.TradeDetail
	_ = c.ShouldBindJSON(&tradeDetail)
	if err := sif.DeleteTradeDetail(tradeDetail); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags TradeDetail
// @Summary 批量删除TradeDetail
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除TradeDetail"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /tradeDetail/deleteTradeDetailByIds [delete]
func DeleteTradeDetailByIds(c *gin.Context) {
	var IDS rif.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := sif.DeleteTradeDetailByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags TradeDetail
// @Summary 更新TradeDetail
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TradeDetail true "更新TradeDetail"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /tradeDetail/updateTradeDetail [put]
func UpdateTradeDetail(c *gin.Context) {
	var tradeDetail mif.TradeDetail
	_ = c.ShouldBindJSON(&tradeDetail)
	if err := sif.UpdateTradeDetail(tradeDetail); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags TradeDetail
// @Summary 用id查询TradeDetail
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TradeDetail true "用id查询TradeDetail"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /tradeDetail/findTradeDetail [get]
func FindTradeDetail(c *gin.Context) {
	var tradeDetail mif.TradeDetail
	_ = c.ShouldBindQuery(&tradeDetail)
	if err, retradeDetail := sif.GetTradeDetail(tradeDetail.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"retradeDetail": retradeDetail}, c)
	}
}

// @Tags TradeDetail
// @Summary 分页获取TradeDetail列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.TradeDetailSearch true "分页获取TradeDetail列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /tradeDetail/getTradeDetailList [get]
func GetTradeDetailList(c *gin.Context) {
	var pageInfo rif.TradeDetailSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := sif.GetTradeDetailInfoList(pageInfo); err != nil {
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
