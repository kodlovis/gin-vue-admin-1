import service from '@/utils/request'

// @Tags BrokerPositionDaily
// @Summary 创建BrokerPositionDaily
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BrokerPositionDaily true "创建BrokerPositionDaily"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /BrokerPositionDaily/createBrokerPositionDaily [post]
export const createBrokerPositionDaily = (data) => {
     return service({
         url: "/brokerPositionDaily/createBrokerPositionDaily",
         method: 'post',
         data
     })
 }


// @Tags BrokerPositionDaily
// @Summary 删除BrokerPositionDaily
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BrokerPositionDaily true "删除BrokerPositionDaily"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /BrokerPositionDaily/deleteBrokerPositionDaily [delete]
 export const deleteBrokerPositionDaily = (data) => {
     return service({
         url: "/brokerPositionDaily/deleteBrokerPositionDaily",
         method: 'delete',
         data
     })
 }

// @Tags BrokerPositionDaily
// @Summary 删除BrokerPositionDaily
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除BrokerPositionDaily"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /BrokerPositionDaily/deleteBrokerPositionDaily [delete]
 export const deleteBrokerPositionDailyByIds = (data) => {
     return service({
         url: "/brokerPositionDaily/deleteBrokerPositionDailyByIds",
         method: 'delete',
         data
     })
 }

// @Tags BrokerPositionDaily
// @Summary 更新BrokerPositionDaily
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BrokerPositionDaily true "更新BrokerPositionDaily"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /BrokerPositionDaily/updateBrokerPositionDaily [put]
 export const updateBrokerPositionDaily = (data) => {
     return service({
         url: "/brokerPositionDaily/updateBrokerPositionDaily",
         method: 'put',
         data
     })
 }


// @Tags BrokerPositionDaily
// @Summary 用id查询BrokerPositionDaily
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BrokerPositionDaily true "用id查询BrokerPositionDaily"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /BrokerPositionDaily/findBrokerPositionDaily [get]
 export const findBrokerPositionDaily = (params) => {
     return service({
         url: "/brokerPositionDaily/findBrokerPositionDaily",
         method: 'get',
         params
     })
 }


// @Tags BrokerPositionDaily
// @Summary 分页获取BrokerPositionDaily列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取BrokerPositionDaily列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /BrokerPositionDaily/getBrokerPositionDailyList [get]
 export const getBrokerPositionDailyList = (params) => {
     return service({
         url: "/brokerPositionDaily/getBrokerPositionDailyList",
         method: 'get',
         params
     })
 }

 
export const loadBrokerPositionExcelData = () => {
    return service({
        url: "/brokerPositionDaily/loadBrokerPositionExcel",
        method: 'get'
    })
}