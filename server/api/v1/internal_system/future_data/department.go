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

// @Tags Department
// @Summary 创建Department
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Department true "创建Department"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /department/createDepartment [post]
func CreateDepartment(c *gin.Context) {
	var department model.Department
	_ = c.ShouldBindJSON(&department)
	if err := service.CreateDepartment(department); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags Department
// @Summary 删除Department
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Department true "删除Department"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /department/deleteDepartment [delete]
func DeleteDepartment(c *gin.Context) {
	var department model.Department
	_ = c.ShouldBindJSON(&department)
	if err := service.DeleteDepartment(department); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Department
// @Summary 批量删除Department
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Department"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /department/deleteDepartmentByIds [delete]
func DeleteDepartmentByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteDepartmentByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags Department
// @Summary 更新Department
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Department true "更新Department"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /department/updateDepartment [put]
func UpdateDepartment(c *gin.Context) {
	var department model.Department
	_ = c.ShouldBindJSON(&department)
	if err := service.UpdateDepartment(department); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags Department
// @Summary 用id查询Department
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Department true "用id查询Department"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /department/findDepartment [get]
func FindDepartment(c *gin.Context) {
	var department model.Department
	_ = c.ShouldBindQuery(&department)
	if err, redepartment := service.GetDepartment(department.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"redepartment": redepartment}, c)
	}
}

// @Tags Department
// @Summary 分页获取Department列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.DepartmentSearch true "分页获取Department列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /department/getDepartmentList [get]
func GetDepartmentList(c *gin.Context) {
	var pageInfo request.DepartmentSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetDepartmentInfoList(pageInfo); err != nil {
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
