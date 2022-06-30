import service from '@/utils/request'

// @Tags Us002ForeignExchangeProfit
// @Summary 创建Us002ForeignExchangeProfit
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Us002ForeignExchangeProfit true "创建Us002ForeignExchangeProfit"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /us002ForeignExchangeProfit/createUs002ForeignExchangeProfit [post]
export const createUs002ForeignExchangeProfit = (data) => {
     return service({
         url: "/us002ForeignExchangeProfit/createUs002ForeignExchangeProfit",
         method: 'post',
         data
     })
 }


// @Tags Us002ForeignExchangeProfit
// @Summary 删除Us002ForeignExchangeProfit
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Us002ForeignExchangeProfit true "删除Us002ForeignExchangeProfit"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /us002ForeignExchangeProfit/deleteUs002ForeignExchangeProfit [delete]
 export const deleteUs002ForeignExchangeProfit = (data) => {
     return service({
         url: "/us002ForeignExchangeProfit/deleteUs002ForeignExchangeProfit",
         method: 'delete',
         data
     })
 }

// @Tags Us002ForeignExchangeProfit
// @Summary 删除Us002ForeignExchangeProfit
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Us002ForeignExchangeProfit"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /us002ForeignExchangeProfit/deleteUs002ForeignExchangeProfit [delete]
 export const deleteUs002ForeignExchangeProfitByIds = (data) => {
     return service({
         url: "/us002ForeignExchangeProfit/deleteUs002ForeignExchangeProfitByIds",
         method: 'delete',
         data
     })
 }

// @Tags Us002ForeignExchangeProfit
// @Summary 更新Us002ForeignExchangeProfit
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Us002ForeignExchangeProfit true "更新Us002ForeignExchangeProfit"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /us002ForeignExchangeProfit/updateUs002ForeignExchangeProfit [put]
 export const updateUs002ForeignExchangeProfit = (data) => {
     return service({
         url: "/us002ForeignExchangeProfit/updateUs002ForeignExchangeProfit",
         method: 'put',
         data
     })
 }


// @Tags Us002ForeignExchangeProfit
// @Summary 用id查询Us002ForeignExchangeProfit
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Us002ForeignExchangeProfit true "用id查询Us002ForeignExchangeProfit"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /us002ForeignExchangeProfit/findUs002ForeignExchangeProfit [get]
 export const findUs002ForeignExchangeProfit = (params) => {
     return service({
         url: "/us002ForeignExchangeProfit/findUs002ForeignExchangeProfit",
         method: 'get',
         params
     })
 }


// @Tags Us002ForeignExchangeProfit
// @Summary 分页获取Us002ForeignExchangeProfit列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取Us002ForeignExchangeProfit列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /us002ForeignExchangeProfit/getUs002ForeignExchangeProfitList [get]
 export const getUs002ForeignExchangeProfitList = (params) => {
     return service({
         url: "/us002ForeignExchangeProfit/getUs002ForeignExchangeProfitList",
         method: 'get',
         params
     })
 }