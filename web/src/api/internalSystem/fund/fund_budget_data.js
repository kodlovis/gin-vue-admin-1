import service from '@/utils/request'

// @Tags BudgetData
// @Summary 创建BudgetData
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BudgetData true "创建BudgetData"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /budget_data/createBudgetData [post]
export const createBudgetData = (data) => {
     return service({
         url: "/budget_data/createBudgetData",
         method: 'post',
         data
     })
 }
 //录入预算数据
 export const createBudgetDatas = (data) => {
    return service({
      url: '/budget_data/createBudgetDatas',
      method: 'post',
      data
    })
  }
  //查询当日数据
  export const clickBudgetData = (data)=>{
    return service({
      url:'/budget_data/clickBudgetData',
      method:'post',
      data
    })
  }
    //查询今日录入数据
    export const getCycleData = (data)=>{
      return service({
        url:'/budget_data/getCycleBudgetData',
        method:'post',
        data
      })
    }
        //查询今日录入数据
        export const getAll = (data)=>{
          return service({
            url:'/budget_data/getReport',
            method:'post',
            data
          })
        }
  //获取确认数据
  export const getConfirmData = (data)=>{
    return service({
      url:'/budget_data/searchConfirmBudgetDate',
      method:'post',
      data
    })
  }
  //数据确认提交
  export const upConfirmData = (data)=>{
    return service({
      url:'/budget_data/createConfirmBudgetDate',
      method:'post',
      data
    })
  }
  //数据确认情况查看
  export const getOneConfirmData = (data)=>{
    return service({
      url:'/budget_data/getConfirmBudgetDate',
      method:'post',
      data
    })
  }

// @Tags BudgetData
// @Summary 删除BudgetData
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BudgetData true "删除BudgetData"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /budget_data/deleteBudgetData [delete]
 export const deleteBudgetData = (data) => {
     return service({
         url: "/budget_data/deleteBudgetData",
         method: 'delete',
         data
     })
 }

// @Tags BudgetData
// @Summary 删除BudgetData
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除BudgetData"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /budget_data/deleteBudgetData [delete]
 export const deleteBudgetDataByIds = (data) => {
     return service({
         url: "/budget_data/deleteBudgetDataByIds",
         method: 'delete',
         data
     })
 }

// @Tags BudgetData
// @Summary 更新BudgetData
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BudgetData true "更新BudgetData"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /budget_data/updateBudgetData [put]
 export const updateBudgetData = (data) => {
     return service({
         url: "/budget_data/updateBudgetData",
         method: 'put',
         data
     })
 }


// @Tags BudgetData
// @Summary 用id查询BudgetData
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.BudgetData true "用id查询BudgetData"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /budget_data/findBudgetData [get]
 export const findBudgetData = (params) => {
     return service({
         url: "/budget_data/findBudgetData",
         method: 'get',
         params
     })
 }


// @Tags BudgetData
// @Summary 分页获取BudgetData列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取BudgetData列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /budget_data/getBudgetDataList [get]
 export const getBudgetDataList = (params) => {
     return service({
         url: "/budget_data/getBudgetDataList",
         method: 'get',
         params
     })
 }