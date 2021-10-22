import service from '@/utils/request'

// @Tags PositionRewardRule
// @Summary 创建PositionRewardRule
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PositionRewardRule true "创建PositionRewardRule"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /positionRewardRule/createPositionRewardRule [post]
export const createPositionRewardRule = (data) => {
     return service({
         url: "/positionRewardRule/createPositionRewardRule",
         method: 'post',
         data
     })
 }


// @Tags PositionRewardRule
// @Summary 删除PositionRewardRule
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PositionRewardRule true "删除PositionRewardRule"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /positionRewardRule/deletePositionRewardRule [delete]
 export const deletePositionRewardRule = (data) => {
     return service({
         url: "/positionRewardRule/deletePositionRewardRule",
         method: 'delete',
         data
     })
 }

// @Tags PositionRewardRule
// @Summary 删除PositionRewardRule
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除PositionRewardRule"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /positionRewardRule/deletePositionRewardRule [delete]
 export const deletePositionRewardRuleByIds = (data) => {
     return service({
         url: "/positionRewardRule/deletePositionRewardRuleByIds",
         method: 'delete',
         data
     })
 }

// @Tags PositionRewardRule
// @Summary 更新PositionRewardRule
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PositionRewardRule true "更新PositionRewardRule"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /positionRewardRule/updatePositionRewardRule [put]
 export const updatePositionRewardRule = (data) => {
     return service({
         url: "/positionRewardRule/updatePositionRewardRule",
         method: 'put',
         data
     })
 }


// @Tags PositionRewardRule
// @Summary 用id查询PositionRewardRule
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PositionRewardRule true "用id查询PositionRewardRule"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /positionRewardRule/findPositionRewardRule [get]
 export const findPositionRewardRule = (params) => {
     return service({
         url: "/positionRewardRule/findPositionRewardRule",
         method: 'get',
         params
     })
 }


// @Tags PositionRewardRule
// @Summary 分页获取PositionRewardRule列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取PositionRewardRule列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /positionRewardRule/getPositionRewardRuleList [get]
 export const getPositionRewardRuleList = (params) => {
     return service({
         url: "/positionRewardRule/getPositionRewardRuleList",
         method: 'get',
         params
     })
 }