// 自动生成模板FutureDeliveryDetailByInput
package future_data

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model/internal_system"
	"time"
)

// 如果含有time.Time 请自行import time包
type FutureDeliveryDetailByInput struct {
	global.GVA_MODEL
	DepartmentName   string                      `json:"departmentName" form:"departmentName" gorm:"column:department_name;comment:;type:varchar;"`
	ProductName      string                      `json:"productName" form:"productName" gorm:"column:product_name;comment:;type:varchar;"`
	AccountId        string                      `json:"accountId" form:"accountId" gorm:"column:account_id;comment:;type:varchar;"`
	ProfitOfPosition float64                     `json:"profitOfPosition" form:"profitOfPosition" gorm:"column:profit_of_position;comment:;type:double precision;size:53;"`
	ProfitOfTrade    float64                     `json:"profitOfTrade" form:"profitOfTrade" gorm:"column:profit_of_trade;comment:;type:double precision;size:53;"`
	TradeFee         float64                     `json:"tradeFee" form:"tradeFee" gorm:"column:trade_fee;comment:;type:double precision;size:53;"`
	Time             time.Time                   `json:"time" form:"time" gorm:"column:time;comment:;type:date;size:0;"`
	AccountInfo      internal_system.AccountInfo `json:"accountInfo" gorm:"foreignKey:AccountId;references:AccountId;comment:"`
}

func (FutureDeliveryDetailByInput) TableName() string {
	return "future.us005_future_delivery_detail_by_input"
}
