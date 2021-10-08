import service from '@/utils/request'

// @Tags ForexFutureDetail
// @Summary 创建ForexFutureDetail
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ForexFutureDetail true "创建ForexFutureDetail"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /forexFutureDetail/createForexFutureDetail [post]
export const createForexFutureDetail = (data) => {
     return service({
         url: "/forexFutureDetail/createForexFutureDetail",
         method: 'post',
         data
     })
 }


// @Tags ForexFutureDetail
// @Summary 删除ForexFutureDetail
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ForexFutureDetail true "删除ForexFutureDetail"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /forexFutureDetail/deleteForexFutureDetail [delete]
 export const deleteForexFutureDetail = (data) => {
     return service({
         url: "/forexFutureDetail/deleteForexFutureDetail",
         method: 'delete',
         data
     })
 }

// @Tags ForexFutureDetail
// @Summary 删除ForexFutureDetail
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ForexFutureDetail"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /forexFutureDetail/deleteForexFutureDetail [delete]
 export const deleteForexFutureDetailByIds = (data) => {
     return service({
         url: "/forexFutureDetail/deleteForexFutureDetailByIds",
         method: 'delete',
         data
     })
 }

// @Tags ForexFutureDetail
// @Summary 更新ForexFutureDetail
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ForexFutureDetail true "更新ForexFutureDetail"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /forexFutureDetail/updateForexFutureDetail [put]
 export const updateForexFutureDetail = (data) => {
     return service({
         url: "/forexFutureDetail/updateForexFutureDetail",
         method: 'put',
         data
     })
 }


// @Tags ForexFutureDetail
// @Summary 用id查询ForexFutureDetail
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ForexFutureDetail true "用id查询ForexFutureDetail"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /forexFutureDetail/findForexFutureDetail [get]
 export const findForexFutureDetail = (params) => {
     return service({
         url: "/forexFutureDetail/findForexFutureDetail",
         method: 'get',
         params
     })
 }


// @Tags ForexFutureDetail
// @Summary 分页获取ForexFutureDetail列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取ForexFutureDetail列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /forexFutureDetail/getForexFutureDetailList [get]
 export const getForexFutureDetailList = (params) => {
     return service({
         url: "/forexFutureDetail/getForexFutureDetailList",
         method: 'get',
         params
     })
 }