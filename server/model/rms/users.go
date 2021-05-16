// 自动生成模板Tag
package rms

import (
	"gin-vue-admin/global"
)

// 如果含有time.Time 请自行import time包
type Users struct {
      global.GVA_MODEL
      Name  string `json:"name" form:"name" gorm:"column:name;comment:角色名称;type:varchar(255);size:255;"`
      RoleID string `json:"roleid" form:"roleid" gorm:"column:role_id;comment:角色id;type:int"`
      ProductID string `json:"productid" form:"productid" gorm:"column:product_id;comment:产品id;type:int;"`
      Password string `json:"password" form:"password" gorm:"column:password;comment:密码;type:varchar(255);size:255;"`

      Macs           []Mac         `gorm:"many2many:users_mac;"`
      Products           []Product        `gorm:"many2many:users_product;"`
      Role   Role `json:"role" gorm:"foreignKey:id;References:role_id;AssociationForeignKey:RoleID;"`
}


func (Users) TableName() string {
  return "users"
}
