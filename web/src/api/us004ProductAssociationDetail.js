import service from '@/utils/request'

// @Tags Us004ProductAssociationDetail
// @Summary 创建Us004ProductAssociationDetail
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Us004ProductAssociationDetail true "创建Us004ProductAssociationDetail"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /us004ProductAssociationDetail/createUs004ProductAssociationDetail [post]
export const createUs004ProductAssociationDetail = (data) => {
     return service({
         url: "/us004ProductAssociationDetail/createUs004ProductAssociationDetail",
         method: 'post',
         data
     })
 }


// @Tags Us004ProductAssociationDetail
// @Summary 删除Us004ProductAssociationDetail
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Us004ProductAssociationDetail true "删除Us004ProductAssociationDetail"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /us004ProductAssociationDetail/deleteUs004ProductAssociationDetail [delete]
 export const deleteUs004ProductAssociationDetail = (data) => {
     return service({
         url: "/us004ProductAssociationDetail/deleteUs004ProductAssociationDetail",
         method: 'delete',
         data
     })
 }

// @Tags Us004ProductAssociationDetail
// @Summary 删除Us004ProductAssociationDetail
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Us004ProductAssociationDetail"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /us004ProductAssociationDetail/deleteUs004ProductAssociationDetail [delete]
 export const deleteUs004ProductAssociationDetailByIds = (data) => {
     return service({
         url: "/us004ProductAssociationDetail/deleteUs004ProductAssociationDetailByIds",
         method: 'delete',
         data
     })
 }

// @Tags Us004ProductAssociationDetail
// @Summary 更新Us004ProductAssociationDetail
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Us004ProductAssociationDetail true "更新Us004ProductAssociationDetail"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /us004ProductAssociationDetail/updateUs004ProductAssociationDetail [put]
 export const updateUs004ProductAssociationDetail = (data) => {
     return service({
         url: "/us004ProductAssociationDetail/updateUs004ProductAssociationDetail",
         method: 'put',
         data
     })
 }


// @Tags Us004ProductAssociationDetail
// @Summary 用id查询Us004ProductAssociationDetail
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Us004ProductAssociationDetail true "用id查询Us004ProductAssociationDetail"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /us004ProductAssociationDetail/findUs004ProductAssociationDetail [get]
 export const findUs004ProductAssociationDetail = (params) => {
     return service({
         url: "/us004ProductAssociationDetail/findUs004ProductAssociationDetail",
         method: 'get',
         params
     })
 }


// @Tags Us004ProductAssociationDetail
// @Summary 分页获取Us004ProductAssociationDetail列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取Us004ProductAssociationDetail列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /us004ProductAssociationDetail/getUs004ProductAssociationDetailList [get]
 export const getUs004ProductAssociationDetailList = (params) => {
     return service({
         url: "/us004ProductAssociationDetail/getUs004ProductAssociationDetailList",
         method: 'get',
         params
     })
 }