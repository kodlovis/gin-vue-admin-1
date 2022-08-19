// 自动生成模板Department
package public

// 如果含有time.Time 请自行import time包
type Department struct {
	ID             uint   `json:"ID"  gorm:"column:id;comment:ID;"`
	DepartmentName string `json:"departmentName" form:"departmentName" gorm:"column:department_name;comment:;type:character varying;"`
	DepartmentCode string `json:"departmentCode" form:"departmentCode" gorm:"column:department_code;comment:;type:character varying;"`
	CompanyCode    string `json:"companyCode" form:"companyCode" gorm:"column:company_code;comment:;type:character varying;"`
}

func (Department) TableName() string {
	return "future.department"
}
