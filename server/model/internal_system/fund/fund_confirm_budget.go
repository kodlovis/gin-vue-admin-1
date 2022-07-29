// 自动生成模板BudgetData
package fund

import (
	"gin-vue-admin/global"
)

// 如果含有time.Time 请自行import time包
type ConfirmBudget struct {
	global.GVA_MODEL
	Department_code string `json:"department_code" form:"department_code" gorm:"column:department_code;comment:部门编号;type:varchar;"`
	Create_user     string `json:"create_user" form:"create_user" gorm:"column:create_user;comment:录入人员;type:varchar;"`
	Create_role     string `json:"create_role" form:"create_role" gorm:"column:create_role;comment:录入角色;type:varchar;"`
	Confirm_date    string `json:"confirm_date" form:"confirm_date" gorm:"column:confirm_date;comment:确认日期;type:datetime;"`
	Is_confirm      string `json:"is_confirm" form:"is_confirm" gorm:"column:is_confirm;comment:是否确认;type:varchar;"`
}

type SearchConfirm struct {
	Department_code string `json:"department_code"`
	Create_role     string `json:"create_role"`
	In_data         string `json:"in_data"`
}

type LimitData struct {
	ChoseDate string `json:"chose_date"`
}

func (ConfirmBudget) TableName() string {
	return "fund.budget_date"
}
