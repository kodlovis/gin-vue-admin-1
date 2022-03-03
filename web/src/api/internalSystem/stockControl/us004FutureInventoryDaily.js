import service from '@/utils/request'

// @Tags Us004FutureInventoryDaily
// @Summary 创建Us004FutureInventoryDaily
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Us004FutureInventoryDaily true "创建Us004FutureInventoryDaily"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /us004FutureInventoryDaily/createUs004FutureInventoryDaily [post]
export const createUs004FutureInventoryDaily = (data) => {
     return service({
         url: "/us004FutureInventoryDaily/createUs004FutureInventoryDaily",
         method: 'post',
         data
     })
 }


// @Tags Us004FutureInventoryDaily
// @Summary 删除Us004FutureInventoryDaily
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Us004FutureInventoryDaily true "删除Us004FutureInventoryDaily"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /us004FutureInventoryDaily/deleteUs004FutureInventoryDaily [delete]
 export const deleteUs004FutureInventoryDaily = (data) => {
     return service({
         url: "/us004FutureInventoryDaily/deleteUs004FutureInventoryDaily",
         method: 'delete',
         data
     })
 }

// @Tags Us004FutureInventoryDaily
// @Summary 删除Us004FutureInventoryDaily
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Us004FutureInventoryDaily"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /us004FutureInventoryDaily/deleteUs004FutureInventoryDaily [delete]
 export const deleteUs004FutureInventoryDailyByIds = (data) => {
     return service({
         url: "/us004FutureInventoryDaily/deleteUs004FutureInventoryDailyByIds",
         method: 'delete',
         data
     })
 }

// @Tags Us004FutureInventoryDaily
// @Summary 更新Us004FutureInventoryDaily
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Us004FutureInventoryDaily true "更新Us004FutureInventoryDaily"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /us004FutureInventoryDaily/updateUs004FutureInventoryDaily [put]
 export const updateUs004FutureInventoryDaily = (data) => {
     return service({
         url: "/us004FutureInventoryDaily/updateUs004FutureInventoryDaily",
         method: 'put',
         data
     })
 }


// @Tags Us004FutureInventoryDaily
// @Summary 用id查询Us004FutureInventoryDaily
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Us004FutureInventoryDaily true "用id查询Us004FutureInventoryDaily"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /us004FutureInventoryDaily/findUs004FutureInventoryDaily [get]
 export const findUs004FutureInventoryDaily = (params) => {
     return service({
         url: "/us004FutureInventoryDaily/findUs004FutureInventoryDaily",
         method: 'get',
         params
     })
 }


// @Tags Us004FutureInventoryDaily
// @Summary 分页获取Us004FutureInventoryDaily列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取Us004FutureInventoryDaily列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /us004FutureInventoryDaily/getUs004FutureInventoryDailyList [get]
 export const getUs004FutureInventoryDailyList = (params) => {
     return service({
         url: "/us004FutureInventoryDaily/getUs004FutureInventoryDailyList",
         method: 'get',
         params
     })

 }

 export const getUs004FutureInventoryDailyType = (params) => {
    return service({
        url: "/us004FutureInventoryDaily/getUs004FutureInventoryDailyType",
        method: 'get',
        params
    })
 }
 
 export const loadInventoryExcelData = () => {
    return service({
        url: "/us004FutureInventoryDaily/loadInventoryExcelData",
        method: 'get'
    })
}