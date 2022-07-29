package fund

import (
	"fmt"
	"gin-vue-admin/global"
	"gin-vue-admin/model/internal_system/fund"
	model "gin-vue-admin/model/internal_system/fund"
	request "gin-vue-admin/model/request/internal_system/fund"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateBudgetData
//@description: 创建BudgetData记录
//@param: budget_data model.BudgetData
//@return: err error

func CreateBudgetData(budget_data model.BudgetData) (err error) {
	err = global.GVA_DB.Create(&budget_data).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteBudgetData
//@description: 删除BudgetData记录
//@param: budget_data model.BudgetData
//@return: err error

func DeleteBudgetData(budget_data model.BudgetData) (err error) {
	err = global.GVA_DB.Delete(&budget_data).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteBudgetDataByIds
//@description: 批量删除BudgetData记录
//@param: ids request.IdsReq
//@return: err error

func DeleteBudgetDataByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.BudgetData{}, "id in ?", ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateBudgetData
//@description: 更新BudgetData记录
//@param: budget_data *model.BudgetData
//@return: err error

func UpdateBudgetData(budget_data model.BudgetData) (err error) {
	err = global.GVA_DB.Save(&budget_data).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetBudgetData
//@description: 根据id获取BudgetData记录
//@param: id uint
//@return: err error, budget_data model.BudgetData

func GetBudgetData(id uint) (err error, budget_data model.BudgetData) {
	err = global.GVA_DB.Where("id = ?", id).First(&budget_data).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetBudgetDataInfoList
//@description: 分页获取BudgetData记录
//@param: info request.BudgetDataSearch
//@return: err error, list interface{}, total int64

func GetBudgetDataInfoList(info request.BudgetDataSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&model.BudgetData{})
	var budget_datas []model.BudgetData
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&budget_datas).Error
	return err, budget_datas, total
}

func SaveBudgetData(budgets []fund.BudgetData) (err error) {

	for _, item := range budgets {
		err = CreateBudgetData(item)
	}
	fmt.Println(err)
	return err
}

func FindDayData(searchInfo fund.SearchDay) (err error, budget []fund.BudgetData) {
	var budgets []fund.BudgetData
	//data := searchInfo.DataDate + " 00:00:00.000"
	db := global.GVA_DB
	// 	select  name,sex,age,first_value(name) over (partition by sex order by name)
	//    from user

	sql := "SELECT distinct in_data,type_code,first_value(in_value) over (partition by type_code order by created_at desc) as in_value," +
		"first_value(comment) over (partition by type_code order by created_at desc) as comment,first_value(create_role) over (partition by type_code order by created_at desc) as create_role FROM fund.budget_datas where create_role =" +
		"'" + searchInfo.Create_role + "'" +
		" and department_code = " + "'" + searchInfo.Department_code + "'" + " and in_data = " + "'" + searchInfo.In_data + "'and to_char( up_date,'yyyy-mm-dd')::date ='" + searchInfo.Up_date + "'"
	err = db.Raw(sql).Scan(&budgets).Error
	return err, budgets
}

//获取每日最新录入收支信息
func GetDayData(searchInfo fund.SearchDay) (err error, budget []fund.BudgetData) {
	var budgets []fund.BudgetData
	//data := searchInfo.DataDate + " 00:00:00.000"
	db := global.GVA_DB
	// 	select  name,sex,age,first_value(name) over (partition by sex order by name)
	//    from user
	var sql string
	if searchInfo.Create_role != "" {
		sql = "with today_all_data as ( select * from  fund.budget_datas where department_code = '" + searchInfo.Department_code + "' and  to_char( up_date,'yyyy-mm-dd')::date = '" + searchInfo.Create_date +
			"' and create_role ='" +
			searchInfo.Create_role + "' ) select * from today_all_data where id in (select distinct first_value(id) over (partition by type_code,in_data order by created_at desc) from today_all_data)"
	} else {
		sql = "with today_all_data as ( select * from  fund.budget_datas where department_code = '" + searchInfo.Department_code + "' and  to_char( up_date,'yyyy-mm-dd')::date = '" + searchInfo.Create_date +
			"'  ) select * from today_all_data where id in (select distinct first_value(id) over (partition by type_code,in_data order by created_at desc) from today_all_data)"
	}

	err = db.Raw(sql).Scan(&budgets).Error
	return err, budgets
}

func GetReport(searchInfo fund.SearchDay) (err error, budget []fund.BudgetData) {
	var budgets []fund.BudgetData
	//data := searchInfo.DataDate + " 00:00:00.000"
	db := global.GVA_DB
	// 	select  name,sex,age,first_value(name) over (partition by sex order by name)
	//    from user
	var sql string

	department_codes := [11]string{"有色事业部", "黑色事业部", "农产事业部", "原料事业部", "钢材事业部", "总经理二室", "总经理一室", "杉贸本部", "交易部", "智维钢材", "智维本部"}
	for _, Department_code := range department_codes {
		var budget_item []fund.BudgetData
		sql = "with today_all_data as ( select * from  fund.budget_datas where department_code = '" + Department_code + "' and  to_char( up_date,'yyyy-mm-dd')::date = '" + searchInfo.Create_date +
			"'  ) select * from today_all_data where id in (select distinct first_value(id) over (partition by type_code,in_data order by created_at desc) from today_all_data)"

		err = db.Raw(sql).Scan(&budget_item).Error
		budgets = append(budgets, budget_item...)
	}

	return err, budgets
}
