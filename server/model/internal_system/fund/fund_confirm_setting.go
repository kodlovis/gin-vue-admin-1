package fund

import (
	"gin-vue-admin/global"
)

// 如果含有time.Time 请自行import time包
type ConfirmSetting struct {
	global.GVA_MODEL
	Is_ok        int    `json:"is_ok" form:"is_ok" gorm:"column:is_ok;comment:是否允许填报;type:bigint;"`
	Confirm_date string `json:"confirm_date" form:"confirm_date" gorm:"column:confirm_date;comment:填报日期;type:datetime;"`
}

type SearchConfirmSetting struct {
	Confirm_date string `json:"confirm_date"`
}

func (ConfirmSetting) TableName() string {
	return "fund.confirm_setting"
}
