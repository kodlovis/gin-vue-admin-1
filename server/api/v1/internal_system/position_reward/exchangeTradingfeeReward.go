package position_reward

import (
	"gin-vue-admin/global"
	mp "gin-vue-admin/model/internal_system/position_reward"
	rp "gin-vue-admin/model/request/internal_system/position_reward"
	"gin-vue-admin/model/response"
	sp "gin-vue-admin/service/internal_system/position_reward"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// @Tags ExchangeTradingfeeReward
// @Summary 创建ExchangeTradingfeeReward
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ExchangeTradingfeeReward true "创建ExchangeTradingfeeReward"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /exchangeTradingfeeReward/createExchangeTradingfeeReward [post]
func CreateExchangeTradingfeeReward(c *gin.Context) {
	var exchangeTradingfeeReward mp.ExchangeTradingfeeReward
	_ = c.ShouldBindJSON(&exchangeTradingfeeReward)
	if err := sp.CreateExchangeTradingfeeReward(exchangeTradingfeeReward); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags ExchangeTradingfeeReward
// @Summary 删除ExchangeTradingfeeReward
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ExchangeTradingfeeReward true "删除ExchangeTradingfeeReward"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /exchangeTradingfeeReward/deleteExchangeTradingfeeReward [delete]
func DeleteExchangeTradingfeeReward(c *gin.Context) {
	var exchangeTradingfeeReward mp.ExchangeTradingfeeReward
	_ = c.ShouldBindJSON(&exchangeTradingfeeReward)
	if err := sp.DeleteExchangeTradingfeeReward(exchangeTradingfeeReward); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags ExchangeTradingfeeReward
// @Summary 批量删除ExchangeTradingfeeReward
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ExchangeTradingfeeReward"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /exchangeTradingfeeReward/deleteExchangeTradingfeeRewardByIds [delete]
func DeleteExchangeTradingfeeRewardByIds(c *gin.Context) {
	var IDS rp.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := sp.DeleteExchangeTradingfeeRewardByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags ExchangeTradingfeeReward
// @Summary 更新ExchangeTradingfeeReward
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ExchangeTradingfeeReward true "更新ExchangeTradingfeeReward"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /exchangeTradingfeeReward/updateExchangeTradingfeeReward [put]
func UpdateExchangeTradingfeeReward(c *gin.Context) {
	var exchangeTradingfeeReward mp.ExchangeTradingfeeReward
	_ = c.ShouldBindJSON(&exchangeTradingfeeReward)
	if err := sp.UpdateExchangeTradingfeeReward(exchangeTradingfeeReward); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags ExchangeTradingfeeReward
// @Summary 用id查询ExchangeTradingfeeReward
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ExchangeTradingfeeReward true "用id查询ExchangeTradingfeeReward"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /exchangeTradingfeeReward/findExchangeTradingfeeReward [get]
func FindExchangeTradingfeeReward(c *gin.Context) {
	var exchangeTradingfeeReward mp.ExchangeTradingfeeReward
	_ = c.ShouldBindQuery(&exchangeTradingfeeReward)
	if err, reexchangeTradingfeeReward := sp.GetExchangeTradingfeeReward(exchangeTradingfeeReward.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reexchangeTradingfeeReward": reexchangeTradingfeeReward}, c)
	}
}

// @Tags ExchangeTradingfeeReward
// @Summary 分页获取ExchangeTradingfeeReward列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.ExchangeTradingfeeRewardSearch true "分页获取ExchangeTradingfeeReward列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /exchangeTradingfeeReward/getExchangeTradingfeeRewardList [get]
func GetExchangeTradingfeeRewardList(c *gin.Context) {
	var pageInfo rp.ExchangeTradingfeeRewardSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := sp.GetExchangeTradingfeeRewardInfoList(pageInfo); err != nil {
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
