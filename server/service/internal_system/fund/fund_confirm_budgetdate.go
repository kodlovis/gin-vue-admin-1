package fund

import (
	// "fmt"

	"gin-vue-admin/global"

	// "gin-vue-admin/model/internal_system/fund"
	"gin-vue-admin/model/internal_system/fund"
	model "gin-vue-admin/model/internal_system/fund"
	// request "gin-vue-admin/model/request/internal_system/fund"
)

func CreateconfirmBudget(confirm_date model.ConfirmBudget) (err error) {
	err = global.GVA_DB.Create(&confirm_date).Error
	return err
}
func FindConfirmBudget(searchInfo fund.SearchConfirm) (err error, budget []fund.BudgetData) {
	var budgets []fund.BudgetData
	db := global.GVA_DB
	// confirm_sql := "select * from fund.budget_date  where created_role =" +
	// 	"'" + searchInfo.Create_role + "'" +
	// 	" and department_code = " + "'" + searchInfo.Department_code + "'" + "order by created_at desc limit 1"
	// _ = db.Raw(confirm_sql).Scan(&confirmBudget).Error
	sql := "(select *,LAST_VALUE(in_value) OVER(PARTITION BY type_code ORDER BY created_at) from fund.budget_datas where create_role =" +
		"'" + searchInfo.Create_role + "'" +
		" and department_code = " + "'" + searchInfo.Department_code + "'" + " and to_char( created_at,'yyyy-mm-dd')::date = " + "'" + searchInfo.In_data + "')ORDER BY in_data"
	// sql := "select *,LAST_VALUE(in_value) OVER(PARTITION BY type_code ORDER BY created_at) from fund.budget_datas where create_role ='" + searchInfo.Create_role +
	// 	"'and department_code = '" + searchInfo.Department_code + "'and created_at > '" + confirmBudget.CreatedAt.Format("2006-01-02 15:04:05") + "'"
	err = db.Raw(sql).Scan(&budgets).Error
	return err, budgets
}

func GetConfirmBudget(limitData fund.LimitData) (err error, confirm []fund.ConfirmBudget) {
	var confirms []fund.ConfirmBudget
	// Department_code
	// db := global.GVA_DB.Model(&model.ConfirmSetting{})
	// _ = db.Where("confirm_date = ?", confirm_date).Find(&confirm_setting).Error

	// var is_confirm model.ConfirmBudget
	// confirm_db := global.GVA_DB.Model(&model.ConfirmBudget{})

	db := global.GVA_DB.Model(&model.ConfirmBudget{})
	err = db.Where("confirm_date = ?", limitData.ChoseDate).Find(&confirms).Error
	// sql := "SELECT * FROM  fund.budget_date WHERE id IN ( select max(id) from fund.budget_date  where to_char( confirm_date,'yyyy-mm-dd')::date = '" + limitData.ChoseDate + "' group by department_code,create_role,confirm_date)"
	// err = db.Raw(sql).Scan(&confirms).Error
	return err, confirms
}
