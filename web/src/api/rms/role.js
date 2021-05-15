import service from '@/utils/request'

// @Roles Role
// @Summary 创建Role
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Role true "创建Role"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /Role/createRole [post]
export const createRole = (data) => {
     return service({
         url: "/Role/createRole",
         method: 'post',
         data
     })
 }


// @Roles Role
// @Summary 删除Role
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Role true "删除Role"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /Role/deleteRole [delete]
 export const deleteRole = (data) => {
     return service({
         url: "/Role/deleteRole",
         method: 'delete',
         data
     })
 }

// @Roles Role
// @Summary 删除Role
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Role"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /Role/deleteRole [delete]
 export const deleteRoleByIds = (data) => {
     return service({
         url: "/Role/deleteRoleByIds",
         method: 'delete',
         data
     })
 }

// @Roles Role
// @Summary 更新Role
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Role true "更新Role"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /Role/updateRole [put]
 export const updateRole = (data) => {
     return service({
         url: "/Role/updateRole",
         method: 'put',
         data
     })
 }


// @Roles Role
// @Summary 用id查询Role
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Role true "用id查询Role"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /Role/findRole [get]
 export const findRole = (params) => {
     return service({
         url: "/Role/findRole",
         method: 'get',
         params
     })
 }


// @Roles Role
// @Summary 分页获取Role列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取Role列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /Role/getRoleList [get]
 export const getRoleList = (params) => {
     return service({
         url: "/Role/getRoleList",
         method: 'get',
         params
     })
 }