import service from '@/utils/request'

// @Tags FutureDeliveryDetailByInput
// @Summary 创建FutureDeliveryDetailByInput
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.FutureDeliveryDetailByInput true "创建FutureDeliveryDetailByInput"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /futureDeliveryDetailByInput/createFutureDeliveryDetailByInput [post]
export const createFutureDeliveryDetailByInput = (data) => {
     return service({
         url: "/futureDeliveryDetailByInput/createFutureDeliveryDetailByInput",
         method: 'post',
         data
     })
 }


// @Tags FutureDeliveryDetailByInput
// @Summary 删除FutureDeliveryDetailByInput
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.FutureDeliveryDetailByInput true "删除FutureDeliveryDetailByInput"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /futureDeliveryDetailByInput/deleteFutureDeliveryDetailByInput [delete]
 export const deleteFutureDeliveryDetailByInput = (data) => {
     return service({
         url: "/futureDeliveryDetailByInput/deleteFutureDeliveryDetailByInput",
         method: 'delete',
         data
     })
 }

// @Tags FutureDeliveryDetailByInput
// @Summary 删除FutureDeliveryDetailByInput
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除FutureDeliveryDetailByInput"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /futureDeliveryDetailByInput/deleteFutureDeliveryDetailByInput [delete]
 export const deleteFutureDeliveryDetailByInputByIds = (data) => {
     return service({
         url: "/futureDeliveryDetailByInput/deleteFutureDeliveryDetailByInputByIds",
         method: 'delete',
         data
     })
 }

// @Tags FutureDeliveryDetailByInput
// @Summary 更新FutureDeliveryDetailByInput
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.FutureDeliveryDetailByInput true "更新FutureDeliveryDetailByInput"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /futureDeliveryDetailByInput/updateFutureDeliveryDetailByInput [put]
 export const updateFutureDeliveryDetailByInput = (data) => {
     return service({
         url: "/futureDeliveryDetailByInput/updateFutureDeliveryDetailByInput",
         method: 'put',
         data
     })
 }


// @Tags FutureDeliveryDetailByInput
// @Summary 用id查询FutureDeliveryDetailByInput
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.FutureDeliveryDetailByInput true "用id查询FutureDeliveryDetailByInput"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /futureDeliveryDetailByInput/findFutureDeliveryDetailByInput [get]
 export const findFutureDeliveryDetailByInput = (params) => {
     return service({
         url: "/futureDeliveryDetailByInput/findFutureDeliveryDetailByInput",
         method: 'get',
         params
     })
 }


// @Tags FutureDeliveryDetailByInput
// @Summary 分页获取FutureDeliveryDetailByInput列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取FutureDeliveryDetailByInput列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /futureDeliveryDetailByInput/getFutureDeliveryDetailByInputList [get]
 export const getFutureDeliveryDetailByInputList = (params) => {
     return service({
         url: "/futureDeliveryDetailByInput/getFutureDeliveryDetailByInputList",
         method: 'get',
         params
     })
 }