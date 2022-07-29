package fund

import (
	"gin-vue-admin/global"
	model "gin-vue-admin/model/internal_system/fund"
	request "gin-vue-admin/model/request/internal_system/fund"
	"gin-vue-admin/model/response"
	service "gin-vue-admin/service/internal_system/fund"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// @Tags RemainingSum
// @Summary 创建RemainingSum
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.RemainingSum true "创建RemainingSum"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /remainingSum/createRemainingSum [post]
func CreateRemainingSum(c *gin.Context) {
	var remainingSum model.RemainingSum
	_ = c.ShouldBindJSON(&remainingSum)
	if err := service.CreateRemainingSum(remainingSum); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags RemainingSum
// @Summary 删除RemainingSum
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.RemainingSum true "删除RemainingSum"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /remainingSum/deleteRemainingSum [delete]
func DeleteRemainingSum(c *gin.Context) {
	var remainingSum model.RemainingSum
	_ = c.ShouldBindJSON(&remainingSum)
	if err := service.DeleteRemainingSum(remainingSum); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags RemainingSum
// @Summary 批量删除RemainingSum
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除RemainingSum"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /remainingSum/deleteRemainingSumByIds [delete]
func DeleteRemainingSumByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteRemainingSumByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags RemainingSum
// @Summary 更新RemainingSum
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.RemainingSum true "更新RemainingSum"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /remainingSum/updateRemainingSum [put]
func UpdateRemainingSum(c *gin.Context) {
	var remainingSum model.RemainingSum
	_ = c.ShouldBindJSON(&remainingSum)
	if err := service.UpdateRemainingSum(remainingSum); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags RemainingSum
// @Summary 用id查询RemainingSum
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.RemainingSum true "用id查询RemainingSum"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /remainingSum/findRemainingSum [get]
func FindRemainingSum(c *gin.Context) {
	var remainingSum model.RemainingSum
	_ = c.ShouldBindQuery(&remainingSum)
	if err, reremainingSum := service.GetRemainingSum(remainingSum.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reremainingSum": reremainingSum}, c)
	}
}

// @Tags RemainingSum
// @Summary 分页获取RemainingSum列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.RemainingSumSearch true "分页获取RemainingSum列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /remainingSum/getRemainingSumList [get]
func GetRemainingSumList(c *gin.Context) {
	var pageInfo request.RemainingSumSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetRemainingSumInfoList(pageInfo); err != nil {
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
func GetOneRemainingSum(c *gin.Context) {
	var remainingSum model.SearchRemainingSum
	_ = c.ShouldBindJSON(&remainingSum)
	if err, list := service.SearchRemainingSum(remainingSum); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List: list,
		}, "获取成功", c)
	}
}
func Getwarehouse(c *gin.Context) {
	var remainingSum model.SearchRemainingSum
	_ = c.ShouldBindJSON(&remainingSum)
	if err, list := service.SearchSum(remainingSum); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List: list,
		}, "获取成功", c)
	}
}
func GetFundRole(c *gin.Context) {
	var User model.SearchUser
	_ = c.ShouldBindJSON(&User)
	if err, list := service.SearchUser(User); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.OkWithDetailed(response.PageResult{}, "获取成功", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List: list,
		}, "获取成功", c)
	}
}
