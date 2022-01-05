// 自动生成模板ExchangeTradingfeeReward
package position_reward

import (
	"time"
)

// 如果含有time.Time 请自行import time包
type ExchangeTradingfeeReward struct {
	ID            uint      `json:"id"  gorm:"column:id;comment:ID;"`
	EffectiveDate time.Time `json:"effectiveDate" form:"effectiveDate" gorm:"column:effective_date;comment:;type:date;size:0;"`
	ExchangeId    string    `json:"exchangeId" form:"exchangeId" gorm:"column:exchange_id;comment:;type:character varying;"`
	ProductCode   string    `json:"productCode" form:"productCode" gorm:"column:product_code;comment:;type:character varying;"`
	RewardRate    float64   `json:"rewardRate" form:"rewardRate" gorm:"column:reward_rate;comment:;type:double precision;size:53;"`
	Description   string    `json:"description" form:"description" gorm:"column:description;comment:;type:character varying;"`
}

func (ExchangeTradingfeeReward) TableName() string {
	return "future.exchange_tradingfee_reward"
}
