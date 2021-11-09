import service from '@/utils/request'

// @Tags AccountInfo
// @Summary 创建AccountInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AccountInfo true "创建AccountInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /accountInfo/createAccountInfo [post]
export const createAccountInfo = (data) => {
     return service({
         url: "/accountInfo/createAccountInfo",
         method: 'post',
         data
     })
 }


// @Tags AccountInfo
// @Summary 删除AccountInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AccountInfo true "删除AccountInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /accountInfo/deleteAccountInfo [delete]
 export const deleteAccountInfo = (data) => {
     return service({
         url: "/accountInfo/deleteAccountInfo",
         method: 'delete',
         data
     })
 }

// @Tags AccountInfo
// @Summary 删除AccountInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除AccountInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /accountInfo/deleteAccountInfo [delete]
 export const deleteAccountInfoByIds = (data) => {
     return service({
         url: "/accountInfo/deleteAccountInfoByIds",
         method: 'delete',
         data
     })
 }

// @Tags AccountInfo
// @Summary 更新AccountInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AccountInfo true "更新AccountInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /accountInfo/updateAccountInfo [put]
 export const updateAccountInfo = (data) => {
     return service({
         url: "/accountInfo/updateAccountInfo",
         method: 'put',
         data
     })
 }


// @Tags AccountInfo
// @Summary 用id查询AccountInfo
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AccountInfo true "用id查询AccountInfo"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /accountInfo/findAccountInfo [get]
 export const findAccountInfo = (params) => {
     return service({
         url: "/accountInfo/findAccountInfo",
         method: 'get',
         params
     })
 }


// @Tags AccountInfo
// @Summary 分页获取AccountInfo列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取AccountInfo列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /accountInfo/getAccountInfoList [get]
 export const getAccountInfoList = (params) => {
     return service({
         url: "/accountInfo/getAccountInfoList",
         method: 'get',
         params
     })
 }