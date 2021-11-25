// 自动生成模板AccountPositionDaily
package position_reward

import (
	"gin-vue-admin/model/internal_system"
	"time"
)

// 如果含有time.Time 请自行import time包
type AccountPositionDaily struct {
	ID          uint                        `json:"id"  gorm:"column:id;comment:ID;"`
	TradingDate time.Time                   `json:"tradingDate" form:"tradingDate" gorm:"column:trading_date;comment:;"`
	BrokerId    string                      `json:"brokerId" form:"brokerId" gorm:"column:broker_id;comment:;type:character varying;"`
	AccountId   string                      `json:"accountId" form:"accountId" gorm:"column:account_id;comment:;type:character varying;"`
	Instrument  string                      `json:"instrument" form:"instrument" gorm:"column:instrument;comment:;type:character varying;"`
	Direction   int                         `json:"direction" form:"direction" gorm:"column:direction;comment:;"`
	HedgeFlag   int                         `json:"hedgeFlag" form:"hedgeFlag" gorm:"column:hedge_flag;comment:;"`
	Amount      int                         `json:"amount" form:"amount" gorm:"column:amount;comment:;"`
	AccountInfo internal_system.AccountInfo `json:"accountInfo" gorm:"foreignKey:AccountId;references:BrokerId;comment:"`
}

func (AccountPositionDaily) TableName() string {
	return "future.us001_account_position_daily"
}
