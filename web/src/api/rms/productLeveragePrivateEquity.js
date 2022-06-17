import service from '@/utils/request'

// @Tags ProductLeveragePrivateEquity
// @Summary 创建ProductLeveragePrivateEquity
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ProductLeveragePrivateEquity true "创建ProductLeveragePrivateEquity"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /productLeveragePrivateEquity/createProductLeveragePrivateEquity [post]
export const createProductLeveragePrivateEquity = (data) => {
     return service({
         url: "/productLeveragePrivateEquity/createProductLeveragePrivateEquity",
         method: 'post',
         data
     })
 }


// @Tags ProductLeveragePrivateEquity
// @Summary 删除ProductLeveragePrivateEquity
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ProductLeveragePrivateEquity true "删除ProductLeveragePrivateEquity"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /productLeveragePrivateEquity/deleteProductLeveragePrivateEquity [delete]
 export const deleteProductLeveragePrivateEquity = (data) => {
     return service({
         url: "/productLeveragePrivateEquity/deleteProductLeveragePrivateEquity",
         method: 'delete',
         data
     })
 }

// @Tags ProductLeveragePrivateEquity
// @Summary 删除ProductLeveragePrivateEquity
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ProductLeveragePrivateEquity"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /productLeveragePrivateEquity/deleteProductLeveragePrivateEquity [delete]
 export const deleteProductLeveragePrivateEquityByIds = (data) => {
     return service({
         url: "/productLeveragePrivateEquity/deleteProductLeveragePrivateEquityByIds",
         method: 'delete',
         data
     })
 }

// @Tags ProductLeveragePrivateEquity
// @Summary 更新ProductLeveragePrivateEquity
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ProductLeveragePrivateEquity true "更新ProductLeveragePrivateEquity"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /productLeveragePrivateEquity/updateProductLeveragePrivateEquity [put]
 export const updateProductLeveragePrivateEquity = (data) => {
     return service({
         url: "/productLeveragePrivateEquity/updateProductLeveragePrivateEquity",
         method: 'put',
         data
     })
 }


// @Tags ProductLeveragePrivateEquity
// @Summary 用id查询ProductLeveragePrivateEquity
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ProductLeveragePrivateEquity true "用id查询ProductLeveragePrivateEquity"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /productLeveragePrivateEquity/findProductLeveragePrivateEquity [get]
 export const findProductLeveragePrivateEquity = (params) => {
     return service({
         url: "/productLeveragePrivateEquity/findProductLeveragePrivateEquity",
         method: 'get',
         params
     })
 }


// @Tags ProductLeveragePrivateEquity
// @Summary 分页获取ProductLeveragePrivateEquity列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取ProductLeveragePrivateEquity列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /productLeveragePrivateEquity/getProductLeveragePrivateEquityList [get]
 export const getProductLeveragePrivateEquityList = (params) => {
     return service({
         url: "/productLeveragePrivateEquity/getProductLeveragePrivateEquityList",
         method: 'get',
         params
     })
 }