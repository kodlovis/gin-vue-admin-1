// 自动生成模板ProductLeveragePrivateEquity
package rms

import (
	"gin-vue-admin/global"
)

// 如果含有time.Time 请自行import time包
type ProductLeveragePrivateEquity struct {
	global.GVA_MODEL
	ProductCode string  `json:"productCode" form:"productCode" gorm:"column:product_code;comment:;type:character varying;"`
	Leverage    float64 `json:"leverage" form:"leverage" gorm:"column:leverage;comment:;type:double precision;size:53;"`
}

func (ProductLeveragePrivateEquity) TableName() string {
	return "product_leverage_private_equity"
}
