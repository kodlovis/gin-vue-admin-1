package rms

import (
	"gin-vue-admin/global"
    mp "gin-vue-admin/model/rms"
    rp "gin-vue-admin/model/request/rms"
    "gin-vue-admin/model/response"
    sp "gin-vue-admin/service/rms"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

// @RiskSettings RiskSetting
// @Summary 创建RiskSetting
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.RiskSetting true "创建RiskSetting"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /RiskSetting/createRiskSetting [post]
func CreateRiskSetting(c *gin.Context) {
	var RiskSetting mp.RiskSetting
	_ = c.ShouldBindJSON(&RiskSetting)
	if err := sp.CreateRiskSetting(RiskSetting); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @RiskSettings RiskSetting
// @Summary 删除RiskSetting
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.RiskSetting true "删除RiskSetting"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /RiskSetting/deleteRiskSetting [delete]
func DeleteRiskSetting(c *gin.Context) {
	var RiskSetting mp.RiskSetting
	_ = c.ShouldBindJSON(&RiskSetting)
	if err := sp.DeleteRiskSetting(RiskSetting); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @RiskSettings RiskSetting
// @Summary 批量删除RiskSetting
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除RiskSetting"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /RiskSetting/deleteRiskSettingByIds [delete]
func DeleteRiskSettingByIds(c *gin.Context) {
	var IDS rp.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := sp.DeleteRiskSettingByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @RiskSettings RiskSetting
// @Summary 更新RiskSetting
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.RiskSetting true "更新RiskSetting"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /RiskSetting/updateRiskSetting [put]
func UpdateRiskSetting(c *gin.Context) {
	var RiskSetting mp.RiskSetting
	_ = c.ShouldBindJSON(&RiskSetting)
	if err := sp.UpdateRiskSetting(&RiskSetting); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @RiskSettings RiskSetting
// @Summary 用id查询RiskSetting
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.RiskSetting true "用id查询RiskSetting"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /RiskSetting/findRiskSetting [get]
func FindRiskSetting(c *gin.Context) {
	var RiskSetting mp.RiskSetting
	_ = c.ShouldBindQuery(&RiskSetting)
	if err, reRiskSetting := sp.GetRiskSetting(RiskSetting.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reRiskSetting": reRiskSetting}, c)
	}
}

// @RiskSettings RiskSetting
// @Summary 分页获取RiskSetting列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.RiskSettingSearch true "分页获取RiskSetting列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /RiskSetting/getRiskSettingList [get]
func GetRiskSettingList(c *gin.Context) {
	var pageInfo rp.RiskSettingSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := sp.GetRiskSettingInfoList(pageInfo); err != nil {
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
//不使用pageinfo
func GetRSIList(c *gin.Context) {
	var pageInfo rp.RiskSettingSearch
	if err, list, total := sp.GetRSIList(); err != nil {
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
