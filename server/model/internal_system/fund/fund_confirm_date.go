// 自动生成模板FundConfirm
package fund

import (
	"gin-vue-admin/global"
)

// 如果含有time.Time 请自行import time包
type FundConfirm struct {
	global.GVA_MODEL
	End_date   string `json:"end_date" form:"end_date" gorm:"column:end_date;comment:终止时间;type:datetime;"`
	Start_date string `json:"start_date" form:"start_date" gorm:"column:start_date;comment:终止时间;type:datetime;"`
}

type SearchFundConfirm struct {
	Start_date string `json:"start_date" form:"start_date" gorm:"column:start_date;comment:终止时间;type:datetime;"`
}

func (FundConfirm) TableName() string {
	return "fund.confirm_date"
}
