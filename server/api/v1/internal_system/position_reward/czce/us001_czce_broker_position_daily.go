package czce

import (
	"gin-vue-admin/global"
	model "gin-vue-admin/model/internal_system/position_reward/czce"
	request "gin-vue-admin/model/request/internal_system/position_reward/czce"
	"gin-vue-admin/model/response"
	service "gin-vue-admin/service/internal_system/position_reward/czce"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// @Tags Us001CzceBrokerPositionDaily
// @Summary 创建Us001CzceBrokerPositionDaily
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Us001CzceBrokerPositionDaily true "创建Us001CzceBrokerPositionDaily"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /us001CzceBrokerPositionDaily/createUs001CzceBrokerPositionDaily [post]
func CreateUs001CzceBrokerPositionDaily(c *gin.Context) {
	var us001CzceBrokerPositionDaily model.Us001CzceBrokerPositionDaily
	_ = c.ShouldBindJSON(&us001CzceBrokerPositionDaily)
	if err := service.CreateUs001CzceBrokerPositionDaily(us001CzceBrokerPositionDaily); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags Us001CzceBrokerPositionDaily
// @Summary 删除Us001CzceBrokerPositionDaily
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Us001CzceBrokerPositionDaily true "删除Us001CzceBrokerPositionDaily"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /us001CzceBrokerPositionDaily/deleteUs001CzceBrokerPositionDaily [delete]
func DeleteUs001CzceBrokerPositionDaily(c *gin.Context) {
	var us001CzceBrokerPositionDaily model.Us001CzceBrokerPositionDaily
	_ = c.ShouldBindJSON(&us001CzceBrokerPositionDaily)
	if err := service.DeleteUs001CzceBrokerPositionDaily(us001CzceBrokerPositionDaily); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Us001CzceBrokerPositionDaily
// @Summary 批量删除Us001CzceBrokerPositionDaily
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Us001CzceBrokerPositionDaily"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /us001CzceBrokerPositionDaily/deleteUs001CzceBrokerPositionDailyByIds [delete]
func DeleteUs001CzceBrokerPositionDailyByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteUs001CzceBrokerPositionDailyByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags Us001CzceBrokerPositionDaily
// @Summary 更新Us001CzceBrokerPositionDaily
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Us001CzceBrokerPositionDaily true "更新Us001CzceBrokerPositionDaily"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /us001CzceBrokerPositionDaily/updateUs001CzceBrokerPositionDaily [put]
func UpdateUs001CzceBrokerPositionDaily(c *gin.Context) {
	var us001CzceBrokerPositionDaily model.Us001CzceBrokerPositionDaily
	_ = c.ShouldBindJSON(&us001CzceBrokerPositionDaily)
	if err := service.UpdateUs001CzceBrokerPositionDaily(us001CzceBrokerPositionDaily); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags Us001CzceBrokerPositionDaily
// @Summary 用id查询Us001CzceBrokerPositionDaily
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Us001CzceBrokerPositionDaily true "用id查询Us001CzceBrokerPositionDaily"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /us001CzceBrokerPositionDaily/findUs001CzceBrokerPositionDaily [get]
func FindUs001CzceBrokerPositionDaily(c *gin.Context) {
	var us001CzceBrokerPositionDaily model.Us001CzceBrokerPositionDaily
	_ = c.ShouldBindQuery(&us001CzceBrokerPositionDaily)
	if err, reus001CzceBrokerPositionDaily := service.GetUs001CzceBrokerPositionDaily(us001CzceBrokerPositionDaily.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reus001CzceBrokerPositionDaily": reus001CzceBrokerPositionDaily}, c)
	}
}

// @Tags Us001CzceBrokerPositionDaily
// @Summary 分页获取Us001CzceBrokerPositionDaily列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.Us001CzceBrokerPositionDailySearch true "分页获取Us001CzceBrokerPositionDaily列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /us001CzceBrokerPositionDaily/getUs001CzceBrokerPositionDailyList [get]
func GetUs001CzceBrokerPositionDailyList(c *gin.Context) {
	var pageInfo request.Us001CzceBrokerPositionDailySearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetUs001CzceBrokerPositionDailyInfoList(pageInfo); err != nil {
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

func LoadUs001CzceBrokerPositionExcelData(c *gin.Context) {
	err := service.ParseUs001CzceBrokerPositionExcel2InfoList()
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
