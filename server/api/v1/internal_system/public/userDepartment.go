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

// @Tags UserDepartment
// @Summary 创建UserDepartment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.UserDepartment true "创建UserDepartment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /userDepartment/createUserDepartment [post]
func CreateUserDepartment(c *gin.Context) {
	var userDepartment model.UserDepartment
	_ = c.ShouldBindJSON(&userDepartment)
	if err := service.CreateUserDepartment(userDepartment); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags UserDepartment
// @Summary 删除UserDepartment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.UserDepartment true "删除UserDepartment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /userDepartment/deleteUserDepartment [delete]
func DeleteUserDepartment(c *gin.Context) {
	var userDepartment model.UserDepartment
	_ = c.ShouldBindJSON(&userDepartment)
	if err := service.DeleteUserDepartment(userDepartment); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags UserDepartment
// @Summary 批量删除UserDepartment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除UserDepartment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /userDepartment/deleteUserDepartmentByIds [delete]
func DeleteUserDepartmentByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteUserDepartmentByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags UserDepartment
// @Summary 更新UserDepartment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.UserDepartment true "更新UserDepartment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /userDepartment/updateUserDepartment [put]
func UpdateUserDepartment(c *gin.Context) {
	var userDepartment model.UserDepartment
	_ = c.ShouldBindJSON(&userDepartment)
	if err := service.UpdateUserDepartment(userDepartment); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags UserDepartment
// @Summary 用id查询UserDepartment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.UserDepartment true "用id查询UserDepartment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /userDepartment/findUserDepartment [get]
func FindUserDepartment(c *gin.Context) {
	var userDepartment model.UserDepartment
	_ = c.ShouldBindQuery(&userDepartment)
	if err, reuserDepartment := service.GetUserDepartment(userDepartment.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reuserDepartment": reuserDepartment}, c)
	}
}

// @Tags UserDepartment
// @Summary 分页获取UserDepartment列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.UserDepartmentSearch true "分页获取UserDepartment列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /userDepartment/getUserDepartmentList [get]
func GetUserDepartmentList(c *gin.Context) {
	var pageInfo request.UserDepartmentSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetUserDepartmentInfoList(pageInfo); err != nil {
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

func GetUserDepartmentListByID(c *gin.Context) {
	var pageInfo request.UserDepartmentSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetUserDepartmentListByID(pageInfo); err != nil {
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
