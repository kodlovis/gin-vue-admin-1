import service from '@/utils/request'

// @Tags RemainingSum
// @Summary 创建RemainingSum
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.RemainingSum true "创建RemainingSum"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /remainingSum/createRemainingSum [post]
export const createRemainingSum = (data) => {
     return service({
         url: "/remainingSum/createRemainingSum",
         method: 'post',
         data
     })
 }
 

// @Tags RemainingSum
// @Summary 删除RemainingSum
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.RemainingSum true "删除RemainingSum"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /remainingSum/deleteRemainingSum [delete]
 export const deleteRemainingSum = (data) => {
     return service({
         url: "/remainingSum/deleteRemainingSum",
         method: 'delete',
         data
     })
 }

// @Tags RemainingSum
// @Summary 删除RemainingSum
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除RemainingSum"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /remainingSum/deleteRemainingSum [delete]
 export const deleteRemainingSumByIds = (data) => {
     return service({
         url: "/remainingSum/deleteRemainingSumByIds",
         method: 'delete',
         data
     })
 }

// @Tags RemainingSum
// @Summary 更新RemainingSum
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.RemainingSum true "更新RemainingSum"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /remainingSum/updateRemainingSum [put]
 export const updateRemainingSum = (data) => {
     return service({
         url: "/remainingSum/updateRemainingSum",
         method: 'put',
         data
     })
 }


// @Tags RemainingSum
// @Summary 用id查询RemainingSum
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.RemainingSum true "用id查询RemainingSum"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /remainingSum/findRemainingSum [get]
 export const findRemainingSum = (params) => {
     return service({
         url: "/remainingSum/findRemainingSum",
         method: 'get',
         params
     })
 }


// @Tags RemainingSum
// @Summary 分页获取RemainingSum列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取RemainingSum列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /remainingSum/getRemainingSumList [get]
 export const getRemainingSumList = (params) => {
     return service({
         url: "/remainingSum/getRemainingSumList",
         method: 'get',
         params
     })
 }

 export const getOneRemainingSum = (data) => {
    return service({
        url: "/remainingSum/getOneRemainingSum",
        method: 'post',
        data
    })
}
export const getwarehouse = (data) => {
    return service({
        url: "/remainingSum/getwarehouse",
        method: 'post',
        data
    })
}

export const getFundRole = (data) => {
    return service({
        url: "/getfund/role",
        method: 'post',
        data
    })
}
