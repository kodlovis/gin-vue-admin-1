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

// @Tags Us004FutureInventoryDaily
// @Summary 创建Us004FutureInventoryDaily
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Us004FutureInventoryDaily true "创建Us004FutureInventoryDaily"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /us004FutureInventoryDaily/createUs004FutureInventoryDaily [post]
func CreateUs004FutureInventoryDaily(c *gin.Context) {
	var us004FutureInventoryDaily mif.Us004FutureInventoryDaily
	_ = c.ShouldBindJSON(&us004FutureInventoryDaily)
	if err := sif.CreateUs004FutureInventoryDaily(us004FutureInventoryDaily); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags Us004FutureInventoryDaily
// @Summary 删除Us004FutureInventoryDaily
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Us004FutureInventoryDaily true "删除Us004FutureInventoryDaily"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /us004FutureInventoryDaily/deleteUs004FutureInventoryDaily [delete]
func DeleteUs004FutureInventoryDaily(c *gin.Context) {
	var us004FutureInventoryDaily mif.Us004FutureInventoryDaily
	_ = c.ShouldBindJSON(&us004FutureInventoryDaily)
	if err := sif.DeleteUs004FutureInventoryDaily(us004FutureInventoryDaily); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Us004FutureInventoryDaily
// @Summary 批量删除Us004FutureInventoryDaily
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Us004FutureInventoryDaily"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /us004FutureInventoryDaily/deleteUs004FutureInventoryDailyByIds [delete]
func DeleteUs004FutureInventoryDailyByIds(c *gin.Context) {
	var IDS rif.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := sif.DeleteUs004FutureInventoryDailyByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags Us004FutureInventoryDaily
// @Summary 更新Us004FutureInventoryDaily
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Us004FutureInventoryDaily true "更新Us004FutureInventoryDaily"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /us004FutureInventoryDaily/updateUs004FutureInventoryDaily [put]
func UpdateUs004FutureInventoryDaily(c *gin.Context) {
	var us004FutureInventoryDaily mif.Us004FutureInventoryDaily
	_ = c.ShouldBindJSON(&us004FutureInventoryDaily)
	if err := sif.UpdateUs004FutureInventoryDaily(us004FutureInventoryDaily); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags Us004FutureInventoryDaily
// @Summary 用id查询Us004FutureInventoryDaily
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Us004FutureInventoryDaily true "用id查询Us004FutureInventoryDaily"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /us004FutureInventoryDaily/findUs004FutureInventoryDaily [get]
func FindUs004FutureInventoryDaily(c *gin.Context) {
	var us004FutureInventoryDaily mif.Us004FutureInventoryDaily
	_ = c.ShouldBindQuery(&us004FutureInventoryDaily)
	if err, reus004FutureInventoryDaily := sif.GetUs004FutureInventoryDaily(us004FutureInventoryDaily.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reus004FutureInventoryDaily": reus004FutureInventoryDaily}, c)
	}
}

// @Tags Us004FutureInventoryDaily
// @Summary 分页获取Us004FutureInventoryDaily列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.Us004FutureInventoryDailySearch true "分页获取Us004FutureInventoryDaily列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /us004FutureInventoryDaily/getUs004FutureInventoryDailyList [get]
func GetUs004FutureInventoryDailyList(c *gin.Context) {
	var pageInfo rif.Us004FutureInventoryDailySearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := sif.GetUs004FutureInventoryDailyInfoList(pageInfo); err != nil {
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

func GetUs004FutureInventoryDailyType(c *gin.Context) {
	var pageInfo rif.Us004FutureInventoryDailySearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := sif.GetUs004FutureInventoryDailyType(pageInfo); err != nil {
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
