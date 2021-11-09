// position_reward
package position_reward

// 如果含有time.Time 请自行import time包
type BrokerPositionDaily struct {
	ID                   uint    `json:"id"  gorm:"column:id;comment:ID;"`
	TradingDate          string  `json:"tradingDate" form:"tradingDate" gorm:"column:trading_date;comment:;type:varchar;"`
	BrokerId             string  `json:"brokerId" form:"brokerId" gorm:"column:broker_id;comment:"`
	ProductCode          string  `json:"productCode" form:"productCode" gorm:"column:product_code;comment:"`
	TotalTradingFee      float64 `json:"totalTradingFee" form:"totalTradingFee" gorm:"column:total_trading_fee;comment:"`
	SpecialTradingFee    float64 `json:"specialTradingFee" form:"specialTradingFee" gorm:"column:special_trading_fee;comment:"`
	AverageDailyPosition float64 `json:"averageDailyPosition" form:"averageDailyPosition" gorm:"column:average_daily_position;comment:"`
	MaximumToOpen        float64 `json:"maximumToOpen" form:"maximumToOpen" gorm:"column:maximum_to_open;comment:"`
	CurrentPosition      float64 `json:"currentPosition" form:"currentPosition" gorm:"column:current_position;comment:"`
}

func (BrokerPositionDaily) TableName() string {
	return "future.us001_broker_position_daily"
}
