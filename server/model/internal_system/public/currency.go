// 自动生成模板Us100Currency
package public

// 如果含有time.Time 请自行import time包
type Us100Currency struct {
	ID           uint   `json:"ID"  gorm:"column:id;comment:ID;"`
	CurrencyCode string `json:"currencyCode" form:"currencyCode" gorm:"column:currency_code;comment:;type:varchar;"`
	CurrencyName string `json:"currencyName" form:"currencyName" gorm:"column:currency_name;comment:;type:varchar;"`
}

func (Us100Currency) TableName() string {
	return "future.us100_currency"
}
