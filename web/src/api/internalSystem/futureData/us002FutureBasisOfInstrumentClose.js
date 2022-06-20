import service from '@/utils/request'

// @Tags Us002FutureBasisOfInstrumentClose
// @Summary 创建Us002FutureBasisOfInstrumentClose
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Us002FutureBasisOfInstrumentClose true "创建Us002FutureBasisOfInstrumentClose"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /us002FutureBasisOfInstrumentClose/createUs002FutureBasisOfInstrumentClose [post]
export const createUs002FutureBasisOfInstrumentClose = (data) => {
     return service({
         url: "/us002FutureBasisOfInstrumentClose/createUs002FutureBasisOfInstrumentClose",
         method: 'post',
         data
     })
 }


// @Tags Us002FutureBasisOfInstrumentClose
// @Summary 删除Us002FutureBasisOfInstrumentClose
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Us002FutureBasisOfInstrumentClose true "删除Us002FutureBasisOfInstrumentClose"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /us002FutureBasisOfInstrumentClose/deleteUs002FutureBasisOfInstrumentClose [delete]
 export const deleteUs002FutureBasisOfInstrumentClose = (data) => {
     return service({
         url: "/us002FutureBasisOfInstrumentClose/deleteUs002FutureBasisOfInstrumentClose",
         method: 'delete',
         data
     })
 }

// @Tags Us002FutureBasisOfInstrumentClose
// @Summary 删除Us002FutureBasisOfInstrumentClose
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Us002FutureBasisOfInstrumentClose"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /us002FutureBasisOfInstrumentClose/deleteUs002FutureBasisOfInstrumentClose [delete]
 export const deleteUs002FutureBasisOfInstrumentCloseByIds = (data) => {
     return service({
         url: "/us002FutureBasisOfInstrumentClose/deleteUs002FutureBasisOfInstrumentCloseByIds",
         method: 'delete',
         data
     })
 }

// @Tags Us002FutureBasisOfInstrumentClose
// @Summary 更新Us002FutureBasisOfInstrumentClose
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Us002FutureBasisOfInstrumentClose true "更新Us002FutureBasisOfInstrumentClose"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /us002FutureBasisOfInstrumentClose/updateUs002FutureBasisOfInstrumentClose [put]
 export const updateUs002FutureBasisOfInstrumentClose = (data) => {
     return service({
         url: "/us002FutureBasisOfInstrumentClose/updateUs002FutureBasisOfInstrumentClose",
         method: 'put',
         data
     })
 }


// @Tags Us002FutureBasisOfInstrumentClose
// @Summary 用id查询Us002FutureBasisOfInstrumentClose
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Us002FutureBasisOfInstrumentClose true "用id查询Us002FutureBasisOfInstrumentClose"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /us002FutureBasisOfInstrumentClose/findUs002FutureBasisOfInstrumentClose [get]
 export const findUs002FutureBasisOfInstrumentClose = (params) => {
     return service({
         url: "/us002FutureBasisOfInstrumentClose/findUs002FutureBasisOfInstrumentClose",
         method: 'get',
         params
     })
 }


// @Tags Us002FutureBasisOfInstrumentClose
// @Summary 分页获取Us002FutureBasisOfInstrumentClose列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取Us002FutureBasisOfInstrumentClose列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /us002FutureBasisOfInstrumentClose/getUs002FutureBasisOfInstrumentCloseList [get]
 export const getUs002FutureBasisOfInstrumentCloseList = (params) => {
     return service({
         url: "/us002FutureBasisOfInstrumentClose/getUs002FutureBasisOfInstrumentCloseList",
         method: 'get',
         params
     })
 }
 
 export const getPositionDeliveryMonthInstrumentList = (params) => {
    return service({
        url: "/us002FutureBasisOfInstrumentClose/getPositionDeliveryMonthInstrumentList",
        method: 'get',
        params
    })
}
export const getBenchmarkInstrumentList = (params) => {
   return service({
       url: "/us002FutureBasisOfInstrumentClose/getBenchmarkInstrumentList",
       method: 'get',
       params
   })
}



export const getNoRiskValue = (params) => {
    return service({
        url: "/us002FutureBasisOfInstrumentClose/getNoRiskValue",
        method: 'get',
        params
    })
}




export const createBasisOfInstrumentCloseList = (data) => {
    return service({
        url: "/us002FutureBasisOfInstrumentClose/createBasisOfInstrumentCloseList",
        method: 'post',
        data
    })
}