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

// @Tags Us002FutureBasisOfInstrumentClose
// @Summary 创建Us002FutureBasisOfInstrumentClose
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Us002FutureBasisOfInstrumentClose true "创建Us002FutureBasisOfInstrumentClose"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /us002FutureBasisOfInstrumentClose/createUs002FutureBasisOfInstrumentClose [post]
func CreateUs002FutureBasisOfInstrumentClose(c *gin.Context) {
	var us002FutureBasisOfInstrumentClose model.Us002FutureBasisOfInstrumentClose
	_ = c.ShouldBindJSON(&us002FutureBasisOfInstrumentClose)
	if err := service.CreateUs002FutureBasisOfInstrumentClose(us002FutureBasisOfInstrumentClose); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags Us002FutureBasisOfInstrumentClose
// @Summary 删除Us002FutureBasisOfInstrumentClose
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Us002FutureBasisOfInstrumentClose true "删除Us002FutureBasisOfInstrumentClose"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /us002FutureBasisOfInstrumentClose/deleteUs002FutureBasisOfInstrumentClose [delete]
func DeleteUs002FutureBasisOfInstrumentClose(c *gin.Context) {
	var us002FutureBasisOfInstrumentClose model.Us002FutureBasisOfInstrumentClose
	_ = c.ShouldBindJSON(&us002FutureBasisOfInstrumentClose)
	if err := service.DeleteUs002FutureBasisOfInstrumentClose(us002FutureBasisOfInstrumentClose); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Us002FutureBasisOfInstrumentClose
// @Summary 批量删除Us002FutureBasisOfInstrumentClose
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Us002FutureBasisOfInstrumentClose"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /us002FutureBasisOfInstrumentClose/deleteUs002FutureBasisOfInstrumentCloseByIds [delete]
func DeleteUs002FutureBasisOfInstrumentCloseByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteUs002FutureBasisOfInstrumentCloseByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags Us002FutureBasisOfInstrumentClose
// @Summary 更新Us002FutureBasisOfInstrumentClose
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Us002FutureBasisOfInstrumentClose true "更新Us002FutureBasisOfInstrumentClose"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /us002FutureBasisOfInstrumentClose/updateUs002FutureBasisOfInstrumentClose [put]
func UpdateUs002FutureBasisOfInstrumentClose(c *gin.Context) {
	var us002FutureBasisOfInstrumentClose model.Us002FutureBasisOfInstrumentClose
	_ = c.ShouldBindJSON(&us002FutureBasisOfInstrumentClose)
	if err := service.UpdateUs002FutureBasisOfInstrumentClose(us002FutureBasisOfInstrumentClose); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags Us002FutureBasisOfInstrumentClose
// @Summary 用id查询Us002FutureBasisOfInstrumentClose
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Us002FutureBasisOfInstrumentClose true "用id查询Us002FutureBasisOfInstrumentClose"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /us002FutureBasisOfInstrumentClose/findUs002FutureBasisOfInstrumentClose [get]
func FindUs002FutureBasisOfInstrumentClose(c *gin.Context) {
	var us002FutureBasisOfInstrumentClose model.Us002FutureBasisOfInstrumentClose
	_ = c.ShouldBindQuery(&us002FutureBasisOfInstrumentClose)
	if err, reus002FutureBasisOfInstrumentClose := service.GetUs002FutureBasisOfInstrumentClose(us002FutureBasisOfInstrumentClose.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reus002FutureBasisOfInstrumentClose": reus002FutureBasisOfInstrumentClose}, c)
	}
}

// @Tags Us002FutureBasisOfInstrumentClose
// @Summary 分页获取Us002FutureBasisOfInstrumentClose列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.Us002FutureBasisOfInstrumentCloseSearch true "分页获取Us002FutureBasisOfInstrumentClose列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /us002FutureBasisOfInstrumentClose/getUs002FutureBasisOfInstrumentCloseList [get]
func GetUs002FutureBasisOfInstrumentCloseList(c *gin.Context) {
	var pageInfo request.Us002FutureBasisOfInstrumentCloseSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetUs002FutureBasisOfInstrumentCloseInfoList(pageInfo); err != nil {
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

//获取交割月合约信息
func GetPositionDeliveryMonthInstrumentList(c *gin.Context) {
	var pageInfo request.PositionDeliveryMonthInstrument
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.PositionDeliveryMonthInstrumentlIST(pageInfo); err != nil {
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
func GetBenchmarkInstrumentList(c *gin.Context) {
	var pageInfo request.PositionDeliveryMonthInstrument
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetBenchmarkInstrumentList(pageInfo); err != nil {
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

//计算无风险价差
func GetNoRiskValue(c *gin.Context) {
	var noRiskValue request.NoRiskValue
	_ = c.ShouldBindQuery(&noRiskValue)
	if err, regetNoRiskValue := service.GetNoRiskValue(noRiskValue); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"regetNoRiskValue": regetNoRiskValue}, c)
	}
}

//批量创建合约信息
func CreateBasisOfInstrumentCloseList(c *gin.Context) {
	var us002FutureBasisOfInstrumentClose request.BasisOInstrumentList
	_ = c.ShouldBindJSON(&us002FutureBasisOfInstrumentClose)
	if err := service.CreateBasisOfInstrumentCloseList(us002FutureBasisOfInstrumentClose); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}
