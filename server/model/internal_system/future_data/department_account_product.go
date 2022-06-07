// 自动生成模板DepartmentAccountProduct
package future_data

import (
	"gin-vue-admin/global"
	"time"
)

// 如果含有time.Time 请自行import time包
type DepartmentAccountProduct struct {
	global.GVA_MODEL
	DepartmentCode string    `json:"departmentCode" form:"departmentCode" gorm:"column:department_code;comment:;type:varchar;"`
	AccountId      string    `json:"accountId" form:"accountId" gorm:"column:account_id;comment:;type:varchar;"`
	EffectiveDate  time.Time `json:"effectiveDate" form:"effectiveDate" gorm:"column:effective_date;comment:;type:date;size:0;"`
	ExpirationDate time.Time `json:"expirationDate" form:"expirationDate" gorm:"column:expiration_date;comment:;type:date;size:0;"`
	ProductCode    string    `json:"productCode" form:"productCode" gorm:"column:product_code;comment:;type:varchar;"`
}

func (DepartmentAccountProduct) TableName() string {
	return "department_account_product"
}
