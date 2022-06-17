import service from '@/utils/request'

// @Userss Users
// @Summary 创建Users
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Users true "创建Users"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /Users/createUsers [post]
export const createUsers = (data) => {
     return service({
         url: "/Users/createUsers",
         method: 'post',
         data
     })
 }


// @Userss Users
// @Summary 删除Users
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Users true "删除Users"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /Users/deleteUsers [delete]
 export const deleteUsers = (data) => {
     return service({
         url: "/Users/deleteUsers",
         method: 'delete',
         data
     })
 }

// @Userss Users
// @Summary 删除Users
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Users"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /Users/deleteUsers [delete]
 export const deleteUsersByIds = (data) => {
     return service({
         url: "/Users/deleteUsersByIds",
         method: 'delete',
         data
     })
 }

// @Userss Users
// @Summary 更新Users
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Users true "更新Users"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /Users/updateUsers [put]
 export const updateUsers = (data) => {
     return service({
         url: "/Users/updateUsers",
         method: 'put',
         data
     })
 }


// @Userss Users
// @Summary 用id查询Users
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Users true "用id查询Users"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /Users/findUsers [get]
 export const findUsers = (params) => {
     return service({
         url: "/Users/findUsers",
         method: 'get',
         params
     })
 }


// @Userss Users
// @Summary 分页获取Users列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取Users列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /Users/getUsersList [get]
 export const getUsersList = (params) => {
     return service({
         url: "/Users/getUsersList",
         method: 'get',
         params
     })
 }

 export const getLastUser = (params) => {
    return service({
        url: "/Users/getLastUser",
        method: 'get',
        params
    })
}

export const createUserMac = (data) => {
    return service({
        url: "/Users/createUserMac",
        method: 'post',
        data
    })
}
export const createUserProduct= (data) => {
    return service({
        url: "/Users/createUserProduct",
        method: 'post',
        data
    })
}

export const removeUserMacProduct = (data) => {
    return service({
        url: "/Users/removeUserMacProduct",
        method: 'delete',
        data
    })
}