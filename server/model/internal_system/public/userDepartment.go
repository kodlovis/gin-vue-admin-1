// 自动生成模板UserDepartment
package public

import (
	"gin-vue-admin/global"
	"time"
)

// 如果含有time.Time 请自行import time包
type UserDepartment struct {
	global.GVA_MODEL
	UserId         int        `json:"userId" form:"userId" gorm:"column:user_id;comment:;type:bigint;size:64;"`
	DepartmentCode string     `json:"departmentCode" form:"departmentCode" gorm:"column:department_code;comment:;type:varchar;"`
	EffectiveDate  time.Time  `json:"effectiveDate" form:"effectiveDate" gorm:"column:effective_date;comment:;type:date;size:0;"`
	ExpirationDate time.Time  `json:"expirationDate" form:"expirationDate" gorm:"column:expiration_date;comment:;type:date;size:0;"`
	DepartmentInfo Department `json:"departmentInfo" gorm:"foreignKey:DepartmentCode;references:DepartmentCode;comment:"`
}

func (UserDepartment) TableName() string {
	return "master_data.us010_user_department"
}
