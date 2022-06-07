import service from '@/utils/request'

// @Tags Us001CzceBrokerPositionDaily
// @Summary 创建Us001CzceBrokerPositionDaily
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Us001CzceBrokerPositionDaily true "创建Us001CzceBrokerPositionDaily"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /us001CzceBrokerPositionDaily/createUs001CzceBrokerPositionDaily [post]
export const createUs001CzceBrokerPositionDaily = (data) => {
     return service({
         url: "/us001CzceBrokerPositionDaily/createUs001CzceBrokerPositionDaily",
         method: 'post',
         data
     })
 }


// @Tags Us001CzceBrokerPositionDaily
// @Summary 删除Us001CzceBrokerPositionDaily
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Us001CzceBrokerPositionDaily true "删除Us001CzceBrokerPositionDaily"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /us001CzceBrokerPositionDaily/deleteUs001CzceBrokerPositionDaily [delete]
 export const deleteUs001CzceBrokerPositionDaily = (data) => {
     return service({
         url: "/us001CzceBrokerPositionDaily/deleteUs001CzceBrokerPositionDaily",
         method: 'delete',
         data
     })
 }

// @Tags Us001CzceBrokerPositionDaily
// @Summary 删除Us001CzceBrokerPositionDaily
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Us001CzceBrokerPositionDaily"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /us001CzceBrokerPositionDaily/deleteUs001CzceBrokerPositionDaily [delete]
 export const deleteUs001CzceBrokerPositionDailyByIds = (data) => {
     return service({
         url: "/us001CzceBrokerPositionDaily/deleteUs001CzceBrokerPositionDailyByIds",
         method: 'delete',
         data
     })
 }

// @Tags Us001CzceBrokerPositionDaily
// @Summary 更新Us001CzceBrokerPositionDaily
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Us001CzceBrokerPositionDaily true "更新Us001CzceBrokerPositionDaily"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /us001CzceBrokerPositionDaily/updateUs001CzceBrokerPositionDaily [put]
 export const updateUs001CzceBrokerPositionDaily = (data) => {
     return service({
         url: "/us001CzceBrokerPositionDaily/updateUs001CzceBrokerPositionDaily",
         method: 'put',
         data
     })
 }


// @Tags Us001CzceBrokerPositionDaily
// @Summary 用id查询Us001CzceBrokerPositionDaily
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Us001CzceBrokerPositionDaily true "用id查询Us001CzceBrokerPositionDaily"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /us001CzceBrokerPositionDaily/findUs001CzceBrokerPositionDaily [get]
 export const findUs001CzceBrokerPositionDaily = (params) => {
     return service({
         url: "/us001CzceBrokerPositionDaily/findUs001CzceBrokerPositionDaily",
         method: 'get',
         params
     })
 }


// @Tags Us001CzceBrokerPositionDaily
// @Summary 分页获取Us001CzceBrokerPositionDaily列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取Us001CzceBrokerPositionDaily列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /us001CzceBrokerPositionDaily/getUs001CzceBrokerPositionDailyList [get]
 export const getUs001CzceBrokerPositionDailyList = (params) => {
     return service({
         url: "/us001CzceBrokerPositionDaily/getUs001CzceBrokerPositionDailyList",
         method: 'get',
         params
     })
 }
 export const loadUs001CzceBrokerPositionExcelData = () => {
     return service({
         url: "/us001CzceBrokerPositionDaily/loadUs001CzceBrokerPositionExcel",
         method: 'get'
     })
 }