// 自动生成模板Tag
package rms

import (
	"gin-vue-admin/global"
)

// 如果含有time.Time 请自行import time包
type Product struct {
      global.GVA_MODEL
      Name  string `json:"name" form:"name" gorm:"column:name;comment:产品名称;type:varchar(255);size:255;"`
      AccountID  string `json:"accountid" form:"accountid" gorm:"column:account_id;comment:账号;type:varchar(255);size:255;"`
      CounterInfo  string `json:"counterInfo" form:"counterInfo" gorm:"column:counter_info;comment:柜台信息;type:varchar(255);size:255;"`
      Password  string `json:"password" form:"password" gorm:"column:password;comment:密码;type:varchar(255);size:255;"`
      NodeInfo  string `json:"nodeInfo" form:"nodeInfo" gorm:"column:node_info;comment:信息;type:varchar(255);size:20000;"`
      Comment  string `json:"comment" form:"comment" gorm:"column:comment;comment:备注;type:varchar(255);size:255;"`
      Description  string `json:"description" form:"description" gorm:"column:description;comment:描述;type:varchar(255);size:255;"`
      //Users  []Users `gorm:"many2many:users_product;"`
}


func (Product) TableName() string {
  return "product"
}
