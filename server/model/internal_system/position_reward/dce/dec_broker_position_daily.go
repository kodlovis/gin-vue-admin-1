// 自动生成模板ExchangeTradingfeeReward
package dce

import (
	"time"
)

// 如果含有time.Time 请自行import time包
type Us001DceBrokerPositionDaily struct {
	ID                         uint      `json:"ID"  gorm:"column:id;comment:ID;"`
	Time                       time.Time `json:"time" form:"time" gorm:"column:time;comment:;type:date;size:0;"`
	BrokerId                   string    `json:"brokerId" form:"brokerId" gorm:"column:broker_id;comment:;type:varchar;"`
	ProductCode                string    `json:"productCode" form:"productCode" gorm:"column:product_code;comment:;type:varchar;"`
	CorporateMarketPosition    float64   `json:"corporateMarketPosition" form:"corporateMarketPosition" gorm:"column:corporate_market_position;comment:;type:double precision;size:53;"`
	CorporateNonMarketPosition float64   `json:"corporateNonMarketPosition" form:"corporateNonMarketPosition" gorm:"column:corporate_non_market_position;comment:;type:double precision;size:53;"`
	OtherMarketPosition        float64   `json:"otherMarketPosition" form:"otherMarketPosition" gorm:"column:other_market_position;comment:;type:double precision;size:53;"`
	OtherNonMarketPosition     float64   `json:"otherNonMarketPosition" form:"otherNonMarketPosition" gorm:"column:other_non_market_position;comment:;type:double precision;size:53;"`
	PayFee                     float64   `json:"payFee" form:"payFee" gorm:"column:pay_fee;comment:;type:double precision;size:53;"`
}

func (Us001DceBrokerPositionDaily) TableName() string {
	return "master_data.us001_dce_broker_position_daily"
}
