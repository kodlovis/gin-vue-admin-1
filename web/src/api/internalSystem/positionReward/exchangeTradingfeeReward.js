import service from '@/utils/request'

// @Tags ExchangeTradingfeeReward
// @Summary 创建ExchangeTradingfeeReward
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ExchangeTradingfeeReward true "创建ExchangeTradingfeeReward"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /exchangeTradingfeeReward/createExchangeTradingfeeReward [post]
export const createExchangeTradingfeeReward = (data) => {
     return service({
         url: "/exchangeTradingfeeReward/createExchangeTradingfeeReward",
         method: 'post',
         data
     })
 }


// @Tags ExchangeTradingfeeReward
// @Summary 删除ExchangeTradingfeeReward
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ExchangeTradingfeeReward true "删除ExchangeTradingfeeReward"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /exchangeTradingfeeReward/deleteExchangeTradingfeeReward [delete]
 export const deleteExchangeTradingfeeReward = (data) => {
     return service({
         url: "/exchangeTradingfeeReward/deleteExchangeTradingfeeReward",
         method: 'delete',
         data
     })
 }

// @Tags ExchangeTradingfeeReward
// @Summary 删除ExchangeTradingfeeReward
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ExchangeTradingfeeReward"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /exchangeTradingfeeReward/deleteExchangeTradingfeeReward [delete]
 export const deleteExchangeTradingfeeRewardByIds = (data) => {
     return service({
         url: "/exchangeTradingfeeReward/deleteExchangeTradingfeeRewardByIds",
         method: 'delete',
         data
     })
 }

// @Tags ExchangeTradingfeeReward
// @Summary 更新ExchangeTradingfeeReward
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ExchangeTradingfeeReward true "更新ExchangeTradingfeeReward"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /exchangeTradingfeeReward/updateExchangeTradingfeeReward [put]
 export const updateExchangeTradingfeeReward = (data) => {
     return service({
         url: "/exchangeTradingfeeReward/updateExchangeTradingfeeReward",
         method: 'put',
         data
     })
 }


// @Tags ExchangeTradingfeeReward
// @Summary 用id查询ExchangeTradingfeeReward
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ExchangeTradingfeeReward true "用id查询ExchangeTradingfeeReward"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /exchangeTradingfeeReward/findExchangeTradingfeeReward [get]
 export const findExchangeTradingfeeReward = (params) => {
     return service({
         url: "/exchangeTradingfeeReward/findExchangeTradingfeeReward",
         method: 'get',
         params
     })
 }


// @Tags ExchangeTradingfeeReward
// @Summary 分页获取ExchangeTradingfeeReward列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取ExchangeTradingfeeReward列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /exchangeTradingfeeReward/getExchangeTradingfeeRewardList [get]
 export const getExchangeTradingfeeRewardList = (params) => {
     return service({
         url: "/exchangeTradingfeeReward/getExchangeTradingfeeRewardList",
         method: 'get',
         params
     })
 }

 export const loadExchangeTradingfeeRewardExcelData = () => {
    return service({
        url: "/exchangeTradingfeeReward/loadExchangeTradingfeeRewardExcel",
        method: 'get'
    })
}
