import service from '@/utils/request'

// @Tags DepartmentAccountProduct
// @Summary 创建DepartmentAccountProduct
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DepartmentAccountProduct true "创建DepartmentAccountProduct"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /departmentAccountProduct/createDepartmentAccountProduct [post]
export const createDepartmentAccountProduct = (data) => {
     return service({
         url: "/departmentAccountProduct/createDepartmentAccountProduct",
         method: 'post',
         data
     })
 }


// @Tags DepartmentAccountProduct
// @Summary 删除DepartmentAccountProduct
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DepartmentAccountProduct true "删除DepartmentAccountProduct"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /departmentAccountProduct/deleteDepartmentAccountProduct [delete]
 export const deleteDepartmentAccountProduct = (data) => {
     return service({
         url: "/departmentAccountProduct/deleteDepartmentAccountProduct",
         method: 'delete',
         data
     })
 }

// @Tags DepartmentAccountProduct
// @Summary 删除DepartmentAccountProduct
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除DepartmentAccountProduct"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /departmentAccountProduct/deleteDepartmentAccountProduct [delete]
 export const deleteDepartmentAccountProductByIds = (data) => {
     return service({
         url: "/departmentAccountProduct/deleteDepartmentAccountProductByIds",
         method: 'delete',
         data
     })
 }

// @Tags DepartmentAccountProduct
// @Summary 更新DepartmentAccountProduct
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DepartmentAccountProduct true "更新DepartmentAccountProduct"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /departmentAccountProduct/updateDepartmentAccountProduct [put]
 export const updateDepartmentAccountProduct = (data) => {
     return service({
         url: "/departmentAccountProduct/updateDepartmentAccountProduct",
         method: 'put',
         data
     })
 }


// @Tags DepartmentAccountProduct
// @Summary 用id查询DepartmentAccountProduct
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DepartmentAccountProduct true "用id查询DepartmentAccountProduct"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /departmentAccountProduct/findDepartmentAccountProduct [get]
 export const findDepartmentAccountProduct = (params) => {
     return service({
         url: "/departmentAccountProduct/findDepartmentAccountProduct",
         method: 'get',
         params
     })
 }


// @Tags DepartmentAccountProduct
// @Summary 分页获取DepartmentAccountProduct列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取DepartmentAccountProduct列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /departmentAccountProduct/getDepartmentAccountProductList [get]
 export const getDepartmentAccountProductList = (params) => {
     return service({
         url: "/departmentAccountProduct/getDepartmentAccountProductList",
         method: 'get',
         params
     })
 }