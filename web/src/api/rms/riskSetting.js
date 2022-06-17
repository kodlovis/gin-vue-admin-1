import service from '@/utils/request'

// @RiskSettings RiskSetting
// @Summary 创建RiskSetting
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.RiskSetting true "创建RiskSetting"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /RiskSetting/createRiskSetting [post]
export const createRiskSetting = (data) => {
     return service({
         url: "/RiskSetting/createRiskSetting",
         method: 'post',
         data
     })
 }


// @RiskSettings RiskSetting
// @Summary 删除RiskSetting
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.RiskSetting true "删除RiskSetting"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /RiskSetting/deleteRiskSetting [delete]
 export const deleteRiskSetting = (data) => {
     return service({
         url: "/RiskSetting/deleteRiskSetting",
         method: 'delete',
         data
     })
 }

// @RiskSettings RiskSetting
// @Summary 删除RiskSetting
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除RiskSetting"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /RiskSetting/deleteRiskSetting [delete]
 export const deleteRiskSettingByIds = (data) => {
     return service({
         url: "/RiskSetting/deleteRiskSettingByIds",
         method: 'delete',
         data
     })
 }

// @RiskSettings RiskSetting
// @Summary 更新RiskSetting
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.RiskSetting true "更新RiskSetting"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /RiskSetting/updateRiskSetting [put]
 export const updateRiskSetting = (data) => {
     return service({
         url: "/RiskSetting/updateRiskSetting",
         method: 'put',
         data
     })
 }


// @RiskSettings RiskSetting
// @Summary 用id查询RiskSetting
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.RiskSetting true "用id查询RiskSetting"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /RiskSetting/findRiskSetting [get]
 export const findRiskSetting = (params) => {
     return service({
         url: "/RiskSetting/findRiskSetting",
         method: 'get',
         params
     })
 }


// @RiskSettings RiskSetting
// @Summary 分页获取RiskSetting列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取RiskSetting列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /RiskSetting/getRiskSettingList [get]
 export const getRiskSettingList = (params) => {
     return service({
         url: "/RiskSetting/getRiskSettingList",
         method: 'get',
         params
     })
 }