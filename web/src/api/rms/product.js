import service from '@/utils/request'

// @Products Product
// @Summary 创建Product
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Product true "创建Product"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /Product/createProduct [post]
export const createProduct = (data) => {
     return service({
         url: "/Product/createProduct",
         method: 'post',
         data
     })
 }


// @Products Product
// @Summary 删除Product
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Product true "删除Product"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /Product/deleteProduct [delete]
 export const deleteProduct = (data) => {
     return service({
         url: "/Product/deleteProduct",
         method: 'delete',
         data
     })
 }

// @Products Product
// @Summary 删除Product
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Product"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /Product/deleteProduct [delete]
 export const deleteProductByIds = (data) => {
     return service({
         url: "/Product/deleteProductByIds",
         method: 'delete',
         data
     })
 }

// @Products Product
// @Summary 更新Product
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Product true "更新Product"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /Product/updateProduct [put]
 export const updateProduct = (data) => {
     return service({
         url: "/Product/updateProduct",
         method: 'put',
         data
     })
 }


// @Products Product
// @Summary 用id查询Product
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Product true "用id查询Product"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /Product/findProduct [get]
 export const findProduct = (params) => {
     return service({
         url: "/Product/findProduct",
         method: 'get',
         params
     })
 }


// @Products Product
// @Summary 分页获取Product列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取Product列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /Product/getProductList [get]
 export const getProductList = (params) => {
     return service({
         url: "/Product/getProductList",
         method: 'get',
         params
     })
 }