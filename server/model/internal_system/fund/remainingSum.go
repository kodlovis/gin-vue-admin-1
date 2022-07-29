// 自动生成模板RemainingSum
package fund

import (
	"gin-vue-admin/global"
)

// 如果含有time.Time 请自行import time包
type RemainingSum struct {
	global.GVA_MODEL
	CreateUser string `json:"createUser" form:"createUser" gorm:"column:create_user;comment:;type:character varying;"`
	Type       string `json:"type" form:"type" gorm:"column:type;comment:;type:character varying;"`
	InValue    int    `json:"inValue" form:"inValue" gorm:"column:in_value;comment:;type:bigint;size:64;"`
	UpDate     string `json:"upDate" form:"upDate" gorm:"column:up_date;comment:;type:datetime;"`
	Company    string `json:"company" form:"company" gorm:"column:company;comment:;type:character varying;"`
}

type SearchRemainingSum struct {
	Type    string `json:"type" form:"type" gorm:"column:type;comment:;type:character varying;"`
	UpDate  string `json:"upDate" form:"upDate" gorm:"column:up_date;comment:;type:datetime;"`
	Company string `json:"company" form:"company" gorm:"column:company;comment:;type:character varying;"`
}

func (RemainingSum) TableName() string {
	return "fund.remaining_sum"
}
