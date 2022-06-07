// 自动生成模板DepartmentAccountProduct
package future_data

import (
	"gin-vue-admin/global"
)

// 如果含有time.Time 请自行import time包
type Department struct {
	global.GVA_MODEL
	DepartmentName string `json:"departmentName" form:"departmentName" gorm:"column:department_name;comment:;type:varchar;"`
	DepartmentCode string `json:"departmentCode" form:"departmentCode" gorm:"column:department_code;comment:;type:varchar;"`
	CompanyCode    string `json:"companyCode" form:"companyCode" gorm:"column:company_code;comment:;type:varchar;"`
}

func (Department) TableName() string {
	return "department"
}
