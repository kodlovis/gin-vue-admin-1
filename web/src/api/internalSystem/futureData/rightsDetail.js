import service from '@/utils/request'

// @Tags RightsDetail
// @Summary 创建RightsDetail
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.RightsDetail true "创建RightsDetail"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /rightsDetail/createRightsDetail [post]
export const createRightsDetail = (data) => {
     return service({
         url: "/rightsDetail/createRightsDetail",
         method: 'post',
         data
     })
 }


// @Tags RightsDetail
// @Summary 删除RightsDetail
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.RightsDetail true "删除RightsDetail"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /rightsDetail/deleteRightsDetail [delete]
 export const deleteRightsDetail = (data) => {
     return service({
         url: "/rightsDetail/deleteRightsDetail",
         method: 'delete',
         data
     })
 }

// @Tags RightsDetail
// @Summary 删除RightsDetail
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除RightsDetail"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /rightsDetail/deleteRightsDetail [delete]
 export const deleteRightsDetailByIds = (data) => {
     return service({
         url: "/rightsDetail/deleteRightsDetailByIds",
         method: 'delete',
         data
     })
 }

// @Tags RightsDetail
// @Summary 更新RightsDetail
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.RightsDetail true "更新RightsDetail"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /rightsDetail/updateRightsDetail [put]
 export const updateRightsDetail = (data) => {
     return service({
         url: "/rightsDetail/updateRightsDetail",
         method: 'put',
         data
     })
 }


// @Tags RightsDetail
// @Summary 用id查询RightsDetail
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.RightsDetail true "用id查询RightsDetail"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /rightsDetail/findRightsDetail [get]
 export const findRightsDetail = (params) => {
     return service({
         url: "/rightsDetail/findRightsDetail",
         method: 'get',
         params
     })
 }


// @Tags RightsDetail
// @Summary 分页获取RightsDetail列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取RightsDetail列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /rightsDetail/getRightsDetailList [get]
 export const getRightsDetailList = (params) => {
     return service({
         url: "/rightsDetail/getRightsDetailList",
         method: 'get',
         params
     })
 }