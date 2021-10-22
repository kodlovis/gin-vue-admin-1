import service from '@/utils/request'

// @Tags AccountPositionDaily
// @Summary 创建AccountPositionDaily
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AccountPositionDaily true "创建AccountPositionDaily"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /accountPositionDaily/createAccountPositionDaily [post]
export const createAccountPositionDaily = (data) => {
     return service({
         url: "/accountPositionDaily/createAccountPositionDaily",
         method: 'post',
         data
     })
 }


// @Tags AccountPositionDaily
// @Summary 删除AccountPositionDaily
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AccountPositionDaily true "删除AccountPositionDaily"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /accountPositionDaily/deleteAccountPositionDaily [delete]
 export const deleteAccountPositionDaily = (data) => {
     return service({
         url: "/accountPositionDaily/deleteAccountPositionDaily",
         method: 'delete',
         data
     })
 }

// @Tags AccountPositionDaily
// @Summary 删除AccountPositionDaily
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除AccountPositionDaily"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /accountPositionDaily/deleteAccountPositionDaily [delete]
 export const deleteAccountPositionDailyByIds = (data) => {
     return service({
         url: "/accountPositionDaily/deleteAccountPositionDailyByIds",
         method: 'delete',
         data
     })
 }

// @Tags AccountPositionDaily
// @Summary 更新AccountPositionDaily
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AccountPositionDaily true "更新AccountPositionDaily"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /accountPositionDaily/updateAccountPositionDaily [put]
 export const updateAccountPositionDaily = (data) => {
     return service({
         url: "/accountPositionDaily/updateAccountPositionDaily",
         method: 'put',
         data
     })
 }


// @Tags AccountPositionDaily
// @Summary 用id查询AccountPositionDaily
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AccountPositionDaily true "用id查询AccountPositionDaily"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /accountPositionDaily/findAccountPositionDaily [get]
 export const findAccountPositionDaily = (params) => {
     return service({
         url: "/accountPositionDaily/findAccountPositionDaily",
         method: 'get',
         params
     })
 }


// @Tags AccountPositionDaily
// @Summary 分页获取AccountPositionDaily列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取AccountPositionDaily列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /accountPositionDaily/getAccountPositionDailyList [get]
 export const getAccountPositionDailyList = (params) => {
     return service({
         url: "/accountPositionDaily/getAccountPositionDailyList",
         method: 'get',
         params
     })
 }