import service from '@/utils/request'

// @Tags Us004FutureStockLme
// @Summary 创建Us004FutureStockLme
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Us004FutureStockLme true "创建Us004FutureStockLme"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /us004FutureStockLme/createUs004FutureStockLme [post]
export const createUs004FutureStockLme = (data) => {
     return service({
         url: "/us004FutureStockLme/createUs004FutureStockLme",
         method: 'post',
         data
     })
 }


// @Tags Us004FutureStockLme
// @Summary 删除Us004FutureStockLme
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Us004FutureStockLme true "删除Us004FutureStockLme"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /us004FutureStockLme/deleteUs004FutureStockLme [delete]
 export const deleteUs004FutureStockLme = (data) => {
     return service({
         url: "/us004FutureStockLme/deleteUs004FutureStockLme",
         method: 'delete',
         data
     })
 }

// @Tags Us004FutureStockLme
// @Summary 删除Us004FutureStockLme
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Us004FutureStockLme"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /us004FutureStockLme/deleteUs004FutureStockLme [delete]
 export const deleteUs004FutureStockLmeByIds = (data) => {
     return service({
         url: "/us004FutureStockLme/deleteUs004FutureStockLmeByIds",
         method: 'delete',
         data
     })
 }

// @Tags Us004FutureStockLme
// @Summary 更新Us004FutureStockLme
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Us004FutureStockLme true "更新Us004FutureStockLme"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /us004FutureStockLme/updateUs004FutureStockLme [put]
 export const updateUs004FutureStockLme = (data) => {
     return service({
         url: "/us004FutureStockLme/updateUs004FutureStockLme",
         method: 'put',
         data
     })
 }


// @Tags Us004FutureStockLme
// @Summary 用id查询Us004FutureStockLme
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Us004FutureStockLme true "用id查询Us004FutureStockLme"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /us004FutureStockLme/findUs004FutureStockLme [get]
 export const findUs004FutureStockLme = (params) => {
     return service({
         url: "/us004FutureStockLme/findUs004FutureStockLme",
         method: 'get',
         params
     })
 }


// @Tags Us004FutureStockLme
// @Summary 分页获取Us004FutureStockLme列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取Us004FutureStockLme列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /us004FutureStockLme/getUs004FutureStockLmeList [get]
 export const getUs004FutureStockLmeList = (params) => {
     return service({
         url: "/us004FutureStockLme/getUs004FutureStockLmeList",
         method: 'get',
         params
     })
 }