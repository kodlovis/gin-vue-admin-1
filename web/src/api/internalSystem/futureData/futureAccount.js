import service from '@/utils/request'

// @Tags FutureAccount
// @Summary 创建FutureAccount
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.FutureAccount true "创建FutureAccount"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /futureAccount/createFutureAccount [post]
export const createFutureAccount = (data) => {
     return service({
         url: "/futureAccount/createFutureAccount",
         method: 'post',
         data
     })
 }


// @Tags FutureAccount
// @Summary 删除FutureAccount
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.FutureAccount true "删除FutureAccount"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /futureAccount/deleteFutureAccount [delete]
 export const deleteFutureAccount = (data) => {
     return service({
         url: "/futureAccount/deleteFutureAccount",
         method: 'delete',
         data
     })
 }

// @Tags FutureAccount
// @Summary 删除FutureAccount
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除FutureAccount"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /futureAccount/deleteFutureAccount [delete]
 export const deleteFutureAccountByIds = (data) => {
     return service({
         url: "/futureAccount/deleteFutureAccountByIds",
         method: 'delete',
         data
     })
 }

// @Tags FutureAccount
// @Summary 更新FutureAccount
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.FutureAccount true "更新FutureAccount"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /futureAccount/updateFutureAccount [put]
 export const updateFutureAccount = (data) => {
     return service({
         url: "/futureAccount/updateFutureAccount",
         method: 'put',
         data
     })
 }


// @Tags FutureAccount
// @Summary 用id查询FutureAccount
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.FutureAccount true "用id查询FutureAccount"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /futureAccount/findFutureAccount [get]
 export const findFutureAccount = (params) => {
     return service({
         url: "/futureAccount/findFutureAccount",
         method: 'get',
         params
     })
 }


// @Tags FutureAccount
// @Summary 分页获取FutureAccount列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取FutureAccount列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /futureAccount/getFutureAccountList [get]
 export const getFutureAccountList = (params) => {
     return service({
         url: "/futureAccount/getFutureAccountList",
         method: 'get',
         params
     })
 }