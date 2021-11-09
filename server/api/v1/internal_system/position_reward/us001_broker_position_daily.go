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

// @Tags BrokerPositionDaily
// @Summary 创建BrokerPositionDaily
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BrokerPositionDaily true "创建BrokerPositionDaily"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /BrokerPositionDaily/createBrokerPositionDaily [post]
func CreateBrokerPositionDaily(c *gin.Context) {
	var BrokerPositionDaily mp.BrokerPositionDaily
	_ = c.ShouldBindJSON(&BrokerPositionDaily)
	if err := sp.CreateBrokerPositionDaily(BrokerPositionDaily); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags BrokerPositionDaily
// @Summary 删除BrokerPositionDaily
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BrokerPositionDaily true "删除BrokerPositionDaily"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /BrokerPositionDaily/deleteBrokerPositionDaily [delete]
func DeleteBrokerPositionDaily(c *gin.Context) {
	var BrokerPositionDaily mp.BrokerPositionDaily
	_ = c.ShouldBindJSON(&BrokerPositionDaily)
	if err := sp.DeleteBrokerPositionDaily(BrokerPositionDaily); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags BrokerPositionDaily
// @Summary 批量删除BrokerPositionDaily
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除BrokerPositionDaily"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /BrokerPositionDaily/deleteBrokerPositionDailyByIds [delete]
func DeleteBrokerPositionDailyByIds(c *gin.Context) {
	var IDS rp.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := sp.DeleteBrokerPositionDailyByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags BrokerPositionDaily
// @Summary 更新BrokerPositionDaily
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BrokerPositionDaily true "更新BrokerPositionDaily"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /BrokerPositionDaily/updateBrokerPositionDaily [put]
func UpdateBrokerPositionDaily(c *gin.Context) {
	var BrokerPositionDaily mp.BrokerPositionDaily
	_ = c.ShouldBindJSON(&BrokerPositionDaily)
	if err := sp.UpdateBrokerPositionDaily(BrokerPositionDaily); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags BrokerPositionDaily
// @Summary 用id查询BrokerPositionDaily
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BrokerPositionDaily true "用id查询BrokerPositionDaily"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /BrokerPositionDaily/findBrokerPositionDaily [get]
func FindBrokerPositionDaily(c *gin.Context) {
	var BrokerPositionDaily mp.BrokerPositionDaily
	_ = c.ShouldBindQuery(&BrokerPositionDaily)
	if err, reUs001BrokerPositionDaily := sp.GetBrokerPositionDaily(BrokerPositionDaily.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reBrokerPositionDaily": reUs001BrokerPositionDaily}, c)
	}
}

// @Tags BrokerPositionDaily
// @Summary 分页获取BrokerPositionDaily列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.BrokerPositionDailySearch true "分页获取BrokerPositionDaily列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /BrokerPositionDaily/getBrokerPositionDailyList [get]
func GetBrokerPositionDailyList(c *gin.Context) {
	var pageInfo rp.BrokerPositionDailySearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := sp.GetBrokerPositionDailyInfoList(pageInfo); err != nil {
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
