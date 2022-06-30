package public

import (
	"gin-vue-admin/global"
	model "gin-vue-admin/model/internal_system/public"
	request "gin-vue-admin/model/request/internal_system/public"
	"gin-vue-admin/model/response"
	service "gin-vue-admin/service/internal_system/public"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// @Tags Us100Currency
// @Summary 创建Us100Currency
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Us100Currency true "创建Us100Currency"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /us100Currency/createUs100Currency [post]
func CreateUs100Currency(c *gin.Context) {
	var us100Currency model.Us100Currency
	_ = c.ShouldBindJSON(&us100Currency)
	if err := service.CreateUs100Currency(us100Currency); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags Us100Currency
// @Summary 删除Us100Currency
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Us100Currency true "删除Us100Currency"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /us100Currency/deleteUs100Currency [delete]
func DeleteUs100Currency(c *gin.Context) {
	var us100Currency model.Us100Currency
	_ = c.ShouldBindJSON(&us100Currency)
	if err := service.DeleteUs100Currency(us100Currency); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Us100Currency
// @Summary 批量删除Us100Currency
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Us100Currency"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /us100Currency/deleteUs100CurrencyByIds [delete]
func DeleteUs100CurrencyByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteUs100CurrencyByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags Us100Currency
// @Summary 更新Us100Currency
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Us100Currency true "更新Us100Currency"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /us100Currency/updateUs100Currency [put]
func UpdateUs100Currency(c *gin.Context) {
	var us100Currency model.Us100Currency
	_ = c.ShouldBindJSON(&us100Currency)
	if err := service.UpdateUs100Currency(us100Currency); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags Us100Currency
// @Summary 用id查询Us100Currency
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Us100Currency true "用id查询Us100Currency"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /us100Currency/findUs100Currency [get]
func FindUs100Currency(c *gin.Context) {
	var us100Currency model.Us100Currency
	_ = c.ShouldBindQuery(&us100Currency)
	if err, reus100Currency := service.GetUs100Currency(us100Currency.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reus100Currency": reus100Currency}, c)
	}
}

// @Tags Us100Currency
// @Summary 分页获取Us100Currency列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.Us100CurrencySearch true "分页获取Us100Currency列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /us100Currency/getUs100CurrencyList [get]
func GetUs100CurrencyList(c *gin.Context) {
	var pageInfo request.Us100CurrencySearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetUs100CurrencyInfoList(pageInfo); err != nil {
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
