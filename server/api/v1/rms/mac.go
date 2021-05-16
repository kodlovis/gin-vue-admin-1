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

// @Macs Mac
// @Summary 创建Mac
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Mac true "创建Mac"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /Mac/createMac [post]
func CreateMac(c *gin.Context) {
	var Mac mp.Mac
	_ = c.ShouldBindJSON(&Mac)
	if err := sp.CreateMac(Mac); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Macs Mac
// @Summary 删除Mac
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Mac true "删除Mac"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /Mac/deleteMac [delete]
func DeleteMac(c *gin.Context) {
	var Mac mp.Mac
	_ = c.ShouldBindJSON(&Mac)
	if err := sp.DeleteMac(Mac); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Macs Mac
// @Summary 批量删除Mac
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Mac"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /Mac/deleteMacByIds [delete]
func DeleteMacByIds(c *gin.Context) {
	var IDS rp.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := sp.DeleteMacByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Macs Mac
// @Summary 更新Mac
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Mac true "更新Mac"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /Mac/updateMac [put]
func UpdateMac(c *gin.Context) {
	var Mac mp.Mac
	_ = c.ShouldBindJSON(&Mac)
	if err := sp.UpdateMac(&Mac); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Macs Mac
// @Summary 用id查询Mac
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Mac true "用id查询Mac"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /Mac/findMac [get]
func FindMac(c *gin.Context) {
	var Mac mp.Mac
	_ = c.ShouldBindQuery(&Mac)
	if err, reMac := sp.GetMac(Mac.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reMac": reMac}, c)
	}
}

// @Macs Mac
// @Summary 分页获取Mac列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.MacSearch true "分页获取Mac列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /Mac/getMacList [get]
func GetMacList(c *gin.Context) {
	var pageInfo rp.MacSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := sp.GetMacInfoList(pageInfo); err != nil {
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
