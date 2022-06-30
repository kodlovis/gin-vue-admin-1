package future_data

import (
	"gin-vue-admin/global"
	model "gin-vue-admin/model/internal_system/future_data"
	request "gin-vue-admin/model/request/internal_system/future_data"
	"gin-vue-admin/model/response"
	service "gin-vue-admin/service/internal_system/future_data"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// @Tags Us002ForeignExchangeProfit
// @Summary 创建Us002ForeignExchangeProfit
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Us002ForeignExchangeProfit true "创建Us002ForeignExchangeProfit"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /us002ForeignExchangeProfit/createUs002ForeignExchangeProfit [post]
func CreateUs002ForeignExchangeProfit(c *gin.Context) {
	var us002ForeignExchangeProfit model.Us002ForeignExchangeProfit
	_ = c.ShouldBindJSON(&us002ForeignExchangeProfit)
	if err := service.CreateUs002ForeignExchangeProfit(us002ForeignExchangeProfit); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags Us002ForeignExchangeProfit
// @Summary 删除Us002ForeignExchangeProfit
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Us002ForeignExchangeProfit true "删除Us002ForeignExchangeProfit"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /us002ForeignExchangeProfit/deleteUs002ForeignExchangeProfit [delete]
func DeleteUs002ForeignExchangeProfit(c *gin.Context) {
	var us002ForeignExchangeProfit model.Us002ForeignExchangeProfit
	_ = c.ShouldBindJSON(&us002ForeignExchangeProfit)
	if err := service.DeleteUs002ForeignExchangeProfit(us002ForeignExchangeProfit); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Us002ForeignExchangeProfit
// @Summary 批量删除Us002ForeignExchangeProfit
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Us002ForeignExchangeProfit"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /us002ForeignExchangeProfit/deleteUs002ForeignExchangeProfitByIds [delete]
func DeleteUs002ForeignExchangeProfitByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteUs002ForeignExchangeProfitByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags Us002ForeignExchangeProfit
// @Summary 更新Us002ForeignExchangeProfit
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Us002ForeignExchangeProfit true "更新Us002ForeignExchangeProfit"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /us002ForeignExchangeProfit/updateUs002ForeignExchangeProfit [put]
func UpdateUs002ForeignExchangeProfit(c *gin.Context) {
	var us002ForeignExchangeProfit model.Us002ForeignExchangeProfit
	_ = c.ShouldBindJSON(&us002ForeignExchangeProfit)
	if err := service.UpdateUs002ForeignExchangeProfit(us002ForeignExchangeProfit); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags Us002ForeignExchangeProfit
// @Summary 用id查询Us002ForeignExchangeProfit
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Us002ForeignExchangeProfit true "用id查询Us002ForeignExchangeProfit"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /us002ForeignExchangeProfit/findUs002ForeignExchangeProfit [get]
func FindUs002ForeignExchangeProfit(c *gin.Context) {
	var us002ForeignExchangeProfit model.Us002ForeignExchangeProfit
	_ = c.ShouldBindQuery(&us002ForeignExchangeProfit)
	if err, reus002ForeignExchangeProfit := service.GetUs002ForeignExchangeProfit(us002ForeignExchangeProfit.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reus002ForeignExchangeProfit": reus002ForeignExchangeProfit}, c)
	}
}

// @Tags Us002ForeignExchangeProfit
// @Summary 分页获取Us002ForeignExchangeProfit列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.Us002ForeignExchangeProfitSearch true "分页获取Us002ForeignExchangeProfit列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /us002ForeignExchangeProfit/getUs002ForeignExchangeProfitList [get]
func GetUs002ForeignExchangeProfitList(c *gin.Context) {
	var pageInfo request.Us002ForeignExchangeProfitSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetUs002ForeignExchangeProfitInfoList(pageInfo); err != nil {
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
