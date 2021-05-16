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

// @Products Product
// @Summary 创建Product
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Product true "创建Product"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /Product/createProduct [post]
func CreateProduct(c *gin.Context) {
	var Product mp.Product
	_ = c.ShouldBindJSON(&Product)
	if err := sp.CreateProduct(Product); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Products Product
// @Summary 删除Product
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Product true "删除Product"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /Product/deleteProduct [delete]
func DeleteProduct(c *gin.Context) {
	var Product mp.Product
	_ = c.ShouldBindJSON(&Product)
	if err := sp.DeleteProduct(Product); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Products Product
// @Summary 批量删除Product
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Product"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /Product/deleteProductByIds [delete]
func DeleteProductByIds(c *gin.Context) {
	var IDS rp.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := sp.DeleteProductByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Products Product
// @Summary 更新Product
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Product true "更新Product"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /Product/updateProduct [put]
func UpdateProduct(c *gin.Context) {
	var Product mp.Product
	_ = c.ShouldBindJSON(&Product)
	if err := sp.UpdateProduct(&Product); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Products Product
// @Summary 用id查询Product
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Product true "用id查询Product"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /Product/findProduct [get]
func FindProduct(c *gin.Context) {
	var Product mp.Product
	_ = c.ShouldBindQuery(&Product)
	if err, reProduct := sp.GetProduct(Product.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reProduct": reProduct}, c)
	}
}

// @Products Product
// @Summary 分页获取Product列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.ProductSearch true "分页获取Product列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /Product/getProductList [get]
func GetProductList(c *gin.Context) {
	var pageInfo rp.ProductSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := sp.GetProductInfoList(pageInfo); err != nil {
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
