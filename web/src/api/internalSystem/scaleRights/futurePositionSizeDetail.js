import service from '@/utils/request'

// @Tags FuturePositionSizeDetail
// @Summary 创建FuturePositionSizeDetail
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.FuturePositionSizeDetail true "创建FuturePositionSizeDetail"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /futurePositionSizeDetail/createFuturePositionSizeDetail [post]
export const createFuturePositionSizeDetail = (data) => {
     return service({
         url: "/futurePositionSizeDetail/createFuturePositionSizeDetail",
         method: 'post',
         data
     })
 }


// @Tags FuturePositionSizeDetail
// @Summary 删除FuturePositionSizeDetail
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.FuturePositionSizeDetail true "删除FuturePositionSizeDetail"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /futurePositionSizeDetail/deleteFuturePositionSizeDetail [delete]
 export const deleteFuturePositionSizeDetail = (data) => {
     return service({
         url: "/futurePositionSizeDetail/deleteFuturePositionSizeDetail",
         method: 'delete',
         data
     })
 }

// @Tags FuturePositionSizeDetail
// @Summary 删除FuturePositionSizeDetail
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除FuturePositionSizeDetail"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /futurePositionSizeDetail/deleteFuturePositionSizeDetail [delete]
 export const deleteFuturePositionSizeDetailByIds = (data) => {
     return service({
         url: "/futurePositionSizeDetail/deleteFuturePositionSizeDetailByIds",
         method: 'delete',
         data
     })
 }

// @Tags FuturePositionSizeDetail
// @Summary 更新FuturePositionSizeDetail
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.FuturePositionSizeDetail true "更新FuturePositionSizeDetail"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /futurePositionSizeDetail/updateFuturePositionSizeDetail [put]
 export const updateFuturePositionSizeDetail = (data) => {
     return service({
         url: "/futurePositionSizeDetail/updateFuturePositionSizeDetail",
         method: 'put',
         data
     })
 }


// @Tags FuturePositionSizeDetail
// @Summary 用id查询FuturePositionSizeDetail
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.FuturePositionSizeDetail true "用id查询FuturePositionSizeDetail"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /futurePositionSizeDetail/findFuturePositionSizeDetail [get]
 export const findFuturePositionSizeDetail = (params) => {
     return service({
         url: "/futurePositionSizeDetail/findFuturePositionSizeDetail",
         method: 'get',
         params
     })
 }


// @Tags FuturePositionSizeDetail
// @Summary 分页获取FuturePositionSizeDetail列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取FuturePositionSizeDetail列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /futurePositionSizeDetail/getFuturePositionSizeDetailList [get]
 export const getFuturePositionSizeDetailList = (params) => {
     return service({
         url: "/futurePositionSizeDetail/getFuturePositionSizeDetailList",
         method: 'get',
         params
     })
 }