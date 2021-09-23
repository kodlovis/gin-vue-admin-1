// 自动生成模板Tag
package rms

import (
	"gin-vue-admin/global"
)

// 如果含有time.Time 请自行import time包
type Mac struct {
	global.GVA_MODEL
	Name string `json:"name" form:"name" gorm:"column:mac_address;comment:mac地址;"`
	//Users  []Users `gorm:"many2many:users_mac;"`
}

func (Mac) TableName() string {
	return "mac"
}
