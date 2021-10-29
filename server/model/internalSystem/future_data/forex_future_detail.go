// 自动生成模板ForexFutureDetail
package future_data

// 如果含有time.Time 请自行import time包
type ForexFutureDetail struct {
	ID             uint    `json:"id"  gorm:"column:id;comment:ID;"`
	Time           string  `json:"time" form:"time" gorm:"column:time;comment:;type:varchar;"`
	ProductName    string  `json:"productName" form:"productName" gorm:"column:product_name;comment:;type:varchar;"`
	AccountId      string  `json:"accountId" form:"accountId" gorm:"column:account_id;comment:;type:varchar;"`
	ProfitByFloat  float64 `json:"profitByFloat" form:"profitByFloat" gorm:"column:profit_by_float;comment:;"`
	ProfitByTrade  float64 `json:"profitByTrade" form:"profitByTrade" gorm:"column:profit_by_trade;comment:;"`
	TradeFee       float64 `json:"tradeFee" form:"tradeFee" gorm:"column:trade_fee;comment:;"`
	DepartmentName string  `json:"departmentName" form:"departmentName" gorm:"column:department_name;comment:;type:varchar;"`
}

func (ForexFutureDetail) TableName() string {
	return "future.forex_future_detail"
}
