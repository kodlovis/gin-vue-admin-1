import service from '@/utils/request'

// @Tags TradeDetail
// @Summary 创建TradeDetail
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TradeDetail true "创建TradeDetail"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /tradeDetail/createTradeDetail [post]
export const createTradeDetail = (data) => {
     return service({
         url: "/tradeDetail/createTradeDetail",
         method: 'post',
         data
     })
 }


// @Tags TradeDetail
// @Summary 删除TradeDetail
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TradeDetail true "删除TradeDetail"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /tradeDetail/deleteTradeDetail [delete]
 export const deleteTradeDetail = (data) => {
     return service({
         url: "/tradeDetail/deleteTradeDetail",
         method: 'delete',
         data
     })
 }

// @Tags TradeDetail
// @Summary 删除TradeDetail
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除TradeDetail"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /tradeDetail/deleteTradeDetail [delete]
 export const deleteTradeDetailByIds = (data) => {
     return service({
         url: "/tradeDetail/deleteTradeDetailByIds",
         method: 'delete',
         data
     })
 }

// @Tags TradeDetail
// @Summary 更新TradeDetail
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TradeDetail true "更新TradeDetail"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /tradeDetail/updateTradeDetail [put]
 export const updateTradeDetail = (data) => {
     return service({
         url: "/tradeDetail/updateTradeDetail",
         method: 'put',
         data
     })
 }


// @Tags TradeDetail
// @Summary 用id查询TradeDetail
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TradeDetail true "用id查询TradeDetail"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /tradeDetail/findTradeDetail [get]
 export const findTradeDetail = (params) => {
     return service({
         url: "/tradeDetail/findTradeDetail",
         method: 'get',
         params
     })
 }


// @Tags TradeDetail
// @Summary 分页获取TradeDetail列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取TradeDetail列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /tradeDetail/getTradeDetailList [get]
 export const getTradeDetailList = (params) => {
     return service({
         url: "/tradeDetail/getTradeDetailList",
         method: 'get',
         params
     })
 }