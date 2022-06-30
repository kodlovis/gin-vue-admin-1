package future_data

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model/internal_system/public"
)

// 如果含有time.Time 请自行import time包
type Us002ForeignExchangeProfit struct {
	global.GVA_MODEL
	Time           string               `json:"time" form:"time" gorm:"column:time;comment:;type:varchar;"`
	Currency       string               `json:"currency" form:"currency" gorm:"column:currency;comment:;type:varchar;"`
	AmountOfProfit float64              `json:"amountOfProfit" form:"amountOfProfit" gorm:"column:amount_of_profit;comment:;type:double;"`
	DepartmentCode string               `json:"departmentCode" form:"departmentCode" gorm:"column:department_code;comment:;type:varchar;"`
	Department     Department           `json:"department" gorm:"foreignKey:DepartmentCode;references:DepartmentCode;comment:"`
	Currencys      public.Us100Currency `json:"currencys" gorm:"foreignKey:Currency;references:CurrencyCode;comment:"`
}

func (Us002ForeignExchangeProfit) TableName() string {
	return "master_data.us002_foreign_exchange_profit"
}
