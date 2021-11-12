import service from '@/utils/request'

// @Tags FutureDeparmentRights
// @Summary 创建FutureDeparmentRights
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.FutureDeparmentRights true "创建FutureDeparmentRights"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /futureDeparmentRights/createFutureDeparmentRights [post]
export const createFutureDeparmentRights = (data) => {
     return service({
         url: "/futureDeparmentRights/createFutureDeparmentRights",
         method: 'post',
         data
     })
 }


// @Tags FutureDeparmentRights
// @Summary 删除FutureDeparmentRights
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.FutureDeparmentRights true "删除FutureDeparmentRights"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /futureDeparmentRights/deleteFutureDeparmentRights [delete]
 export const deleteFutureDeparmentRights = (data) => {
     return service({
         url: "/futureDeparmentRights/deleteFutureDeparmentRights",
         method: 'delete',
         data
     })
 }

// @Tags FutureDeparmentRights
// @Summary 删除FutureDeparmentRights
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除FutureDeparmentRights"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /futureDeparmentRights/deleteFutureDeparmentRights [delete]
 export const deleteFutureDeparmentRightsByIds = (data) => {
     return service({
         url: "/futureDeparmentRights/deleteFutureDeparmentRightsByIds",
         method: 'delete',
         data
     })
 }

// @Tags FutureDeparmentRights
// @Summary 更新FutureDeparmentRights
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.FutureDeparmentRights true "更新FutureDeparmentRights"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /futureDeparmentRights/updateFutureDeparmentRights [put]
 export const updateFutureDeparmentRights = (data) => {
     return service({
         url: "/futureDeparmentRights/updateFutureDeparmentRights",
         method: 'put',
         data
     })
 }


// @Tags FutureDeparmentRights
// @Summary 用id查询FutureDeparmentRights
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.FutureDeparmentRights true "用id查询FutureDeparmentRights"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /futureDeparmentRights/findFutureDeparmentRights [get]
 export const findFutureDeparmentRights = (params) => {
     return service({
         url: "/futureDeparmentRights/findFutureDeparmentRights",
         method: 'get',
         params
     })
 }


// @Tags FutureDeparmentRights
// @Summary 分页获取FutureDeparmentRights列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取FutureDeparmentRights列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /futureDeparmentRights/getFutureDeparmentRightsList [get]
 export const getFutureDeparmentRightsList = (params) => {
     return service({
         url: "/futureDeparmentRights/getFutureDeparmentRightsList",
         method: 'get',
         params
     })
 }