import service from '@/utils/request'

// @Macs Mac
// @Summary 创建Mac
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Mac true "创建Mac"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /Mac/createMac [post]
export const createMac = (data) => {
     return service({
         url: "/Mac/createMac",
         method: 'post',
         data
     })
 }


// @Macs Mac
// @Summary 删除Mac
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Mac true "删除Mac"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /Mac/deleteMac [delete]
 export const deleteMac = (data) => {
     return service({
         url: "/Mac/deleteMac",
         method: 'delete',
         data
     })
 }

// @Macs Mac
// @Summary 删除Mac
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Mac"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /Mac/deleteMac [delete]
 export const deleteMacByIds = (data) => {
     return service({
         url: "/Mac/deleteMacByIds",
         method: 'delete',
         data
     })
 }

// @Macs Mac
// @Summary 更新Mac
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Mac true "更新Mac"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /Mac/updateMac [put]
 export const updateMac = (data) => {
     return service({
         url: "/Mac/updateMac",
         method: 'put',
         data
     })
 }


// @Macs Mac
// @Summary 用id查询Mac
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Mac true "用id查询Mac"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /Mac/findMac [get]
 export const findMac = (params) => {
     return service({
         url: "/Mac/findMac",
         method: 'get',
         params
     })
 }


// @Macs Mac
// @Summary 分页获取Mac列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取Mac列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /Mac/getMacList [get]
 export const getMacList = (params) => {
     return service({
         url: "/Mac/getMacList",
         method: 'get',
         params
     })
 }