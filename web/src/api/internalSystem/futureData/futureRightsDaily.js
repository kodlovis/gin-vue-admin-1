import service from '@/utils/request'

// @Tags FutureRightsDaily
// @Summary 创建FutureRightsDaily
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.FutureRightsDaily true "创建FutureRightsDaily"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /futureRightsDaily/createFutureRightsDaily [post]
export const createFutureRightsDaily = (data) => {
     return service({
         url: "/futureRightsDaily/createFutureRightsDaily",
         method: 'post',
         data
     })
 }


// @Tags FutureRightsDaily
// @Summary 删除FutureRightsDaily
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.FutureRightsDaily true "删除FutureRightsDaily"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /futureRightsDaily/deleteFutureRightsDaily [delete]
 export const deleteFutureRightsDaily = (data) => {
     return service({
         url: "/futureRightsDaily/deleteFutureRightsDaily",
         method: 'delete',
         data
     })
 }

// @Tags FutureRightsDaily
// @Summary 删除FutureRightsDaily
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除FutureRightsDaily"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /futureRightsDaily/deleteFutureRightsDaily [delete]
 export const deleteFutureRightsDailyByIds = (data) => {
     return service({
         url: "/futureRightsDaily/deleteFutureRightsDailyByIds",
         method: 'delete',
         data
     })
 }

// @Tags FutureRightsDaily
// @Summary 更新FutureRightsDaily
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.FutureRightsDaily true "更新FutureRightsDaily"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /futureRightsDaily/updateFutureRightsDaily [put]
 export const updateFutureRightsDaily = (data) => {
     return service({
         url: "/futureRightsDaily/updateFutureRightsDaily",
         method: 'put',
         data
     })
 }


// @Tags FutureRightsDaily
// @Summary 用id查询FutureRightsDaily
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.FutureRightsDaily true "用id查询FutureRightsDaily"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /futureRightsDaily/findFutureRightsDaily [get]
 export const findFutureRightsDaily = (params) => {
     return service({
         url: "/futureRightsDaily/findFutureRightsDaily",
         method: 'get',
         params
     })
 }


// @Tags FutureRightsDaily
// @Summary 分页获取FutureRightsDaily列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取FutureRightsDaily列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /futureRightsDaily/getFutureRightsDailyList [get]
 export const getFutureRightsDailyList = (params) => {
     return service({
         url: "/futureRightsDaily/getFutureRightsDailyList",
         method: 'get',
         params
     })
 }