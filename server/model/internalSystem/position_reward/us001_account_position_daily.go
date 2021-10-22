// 自动生成模板AccountPositionDaily
package position_reward

import (
	"gin-vue-admin/global"
)

// 如果含有time.Time 请自行import time包
type AccountPositionDaily struct {
	global.GVA_MODEL
	TradingDate string  `json:"tradingDate" form:"tradingDate" gorm:"column:trading_date;comment:"`
	BrokerId    string  `json:"brokerId" form:"brokerId" gorm:"column:broker_id;comment:;type:character varying;"`
	AccountId   string  `json:"accountId" form:"accountId" gorm:"column:account_id;comment:;type:character varying;"`
	Instrument  string  `json:"instrument" form:"instrument" gorm:"column:instrument;comment:;type:character varying;"`
	Direction   float64 `json:"direction" form:"direction" gorm:"column:direction;comment:;type:integer;"`
	HedgeFlag   float64 `json:"hedgeFlag" form:"hedgeFlag" gorm:"column:hedge_flag;comment:;type:integer;"`
	Amount      float64 `json:"amount" form:"amount" gorm:"column:amount;comment:;type:integer;"`
}

func (AccountPositionDaily) TableName() string {
	return "us001_account_position_daily"
}
