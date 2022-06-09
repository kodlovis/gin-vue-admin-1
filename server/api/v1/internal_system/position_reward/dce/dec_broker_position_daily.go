package dce

import (
	"gin-vue-admin/global"
	model "gin-vue-admin/model/internal_system/position_reward/dce"
	request "gin-vue-admin/model/request/internal_system/position_reward/dce"
	"gin-vue-admin/model/response"
	service "gin-vue-admin/service/internal_system/position_reward/dce"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// @Tags Us001DceBrokerPositionDaily
// @Summary 创建Us001DceBrokerPositionDaily
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Us001DceBrokerPositionDaily true "创建Us001DceBrokerPositionDaily"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /us001DceBrokerPositionDaily/createUs001DceBrokerPositionDaily [post]
func CreateUs001DceBrokerPositionDaily(c *gin.Context) {
	var us001DceBrokerPositionDaily model.Us001DceBrokerPositionDaily
	_ = c.ShouldBindJSON(&us001DceBrokerPositionDaily)
	if err := service.CreateUs001DceBrokerPositionDaily(us001DceBrokerPositionDaily); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags Us001DceBrokerPositionDaily
// @Summary 删除Us001DceBrokerPositionDaily
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Us001DceBrokerPositionDaily true "删除Us001DceBrokerPositionDaily"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /us001DceBrokerPositionDaily/deleteUs001DceBrokerPositionDaily [delete]
func DeleteUs001DceBrokerPositionDaily(c *gin.Context) {
	var us001DceBrokerPositionDaily model.Us001DceBrokerPositionDaily
	_ = c.ShouldBindJSON(&us001DceBrokerPositionDaily)
	if err := service.DeleteUs001DceBrokerPositionDaily(us001DceBrokerPositionDaily); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Us001DceBrokerPositionDaily
// @Summary 批量删除Us001DceBrokerPositionDaily
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Us001DceBrokerPositionDaily"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /us001DceBrokerPositionDaily/deleteUs001DceBrokerPositionDailyByIds [delete]
func DeleteUs001DceBrokerPositionDailyByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteUs001DceBrokerPositionDailyByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags Us001DceBrokerPositionDaily
// @Summary 更新Us001DceBrokerPositionDaily
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Us001DceBrokerPositionDaily true "更新Us001DceBrokerPositionDaily"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /us001DceBrokerPositionDaily/updateUs001DceBrokerPositionDaily [put]
func UpdateUs001DceBrokerPositionDaily(c *gin.Context) {
	var us001DceBrokerPositionDaily model.Us001DceBrokerPositionDaily
	_ = c.ShouldBindJSON(&us001DceBrokerPositionDaily)
	if err := service.UpdateUs001DceBrokerPositionDaily(us001DceBrokerPositionDaily); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags Us001DceBrokerPositionDaily
// @Summary 用id查询Us001DceBrokerPositionDaily
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Us001DceBrokerPositionDaily true "用id查询Us001DceBrokerPositionDaily"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /us001DceBrokerPositionDaily/findUs001DceBrokerPositionDaily [get]
func FindUs001DceBrokerPositionDaily(c *gin.Context) {
	var us001DceBrokerPositionDaily model.Us001DceBrokerPositionDaily
	_ = c.ShouldBindQuery(&us001DceBrokerPositionDaily)
	if err, reus001DceBrokerPositionDaily := service.GetUs001DceBrokerPositionDaily(us001DceBrokerPositionDaily.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reus001DceBrokerPositionDaily": reus001DceBrokerPositionDaily}, c)
	}
}

// @Tags Us001DceBrokerPositionDaily
// @Summary 分页获取Us001DceBrokerPositionDaily列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.Us001DceBrokerPositionDailySearch true "分页获取Us001DceBrokerPositionDaily列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /us001DceBrokerPositionDaily/getUs001DceBrokerPositionDailyList [get]
func GetUs001DceBrokerPositionDailyList(c *gin.Context) {
	var pageInfo request.Us001DceBrokerPositionDailySearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetUs001DceBrokerPositionDailyInfoList(pageInfo); err != nil {
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

func LoadUs001DceBrokerPositionExcelData(c *gin.Context) {
	err := service.ParseUs001DceBrokerPositionExcel2InfoList()
	if err != nil {
		global.GVA_LOG.Error("加载数据失败", zap.Any("err", err))
		response.FailWithMessage("加载数据失败", c)
		return
	} else {
		response.OkWithMessage("创建成功", c)
	}
	// response.OkWithDetailed(response.PageResult{
	// 	List:     menus,
	// 	Total:    int64(len(menus)),
	// 	Page:     1,
	// 	PageSize: 999,
	// }, "加载数据成功", c)
}
