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

// @Userss Users
// @Summary 创建Users
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Users true "创建Users"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /Users/createUsers [post]
func CreateUsers(c *gin.Context) {
	var Users mp.Users
	_ = c.ShouldBindJSON(&Users)
	if err := sp.CreateUsers(Users); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Userss Users
// @Summary 删除Users
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Users true "删除Users"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /Users/deleteUsers [delete]
func DeleteUsers(c *gin.Context) {
	var Users mp.Users
	_ = c.ShouldBindJSON(&Users)
	if err := sp.DeleteUsers(Users); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Userss Users
// @Summary 批量删除Users
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Users"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /Users/deleteUsersByIds [delete]
func DeleteUsersByIds(c *gin.Context) {
	var IDS rp.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := sp.DeleteUsersByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Userss Users
// @Summary 更新Users
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Users true "更新Users"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /Users/updateUsers [put]
func UpdateUsers(c *gin.Context) {
	var Users mp.Users
	_ = c.ShouldBindJSON(&Users)
	if err := sp.UpdateUsers(&Users); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Userss Users
// @Summary 用id查询Users
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Users true "用id查询Users"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /Users/findUsers [get]
func FindUsers(c *gin.Context) {
	var Users mp.Users
	_ = c.ShouldBindQuery(&Users)
	if err, reUsers := sp.GetUsers(Users.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reUsers": reUsers}, c)
	}
}

// @Userss Users
// @Summary 分页获取Users列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.UsersSearch true "分页获取Users列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /Users/getUsersList [get]
func GetUsersList(c *gin.Context) {
	var pageInfo rp.UsersSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := sp.GetUsersInfoList(pageInfo); err != nil {
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


func GetUIList(c *gin.Context) {
	var pageInfo rp.UsersSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total,code := sp.GetUIList(pageInfo.ID,pageInfo.Password,pageInfo.Mac); err != nil {
	    global.GVA_LOG.Error("获取失败", zap.Any("err", err))
        response.FailWithMessage("获取失败", c)
    } else {
        response.OkWithDetailed(response.PageResult{
            List:     list,
            Total:    total,
            Page:     pageInfo.Page,
            PageSize: pageInfo.PageSize,
			Code: code,
        }, "获取成功", c)
    }
}

func GetUIListByServer(c *gin.Context) {
	var pageInfo rp.UsersSearch
	if err, list, total := sp.GetUIListByServer(); err != nil {
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

func GetLastUser(c *gin.Context) {
	var User mp.Users
	_ = c.ShouldBindQuery(&User)
	if err, reUser := sp.GetLastUser(); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reUser": reUser}, c)
	}
}
func CreateUserMac(c *gin.Context) {
	var res rp.UserMacList
	_ = c.ShouldBindJSON(&res)
	if err := sp.CreateUserMac(res.UserMacList); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}
func CreateUserProduct(c *gin.Context) {
	var res rp.UserProductList
	_ = c.ShouldBindJSON(&res)
	if err := sp.CreateUserProduct(res.UserProductList); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}
func RemoveUserMacProduct(c *gin.Context) {
	var User mp.Users
	_ = c.ShouldBindJSON(&User)
	if err := sp.RemoveUserMacProduct(User); err != nil {
		global.GVA_LOG.Error("清除失败!", zap.Any("err", err))
		response.FailWithMessage("清除失败", c)
	} else {
		response.OkWithMessage("清除成功", c)
	}
}