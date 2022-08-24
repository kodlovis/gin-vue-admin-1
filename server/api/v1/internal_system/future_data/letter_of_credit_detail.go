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

// @Tags LetterOfCreditDetail
// @Summary 创建LetterOfCreditDetail
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.LetterOfCreditDetail true "创建LetterOfCreditDetail"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /letterOfCreditDetail/createLetterOfCreditDetail [post]
func CreateLetterOfCreditDetail(c *gin.Context) {
	var letterOfCreditDetail model.LetterOfCreditDetail
	_ = c.ShouldBindJSON(&letterOfCreditDetail)
	if err := service.CreateLetterOfCreditDetail(letterOfCreditDetail); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags LetterOfCreditDetail
// @Summary 删除LetterOfCreditDetail
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.LetterOfCreditDetail true "删除LetterOfCreditDetail"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /letterOfCreditDetail/deleteLetterOfCreditDetail [delete]
func DeleteLetterOfCreditDetail(c *gin.Context) {
	var letterOfCreditDetail model.LetterOfCreditDetail
	_ = c.ShouldBindJSON(&letterOfCreditDetail)
	if err := service.DeleteLetterOfCreditDetail(letterOfCreditDetail); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags LetterOfCreditDetail
// @Summary 批量删除LetterOfCreditDetail
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除LetterOfCreditDetail"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /letterOfCreditDetail/deleteLetterOfCreditDetailByIds [delete]
func DeleteLetterOfCreditDetailByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteLetterOfCreditDetailByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags LetterOfCreditDetail
// @Summary 更新LetterOfCreditDetail
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.LetterOfCreditDetail true "更新LetterOfCreditDetail"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /letterOfCreditDetail/updateLetterOfCreditDetail [put]
func UpdateLetterOfCreditDetail(c *gin.Context) {
	var letterOfCreditDetail model.LetterOfCreditDetail
	_ = c.ShouldBindJSON(&letterOfCreditDetail)
	if err := service.UpdateLetterOfCreditDetail(letterOfCreditDetail); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags LetterOfCreditDetail
// @Summary 用id查询LetterOfCreditDetail
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.LetterOfCreditDetail true "用id查询LetterOfCreditDetail"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /letterOfCreditDetail/findLetterOfCreditDetail [get]
func FindLetterOfCreditDetail(c *gin.Context) {
	var letterOfCreditDetail model.LetterOfCreditDetail
	_ = c.ShouldBindQuery(&letterOfCreditDetail)
	if err, reletterOfCreditDetail := service.GetLetterOfCreditDetail(letterOfCreditDetail.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reletterOfCreditDetail": reletterOfCreditDetail}, c)
	}
}

// @Tags LetterOfCreditDetail
// @Summary 分页获取LetterOfCreditDetail列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.LetterOfCreditDetailSearch true "分页获取LetterOfCreditDetail列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /letterOfCreditDetail/getLetterOfCreditDetailList [get]
func GetLetterOfCreditDetailList(c *gin.Context) {
	var pageInfo request.LetterOfCreditDetailSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetLetterOfCreditDetailInfoList(pageInfo); err != nil {
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

//获取未填写购置汇率的信用证
func GetLetterOfCreditDetailListWithNoPurchaseRate(c *gin.Context) {
	var pageInfo request.LetterOfCreditDetailSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetLetterOfCreditDetailListWithNoPurchaseRate(pageInfo); err != nil {
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
func UpdateLetterOfCreditPurchaseRate(c *gin.Context) {
	var letterOfCreditDetail model.LetterOfCreditDetail
	_ = c.ShouldBindJSON(&letterOfCreditDetail)
	if err := service.UpdateLetterOfCreditPurchaseRate(letterOfCreditDetail); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}
