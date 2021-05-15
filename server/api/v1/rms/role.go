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

// @Roles Role
// @Summary 创建Role
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Role true "创建Role"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /Role/createRole [post]
func CreateRole(c *gin.Context) {
	var Role mp.Role
	_ = c.ShouldBindJSON(&Role)
	if err := sp.CreateRole(Role); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Roles Role
// @Summary 删除Role
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Role true "删除Role"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /Role/deleteRole [delete]
func DeleteRole(c *gin.Context) {
	var Role mp.Role
	_ = c.ShouldBindJSON(&Role)
	if err := sp.DeleteRole(Role); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Roles Role
// @Summary 批量删除Role
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Role"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /Role/deleteRoleByIds [delete]
func DeleteRoleByIds(c *gin.Context) {
	var IDS rp.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := sp.DeleteRoleByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Roles Role
// @Summary 更新Role
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Role true "更新Role"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /Role/updateRole [put]
func UpdateRole(c *gin.Context) {
	var Role mp.Role
	_ = c.ShouldBindJSON(&Role)
	if err := sp.UpdateRole(&Role); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Roles Role
// @Summary 用id查询Role
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Role true "用id查询Role"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /Role/findRole [get]
func FindRole(c *gin.Context) {
	var Role mp.Role
	_ = c.ShouldBindQuery(&Role)
	if err, reRole := sp.GetRole(Role.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reRole": reRole}, c)
	}
}

// @Roles Role
// @Summary 分页获取Role列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.RoleSearch true "分页获取Role列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /Role/getRoleList [get]
func GetRoleList(c *gin.Context) {
	var pageInfo rp.RoleSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := sp.GetRoleInfoList(pageInfo); err != nil {
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
