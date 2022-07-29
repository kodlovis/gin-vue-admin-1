import service from '@/utils/request'

// @Tags FundConfirm
// @Summary 创建FundConfirm
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.FundConfirm true "创建FundConfirm"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /fundConfirmDate/createFundConfirm [post]
export const createFundConfirm = (data) => {
     return service({
         url: "/fundConfirmDate/createFundConfirm",
         method: 'post',
         data
     })
 }


// @Tags FundConfirm
// @Summary 删除FundConfirm
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.FundConfirm true "删除FundConfirm"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /fundConfirmDate/deleteFundConfirm [delete]
 export const deleteFundConfirm = (data) => {
     return service({
         url: "/fundConfirmDate/deleteFundConfirm",
         method: 'delete',
         data
     })
 }

// @Tags FundConfirm
// @Summary 删除FundConfirm
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除FundConfirm"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /fundConfirmDate/deleteFundConfirm [delete]
 export const deleteFundConfirmByIds = (data) => {
     return service({
         url: "/fundConfirmDate/deleteFundConfirmByIds",
         method: 'delete',
         data
     })
 }

// @Tags FundConfirm
// @Summary 更新FundConfirm
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.FundConfirm true "更新FundConfirm"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /fundConfirmDate/updateFundConfirm [put]
 export const updateFundConfirm = (data) => {
     return service({
         url: "/fundConfirmDate/updateFundConfirm",
         method: 'put',
         data
     })
 }


// @Tags FundConfirm
// @Summary 用id查询FundConfirm
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.FundConfirm true "用id查询FundConfirm"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /fundConfirmDate/findFundConfirm [get]
 export const findFundConfirm = (params) => {
     return service({
         url: "/fundConfirmDate/findFundConfirm",
         method: 'get',
         params
     })
 }

 export const getLast = () => {
    return service({
        url: "/fundConfirmDate/getLastConfirmDate",
        method: 'get',
    })
}

export const getCycleConfirmDate = (data) => {
    return service({
        url: "/fundConfirmDate/getCycleConfirmDate",
        method: 'post',
        data
    })
}
// @Tags FundConfirm
// @Summary 分页获取FundConfirm列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取FundConfirm列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /fundConfirmDate/getFundConfirmList [get]
 export const getFundConfirmList = (params) => {
     return service({
         url: "/fundConfirmDate/getFundConfirmList",
         method: 'get',
         params
     })
 }

 export const getConfirmSetting = (data) => {
    return service({
        url: "/fundConfirmDate/getConfirmSetting",
        method: 'post',
        data
    })
}

export const upConfirmSetting = (data) => {
    return service({
        url: "/fundConfirmDate/updateConfirmSetting",
        method: 'put',
        data
    })
}