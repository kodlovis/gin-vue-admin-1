// 自动生成模板Tag
package rms

import (
	"gin-vue-admin/global"
)

// 如果含有time.Time 请自行import time包
type Role struct {
	global.GVA_MODEL
	Name        string `json:"name" form:"name" gorm:"column:name;comment:角色名称;"`
	Description string `json:"description" form:"description" gorm:"column:description;comment:说明;"`
}

func (Role) TableName() string {
	return "role"
}
