// 自动生成模板FutureDepartmentScaleDetail
package scale_rights

// 如果含有time.Time 请自行import time包
type FutureDepartmentScaleDetail struct {
	ID             uint    `json:"id"  gorm:"column:id;comment:ID;"`
	Time           string  `json:"time" form:"time" gorm:"column:time;comment:;type:varchar;"`
	DepartmentName string  `json:"departmentName" form:"departmentName" gorm:"column:department_name;comment:;type:varchar;"`
	Type           string  `json:"type" form:"type" gorm:"column:type;comment:;type:varchar;"`
	ExistingScale  float64 `json:"existingScale" form:"existingScale" gorm:"column:existing_scale;comment:;type:double;"`
}

func (FutureDepartmentScaleDetail) TableName() string {
	return "future.us002_future_department_scale_detail"
}
