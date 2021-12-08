// 自动生成模板TradeDetail
package future_data

import "time"

// 如果含有time.Time 请自行import time包
type TradeDetail struct {
	ID            uint      `json:"id"  gorm:"column:id;comment:ID;"`
	Time          time.Time `json:"time" form:"time" gorm:"column:time;comment:;type:date;size:0;"`
	AccountId     string    `json:"accountId" form:"accountId" gorm:"column:account_id;comment:;type:varchar;"`
	Instrument    string    `json:"instrument" form:"instrument" gorm:"column:instrument;comment:;type:varchar;"`
	TradeTime     string    `json:"tradeTime" form:"tradeTime" gorm:"column:trade_time;comment:;type:varchar;"`
	TradeId       string    `json:"tradeId" form:"tradeId" gorm:"column:trade_id;comment:;type:varchar;"`
	Direction     string    `json:"direction" form:"direction" gorm:"column:direction;comment:;type:varchar;"`
	OffsetFlag    string    `json:"offsetFlag" form:"offsetFlag" gorm:"column:offset_flag;comment:;type:varchar;"`
	HedgeFlag     string    `json:"hedgeFlag" form:"hedgeFlag" gorm:"column:hedge_flag;comment:;type:varchar;"`
	TradePrice    float64   `json:"tradePrice" form:"tradePrice" gorm:"column:trade_price;comment:;type:double;"`
	Volume        int       `json:"volume" form:"volume" gorm:"column:volume;comment:;type:int;"`
	Amount        float64   `json:"amount" form:"amount" gorm:"column:amount;comment:;type:double;"`
	TradeFee      float64   `json:"tradeFee" form:"tradeFee" gorm:"column:trade_fee;comment:;type:double;"`
	ProfitByTrade float64   `json:"profitByTrade" form:"profitByTrade" gorm:"column:profit_by_trade;comment:;type:double;"`
}

func (TradeDetail) TableName() string {
	return "future.trade_detail"
}
