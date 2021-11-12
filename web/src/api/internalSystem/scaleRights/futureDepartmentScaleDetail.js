import service from '@/utils/request'

// @Tags FutureDepartmentScaleDetail
// @Summary 创建FutureDepartmentScaleDetail
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.FutureDepartmentScaleDetail true "创建FutureDepartmentScaleDetail"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /futureDepartmentScaleDetail/createFutureDepartmentScaleDetail [post]
export const createFutureDepartmentScaleDetail = (data) => {
     return service({
         url: "/futureDepartmentScaleDetail/createFutureDepartmentScaleDetail",
         method: 'post',
         data
     })
 }


// @Tags FutureDepartmentScaleDetail
// @Summary 删除FutureDepartmentScaleDetail
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.FutureDepartmentScaleDetail true "删除FutureDepartmentScaleDetail"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /futureDepartmentScaleDetail/deleteFutureDepartmentScaleDetail [delete]
 export const deleteFutureDepartmentScaleDetail = (data) => {
     return service({
         url: "/futureDepartmentScaleDetail/deleteFutureDepartmentScaleDetail",
         method: 'delete',
         data
     })
 }

// @Tags FutureDepartmentScaleDetail
// @Summary 删除FutureDepartmentScaleDetail
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除FutureDepartmentScaleDetail"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /futureDepartmentScaleDetail/deleteFutureDepartmentScaleDetail [delete]
 export const deleteFutureDepartmentScaleDetailByIds = (data) => {
     return service({
         url: "/futureDepartmentScaleDetail/deleteFutureDepartmentScaleDetailByIds",
         method: 'delete',
         data
     })
 }

// @Tags FutureDepartmentScaleDetail
// @Summary 更新FutureDepartmentScaleDetail
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.FutureDepartmentScaleDetail true "更新FutureDepartmentScaleDetail"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /futureDepartmentScaleDetail/updateFutureDepartmentScaleDetail [put]
 export const updateFutureDepartmentScaleDetail = (data) => {
     return service({
         url: "/futureDepartmentScaleDetail/updateFutureDepartmentScaleDetail",
         method: 'put',
         data
     })
 }


// @Tags FutureDepartmentScaleDetail
// @Summary 用id查询FutureDepartmentScaleDetail
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.FutureDepartmentScaleDetail true "用id查询FutureDepartmentScaleDetail"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /futureDepartmentScaleDetail/findFutureDepartmentScaleDetail [get]
 export const findFutureDepartmentScaleDetail = (params) => {
     return service({
         url: "/futureDepartmentScaleDetail/findFutureDepartmentScaleDetail",
         method: 'get',
         params
     })
 }


// @Tags FutureDepartmentScaleDetail
// @Summary 分页获取FutureDepartmentScaleDetail列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取FutureDepartmentScaleDetail列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /futureDepartmentScaleDetail/getFutureDepartmentScaleDetailList [get]
 export const getFutureDepartmentScaleDetailList = (params) => {
     return service({
         url: "/futureDepartmentScaleDetail/getFutureDepartmentScaleDetailList",
         method: 'get',
         params
     })
 }