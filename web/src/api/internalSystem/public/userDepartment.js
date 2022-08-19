import service from '@/utils/request'

// @Tags UserDepartment
// @Summary 创建UserDepartment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.UserDepartment true "创建UserDepartment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /userDepartment/createUserDepartment [post]
export const createUserDepartment = (data) => {
     return service({
         url: "/userDepartment/createUserDepartment",
         method: 'post',
         data
     })
 }


// @Tags UserDepartment
// @Summary 删除UserDepartment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.UserDepartment true "删除UserDepartment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /userDepartment/deleteUserDepartment [delete]
 export const deleteUserDepartment = (data) => {
     return service({
         url: "/userDepartment/deleteUserDepartment",
         method: 'delete',
         data
     })
 }

// @Tags UserDepartment
// @Summary 删除UserDepartment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除UserDepartment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /userDepartment/deleteUserDepartment [delete]
 export const deleteUserDepartmentByIds = (data) => {
     return service({
         url: "/userDepartment/deleteUserDepartmentByIds",
         method: 'delete',
         data
     })
 }

// @Tags UserDepartment
// @Summary 更新UserDepartment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.UserDepartment true "更新UserDepartment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /userDepartment/updateUserDepartment [put]
 export const updateUserDepartment = (data) => {
     return service({
         url: "/userDepartment/updateUserDepartment",
         method: 'put',
         data
     })
 }


// @Tags UserDepartment
// @Summary 用id查询UserDepartment
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.UserDepartment true "用id查询UserDepartment"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /userDepartment/findUserDepartment [get]
 export const findUserDepartment = (params) => {
     return service({
         url: "/userDepartment/findUserDepartment",
         method: 'get',
         params
     })
 }


// @Tags UserDepartment
// @Summary 分页获取UserDepartment列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取UserDepartment列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /userDepartment/getUserDepartmentList [get]
 export const getUserDepartmentList = (params) => {
     return service({
         url: "/userDepartment/getUserDepartmentList",
         method: 'get',
         params
     })
 }

 

 export const getUserDepartmentListByID = (params) => {
    return service({
        url: "/userDepartment/getUserDepartmentListByID",
        method: 'get',
        params
    })
}