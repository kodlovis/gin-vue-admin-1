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

// @Tags PositionRewardRule
// @Summary 创建PositionRewardRule
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PositionRewardRule true "创建PositionRewardRule"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /positionRewardRule/createPositionRewardRule [post]
func CreatePositionRewardRule(c *gin.Context) {
	var positionRewardRule mp.PositionRewardRule
	_ = c.ShouldBindJSON(&positionRewardRule)
	if err := sp.CreatePositionRewardRule(positionRewardRule); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags PositionRewardRule
// @Summary 删除PositionRewardRule
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PositionRewardRule true "删除PositionRewardRule"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /positionRewardRule/deletePositionRewardRule [delete]
func DeletePositionRewardRule(c *gin.Context) {
	var positionRewardRule mp.PositionRewardRule
	_ = c.ShouldBindJSON(&positionRewardRule)
	if err := sp.DeletePositionRewardRule(positionRewardRule); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags PositionRewardRule
// @Summary 批量删除PositionRewardRule
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除PositionRewardRule"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /positionRewardRule/deletePositionRewardRuleByIds [delete]
func DeletePositionRewardRuleByIds(c *gin.Context) {
	var IDS rp.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := sp.DeletePositionRewardRuleByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags PositionRewardRule
// @Summary 更新PositionRewardRule
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PositionRewardRule true "更新PositionRewardRule"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /positionRewardRule/updatePositionRewardRule [put]
func UpdatePositionRewardRule(c *gin.Context) {
	var positionRewardRule mp.PositionRewardRule
	_ = c.ShouldBindJSON(&positionRewardRule)
	if err := sp.UpdatePositionRewardRule(positionRewardRule); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags PositionRewardRule
// @Summary 用id查询PositionRewardRule
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PositionRewardRule true "用id查询PositionRewardRule"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /positionRewardRule/findPositionRewardRule [get]
func FindPositionRewardRule(c *gin.Context) {
	var positionRewardRule mp.PositionRewardRule
	_ = c.ShouldBindQuery(&positionRewardRule)
	if err, repositionRewardRule := sp.GetPositionRewardRule(positionRewardRule.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"repositionRewardRule": repositionRewardRule}, c)
	}
}

// @Tags PositionRewardRule
// @Summary 分页获取PositionRewardRule列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PositionRewardRuleSearch true "分页获取PositionRewardRule列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /positionRewardRule/getPositionRewardRuleList [get]
func GetPositionRewardRuleList(c *gin.Context) {
	var pageInfo rp.PositionRewardRuleSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := sp.GetPositionRewardRuleInfoList(pageInfo); err != nil {
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
