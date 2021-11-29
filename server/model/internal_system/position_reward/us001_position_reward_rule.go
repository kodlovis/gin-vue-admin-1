// 自动生成模板PositionRewardRule
package position_reward

import "gin-vue-admin/model/internal_system"

// 如果含有time.Time 请自行import time包
type PositionRewardRule struct {
	ID             uint                                `json:"id"  gorm:"column:id;comment:ID;"`
	ExchangeId     string                              `json:"exchangeId" form:"exchangeId" gorm:"column:exchange_id;comment:"`
	BrokerId       string                              `json:"brokerId" form:"brokerId" gorm:"column:broker_id;comment:"`
	ProductCode    string                              `json:"productCode" form:"productCode" gorm:"column:product_code;comment:"`
	BasePosition   float64                             `json:"basePosition" form:"basePosition" gorm:"column:base_position;comment:"`
	RewardType     int                                 `json:"rewardType" form:"rewardType" gorm:"column:reward_type;comment:"`
	RewardAmount   float64                             `json:"rewardAmount" form:"rewardAmount" gorm:"column:reward_amount;comment:"`
	EffectiveDate  string                              `json:"effectiveDate" form:"effectiveDate" gorm:"column:effective_date;comment:;type:varchar;"`
	ExpirationDate string                              `json:"expirationDate" form:"expirationDate" gorm:"column:expiration_date;comment:;type:varchar;"`
	ProductInfo    internal_system.ExchangeProductInfo `json:"productInfo" gorm:"foreignKey:ProductCode;references:ProductCode;comment:"`
	AccountInfo    internal_system.AccountInfo         `json:"accountInfo" gorm:"foreignKey:AccountId;references:BrokerId;comment:"`
}

func (PositionRewardRule) TableName() string {
	return "future.us001_position_reward_rule"
}
