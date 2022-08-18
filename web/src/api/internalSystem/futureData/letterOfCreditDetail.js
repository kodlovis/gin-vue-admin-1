import service from '@/utils/request'

// @Tags LetterOfCreditDetail
// @Summary 创建LetterOfCreditDetail
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.LetterOfCreditDetail true "创建LetterOfCreditDetail"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /letterOfCreditDetail/createLetterOfCreditDetail [post]
export const createLetterOfCreditDetail = (data) => {
     return service({
         url: "/letterOfCreditDetail/createLetterOfCreditDetail",
         method: 'post',
         data
     })
 }


// @Tags LetterOfCreditDetail
// @Summary 删除LetterOfCreditDetail
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.LetterOfCreditDetail true "删除LetterOfCreditDetail"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /letterOfCreditDetail/deleteLetterOfCreditDetail [delete]
 export const deleteLetterOfCreditDetail = (data) => {
     return service({
         url: "/letterOfCreditDetail/deleteLetterOfCreditDetail",
         method: 'delete',
         data
     })
 }

// @Tags LetterOfCreditDetail
// @Summary 删除LetterOfCreditDetail
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除LetterOfCreditDetail"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /letterOfCreditDetail/deleteLetterOfCreditDetail [delete]
 export const deleteLetterOfCreditDetailByIds = (data) => {
     return service({
         url: "/letterOfCreditDetail/deleteLetterOfCreditDetailByIds",
         method: 'delete',
         data
     })
 }

// @Tags LetterOfCreditDetail
// @Summary 更新LetterOfCreditDetail
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.LetterOfCreditDetail true "更新LetterOfCreditDetail"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /letterOfCreditDetail/updateLetterOfCreditDetail [put]
 export const updateLetterOfCreditDetail = (data) => {
     return service({
         url: "/letterOfCreditDetail/updateLetterOfCreditDetail",
         method: 'put',
         data
     })
 }


// @Tags LetterOfCreditDetail
// @Summary 用id查询LetterOfCreditDetail
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.LetterOfCreditDetail true "用id查询LetterOfCreditDetail"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /letterOfCreditDetail/findLetterOfCreditDetail [get]
 export const findLetterOfCreditDetail = (params) => {
     return service({
         url: "/letterOfCreditDetail/findLetterOfCreditDetail",
         method: 'get',
         params
     })
 }


// @Tags LetterOfCreditDetail
// @Summary 分页获取LetterOfCreditDetail列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取LetterOfCreditDetail列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /letterOfCreditDetail/getLetterOfCreditDetailList [get]
 export const getLetterOfCreditDetailList = (params) => {
     return service({
         url: "/letterOfCreditDetail/getLetterOfCreditDetailList",
         method: 'get',
         params
     })
 }