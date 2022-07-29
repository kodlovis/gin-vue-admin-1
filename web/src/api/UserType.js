import service from '@/utils/request'

// @Tags UserType
// @Summary 创建UserType
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.UserType true "创建UserType"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /ut/createUserType [post]
export const createUserType = (data) => {
     return service({
         url: "/ut/createUserType",
         method: 'post',
         data
     })
 }


// @Tags UserType
// @Summary 删除UserType
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.UserType true "删除UserType"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /ut/deleteUserType [delete]
 export const deleteUserType = (data) => {
     return service({
         url: "/ut/deleteUserType",
         method: 'delete',
         data
     })
 }

// @Tags UserType
// @Summary 删除UserType
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除UserType"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /ut/deleteUserType [delete]
 export const deleteUserTypeByIds = (data) => {
     return service({
         url: "/ut/deleteUserTypeByIds",
         method: 'delete',
         data
     })
 }

// @Tags UserType
// @Summary 更新UserType
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.UserType true "更新UserType"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /ut/updateUserType [put]
 export const updateUserType = (data) => {
     return service({
         url: "/ut/updateUserType",
         method: 'put',
         data
     })
 }


// @Tags UserType
// @Summary 用id查询UserType
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.UserType true "用id查询UserType"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /ut/findUserType [get]
 export const findUserType = (params) => {
     return service({
         url: "/ut/findUserType",
         method: 'get',
         params
     })
 }


// @Tags UserType
// @Summary 分页获取UserType列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取UserType列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /ut/getUserTypeList [get]
 export const getUserTypeList = (params) => {
     return service({
         url: "/ut/getUserTypeList",
         method: 'get',
         params
     })
 }