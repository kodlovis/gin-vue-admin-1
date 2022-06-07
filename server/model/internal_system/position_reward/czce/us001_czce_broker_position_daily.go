// 自动生成模板ExchangeTradingfeeReward
package czce

import (
	"time"
)

// 如果含有time.Time 请自行import time包
type Us001CzceBrokerPositionDaily struct {
	ID                              uint      `json:"ID"  gorm:"column:id;comment:ID;"`
	Time                            time.Time `json:"time" form:"time" gorm:"column:time;comment:;type:date;size:0;"`
	ExchangeId                      string    `json:"exchangeId" form:"exchangeId" gorm:"column:exchange_id;comment:;type:varchar;"`
	BrokerId                        string    `json:"brokerId" form:"brokerId" gorm:"column:broker_id;comment:"`
	ProductCode                     string    `json:"productCode" form:"productCode" gorm:"column:product_code;comment:;type:varchar;"`
	CustomerDailyPosition           float64   `json:"customerDailyPosition" form:"customerDailyPosition" gorm:"column:customer_daily_position;comment:"`
	InstitutionalDailyPosition      float64   `json:"institutionalDailyPosition" form:"institutionalDailyPosition" gorm:"column:institutional_daily_position;comment:;type:double precision;size:53;"`
	SpecialLawCustomerDailyPosition float64   `json:"specialLawCustomerDailyPosition" form:"specialLawCustomerDailyPosition" gorm:"column:special_law_customer_daily_position;comment:;type:double precision;size:53;"`
	CorporateCustomerDailyPosition  float64   `json:"corporateCustomerDailyPosition" form:"corporateCustomerDailyPosition" gorm:"column:corporate_customer_daily_position;comment:;type:double precision;size:53;"`
	DailyAveragePosition            float64   `json:"dailyAveragePosition" form:"dailyAveragePosition" gorm:"column:daily_average_position;comment:;type:double precision;size:53;"`
	//LastYearLegalPersionPositionRatio float64   `json:"lastYearLegalPersionPositionRatio" form:"lastYearLegalPersionPositionRatio" gorm:"column:last_year_legal_persion_position_ratio;comment:;type:double precision;size:53;"`
	KeyContractMonthlyAveragePosition float64 `json:"keyContractMonthlyAveragePosition" form:"keyContractMonthlyAveragePosition" gorm:"column:key_contract_monthly_average_position;comment:;type:double precision;size:53;"`
	TheHandlingFeeDueTodayThisMonth   float64 `json:"theHandlingFeeDueTodayThisMonth" form:"theHandlingFeeDueTodayThisMonth" gorm:"column:the_handling_fee_due_today_this_month;comment:;type:double precision;size:53;"`
	CorporatePositionMultiplier       float64 `json:"corporatePositionMultiplier" form:"corporatePositionMultiplier" gorm:"column:corporate_position_multiplier;comment:;type:double precision;size:53;"`
}

func (Us001CzceBrokerPositionDaily) TableName() string {
	return "future.us001_czce_broker_position_daily"
}
