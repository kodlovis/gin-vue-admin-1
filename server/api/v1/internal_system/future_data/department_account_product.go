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

// @Tags DepartmentAccountProduct
// @Summary 创建DepartmentAccountProduct
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DepartmentAccountProduct true "创建DepartmentAccountProduct"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /departmentAccountProduct/createDepartmentAccountProduct [post]
func CreateDepartmentAccountProduct(c *gin.Context) {
	var departmentAccountProduct model.DepartmentAccountProduct
	_ = c.ShouldBindJSON(&departmentAccountProduct)
	if err := service.CreateDepartmentAccountProduct(departmentAccountProduct); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags DepartmentAccountProduct
// @Summary 删除DepartmentAccountProduct
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DepartmentAccountProduct true "删除DepartmentAccountProduct"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /departmentAccountProduct/deleteDepartmentAccountProduct [delete]
func DeleteDepartmentAccountProduct(c *gin.Context) {
	var departmentAccountProduct model.DepartmentAccountProduct
	_ = c.ShouldBindJSON(&departmentAccountProduct)
	if err := service.DeleteDepartmentAccountProduct(departmentAccountProduct); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags DepartmentAccountProduct
// @Summary 批量删除DepartmentAccountProduct
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除DepartmentAccountProduct"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /departmentAccountProduct/deleteDepartmentAccountProductByIds [delete]
func DeleteDepartmentAccountProductByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteDepartmentAccountProductByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags DepartmentAccountProduct
// @Summary 更新DepartmentAccountProduct
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DepartmentAccountProduct true "更新DepartmentAccountProduct"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /departmentAccountProduct/updateDepartmentAccountProduct [put]
func UpdateDepartmentAccountProduct(c *gin.Context) {
	var departmentAccountProduct model.DepartmentAccountProduct
	_ = c.ShouldBindJSON(&departmentAccountProduct)
	if err := service.UpdateDepartmentAccountProduct(departmentAccountProduct); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags DepartmentAccountProduct
// @Summary 用id查询DepartmentAccountProduct
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DepartmentAccountProduct true "用id查询DepartmentAccountProduct"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /departmentAccountProduct/findDepartmentAccountProduct [get]
func FindDepartmentAccountProduct(c *gin.Context) {
	var departmentAccountProduct model.DepartmentAccountProduct
	_ = c.ShouldBindQuery(&departmentAccountProduct)
	if err, redepartmentAccountProduct := service.GetDepartmentAccountProduct(departmentAccountProduct.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"redepartmentAccountProduct": redepartmentAccountProduct}, c)
	}
}

// @Tags DepartmentAccountProduct
// @Summary 分页获取DepartmentAccountProduct列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.DepartmentAccountProductSearch true "分页获取DepartmentAccountProduct列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /departmentAccountProduct/getDepartmentAccountProductList [get]
func GetDepartmentAccountProductList(c *gin.Context) {
	var pageInfo request.DepartmentAccountProductSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetDepartmentAccountProductInfoList(pageInfo); err != nil {
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
