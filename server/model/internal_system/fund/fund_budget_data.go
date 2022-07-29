// 自动生成模板BudgetData
package fund

import (
	"gin-vue-admin/global"
)

// 如果含有time.Time 请自行import time包
type BudgetData struct {
	global.GVA_MODEL
	In_data         string `json:"in_data" form:"in_data" gorm:"column:in_data;comment:录入日期;type:datetime;"`
	Type_code       string `json:"type_code" form:"type_code" gorm:"column:type_code;comment:录入类型;type:varchar;"`
	In_value        int    `json:"in_value" form:"in_value" gorm:"column:in_value;comment:录入数据;type:bigint;"`
	Comment         string `json:"comment" form:"comment" gorm:"column:comment;comment:数据备注;type:varchar;"`
	Currency        string `json:"currency" form:"currency" gorm:"column:currency;comment:币种;type:varchar;"`
	Department_code string `json:"department_code" form:"department_code" gorm:"column:department_code;comment:部门编号;type:varchar;"`
	Create_user     string `json:"create_user" form:"create_user" gorm:"column:create_user;comment:录入人员;type:varchar;"`
	Create_role     string `json:"create_role" form:"create_role" gorm:"column:create_role;comment:录入角色;type:varchar;"`
	Up_date         string `json:"up_date" form:"up_date" gorm:"column:up_date;comment:填报日期;type:datetime;"`
}

type SearchDay struct {
	Department_code string `json:"department_code"`
	Create_role     string `json:"create_role"`
	In_data         string `json:"in_data"`
	Create_date     string `json:"create_date"`
	Up_date         string `json:"up_date"`
}

func (BudgetData) TableName() string {
	return "fund.budget_datas"
}
