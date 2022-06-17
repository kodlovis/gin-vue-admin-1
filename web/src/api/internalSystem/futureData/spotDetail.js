import service from '@/utils/request'

// @Tags SpotDetail
// @Summary 创建SpotDetail
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SpotDetail true "创建SpotDetail"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /spotDetail/createSpotDetail [post]
export const createSpotDetail = (data) => {
     return service({
         url: "/spotDetail/createSpotDetail",
         method: 'post',
         data
     })
 }


// @Tags SpotDetail
// @Summary 删除SpotDetail
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SpotDetail true "删除SpotDetail"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /spotDetail/deleteSpotDetail [delete]
 export const deleteSpotDetail = (data) => {
     return service({
         url: "/spotDetail/deleteSpotDetail",
         method: 'delete',
         data
     })
 }

// @Tags SpotDetail
// @Summary 删除SpotDetail
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除SpotDetail"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /spotDetail/deleteSpotDetail [delete]
 export const deleteSpotDetailByIds = (data) => {
     return service({
         url: "/spotDetail/deleteSpotDetailByIds",
         method: 'delete',
         data
     })
 }

// @Tags SpotDetail
// @Summary 更新SpotDetail
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SpotDetail true "更新SpotDetail"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /spotDetail/updateSpotDetail [put]
 export const updateSpotDetail = (data) => {
     return service({
         url: "/spotDetail/updateSpotDetail",
         method: 'put',
         data
     })
 }


// @Tags SpotDetail
// @Summary 用id查询SpotDetail
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SpotDetail true "用id查询SpotDetail"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /spotDetail/findSpotDetail [get]
 export const findSpotDetail = (params) => {
     return service({
         url: "/spotDetail/findSpotDetail",
         method: 'get',
         params
     })
 }


// @Tags SpotDetail
// @Summary 分页获取SpotDetail列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取SpotDetail列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /spotDetail/getSpotDetailList [get]
 export const getSpotDetailList = (params) => {
     return service({
         url: "/spotDetail/getSpotDetailList",
         method: 'get',
         params
     })
 }