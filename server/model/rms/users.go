// 自动生成模板Tag
package rms

import (
	"gin-vue-admin/global"
)

// 如果含有time.Time 请自行import time包
type Users struct {
	global.GVA_MODEL
	Name   string `json:"name" form:"name" gorm:"column:name;comment:角色名称;"`
	RoleID uint   `json:"roleid" form:"roleid" gorm:"column:role_id;comment:角色id"`
	//ProductID string `json:"productid" form:"productid" gorm:"column:product_id;comment:产品id;type:int;"`
	Password string `json:"password" form:"password" gorm:"column:password;comment:密码;"`

	Macs     []Mac     `gorm:"many2many:users_mac;"`
	Products []Product `gorm:"many2many:users_product;"`
	//Role     Role      `json:"role" gorm:"foreignKey:id;AssociationForeignKey:RoleID;"`
	Role Role `json:"role" gorm:"AssociationForeignKey:RoleID;"`
}

func (Users) TableName() string {
	return "users"
}
