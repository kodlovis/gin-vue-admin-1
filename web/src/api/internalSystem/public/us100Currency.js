import service from '@/utils/request'

// @Tags Us100Currency
// @Summary 创建Us100Currency
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Us100Currency true "创建Us100Currency"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /us100Currency/createUs100Currency [post]
export const createUs100Currency = (data) => {
     return service({
         url: "/us100Currency/createUs100Currency",
         method: 'post',
         data
     })
 }


// @Tags Us100Currency
// @Summary 删除Us100Currency
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Us100Currency true "删除Us100Currency"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /us100Currency/deleteUs100Currency [delete]
 export const deleteUs100Currency = (data) => {
     return service({
         url: "/us100Currency/deleteUs100Currency",
         method: 'delete',
         data
     })
 }

// @Tags Us100Currency
// @Summary 删除Us100Currency
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Us100Currency"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /us100Currency/deleteUs100Currency [delete]
 export const deleteUs100CurrencyByIds = (data) => {
     return service({
         url: "/us100Currency/deleteUs100CurrencyByIds",
         method: 'delete',
         data
     })
 }

// @Tags Us100Currency
// @Summary 更新Us100Currency
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Us100Currency true "更新Us100Currency"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /us100Currency/updateUs100Currency [put]
 export const updateUs100Currency = (data) => {
     return service({
         url: "/us100Currency/updateUs100Currency",
         method: 'put',
         data
     })
 }


// @Tags Us100Currency
// @Summary 用id查询Us100Currency
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Us100Currency true "用id查询Us100Currency"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /us100Currency/findUs100Currency [get]
 export const findUs100Currency = (params) => {
     return service({
         url: "/us100Currency/findUs100Currency",
         method: 'get',
         params
     })
 }


// @Tags Us100Currency
// @Summary 分页获取Us100Currency列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取Us100Currency列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /us100Currency/getUs100CurrencyList [get]
 export const getUs100CurrencyList = (params) => {
     return service({
         url: "/us100Currency/getUs100CurrencyList",
         method: 'get',
         params
     })
 }