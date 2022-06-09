import service from '@/utils/request'

// @Tags Us001DceBrokerPositionDaily
// @Summary 创建Us001DceBrokerPositionDaily
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Us001DceBrokerPositionDaily true "创建Us001DceBrokerPositionDaily"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /us001DceBrokerPositionDaily/createUs001DceBrokerPositionDaily [post]
export const createUs001DceBrokerPositionDaily = (data) => {
     return service({
         url: "/us001DceBrokerPositionDaily/createUs001DceBrokerPositionDaily",
         method: 'post',
         data
     })
 }


// @Tags Us001DceBrokerPositionDaily
// @Summary 删除Us001DceBrokerPositionDaily
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Us001DceBrokerPositionDaily true "删除Us001DceBrokerPositionDaily"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /us001DceBrokerPositionDaily/deleteUs001DceBrokerPositionDaily [delete]
 export const deleteUs001DceBrokerPositionDaily = (data) => {
     return service({
         url: "/us001DceBrokerPositionDaily/deleteUs001DceBrokerPositionDaily",
         method: 'delete',
         data
     })
 }

// @Tags Us001DceBrokerPositionDaily
// @Summary 删除Us001DceBrokerPositionDaily
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Us001DceBrokerPositionDaily"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /us001DceBrokerPositionDaily/deleteUs001DceBrokerPositionDaily [delete]
 export const deleteUs001DceBrokerPositionDailyByIds = (data) => {
     return service({
         url: "/us001DceBrokerPositionDaily/deleteUs001DceBrokerPositionDailyByIds",
         method: 'delete',
         data
     })
 }

// @Tags Us001DceBrokerPositionDaily
// @Summary 更新Us001DceBrokerPositionDaily
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Us001DceBrokerPositionDaily true "更新Us001DceBrokerPositionDaily"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /us001DceBrokerPositionDaily/updateUs001DceBrokerPositionDaily [put]
 export const updateUs001DceBrokerPositionDaily = (data) => {
     return service({
         url: "/us001DceBrokerPositionDaily/updateUs001DceBrokerPositionDaily",
         method: 'put',
         data
     })
 }


// @Tags Us001DceBrokerPositionDaily
// @Summary 用id查询Us001DceBrokerPositionDaily
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Us001DceBrokerPositionDaily true "用id查询Us001DceBrokerPositionDaily"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /us001DceBrokerPositionDaily/findUs001DceBrokerPositionDaily [get]
 export const findUs001DceBrokerPositionDaily = (params) => {
     return service({
         url: "/us001DceBrokerPositionDaily/findUs001DceBrokerPositionDaily",
         method: 'get',
         params
     })
 }


// @Tags Us001DceBrokerPositionDaily
// @Summary 分页获取Us001DceBrokerPositionDaily列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取Us001DceBrokerPositionDaily列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /us001DceBrokerPositionDaily/getUs001DceBrokerPositionDailyList [get]
 export const getUs001DceBrokerPositionDailyList = (params) => {
     return service({
         url: "/us001DceBrokerPositionDaily/getUs001DceBrokerPositionDailyList",
         method: 'get',
         params
     })
 }
 export const loadUs001DceBrokerPositionExcelData = () => {
     return service({
         url: "/us001DceBrokerPositionDaily/loadUs001DceBrokerPositionExcel",
         method: 'get'
     })
 }