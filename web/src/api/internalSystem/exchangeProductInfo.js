import service from '@/utils/request'

// @Tags ExchangeProductInfo
// @Summary 创建ExchangeProductInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ExchangeProductInfo true "创建ExchangeProductInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /exchangeProductInfo/createExchangeProductInfo [post]
export const createExchangeProductInfo = (data) => {
     return service({
         url: "/exchangeProductInfo/createExchangeProductInfo",
         method: 'post',
         data
     })
 }


// @Tags ExchangeProductInfo
// @Summary 删除ExchangeProductInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ExchangeProductInfo true "删除ExchangeProductInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /exchangeProductInfo/deleteExchangeProductInfo [delete]
 export const deleteExchangeProductInfo = (data) => {
     return service({
         url: "/exchangeProductInfo/deleteExchangeProductInfo",
         method: 'delete',
         data
     })
 }

// @Tags ExchangeProductInfo
// @Summary 删除ExchangeProductInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ExchangeProductInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /exchangeProductInfo/deleteExchangeProductInfo [delete]
 export const deleteExchangeProductInfoByIds = (data) => {
     return service({
         url: "/exchangeProductInfo/deleteExchangeProductInfoByIds",
         method: 'delete',
         data
     })
 }

// @Tags ExchangeProductInfo
// @Summary 更新ExchangeProductInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ExchangeProductInfo true "更新ExchangeProductInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /exchangeProductInfo/updateExchangeProductInfo [put]
 export const updateExchangeProductInfo = (data) => {
     return service({
         url: "/exchangeProductInfo/updateExchangeProductInfo",
         method: 'put',
         data
     })
 }


// @Tags ExchangeProductInfo
// @Summary 用id查询ExchangeProductInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ExchangeProductInfo true "用id查询ExchangeProductInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /exchangeProductInfo/findExchangeProductInfo [get]
 export const findExchangeProductInfo = (params) => {
     return service({
         url: "/exchangeProductInfo/findExchangeProductInfo",
         method: 'get',
         params
     })
 }


// @Tags ExchangeProductInfo
// @Summary 分页获取ExchangeProductInfo列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取ExchangeProductInfo列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /exchangeProductInfo/getExchangeProductInfoList [get]
 export const getExchangeProductInfoList = (params) => {
     return service({
         url: "/exchangeProductInfo/getExchangeProductInfoList",
         method: 'get',
         params
     })
 }