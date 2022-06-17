// 自动生成模板Tag
package rms

import (
	"gin-vue-admin/global"
)

// 如果含有time.Time 请自行import time包
type RiskSetting struct {
      global.GVA_MODEL
      Name  string `json:"name" form:"name" gorm:"column:name;comment:名称;type:varchar(255);size:255;"`
      RiskSettingPoint  string `json:"code" form:"code" gorm:"column:risk_point;comment:风险点;type:varchar(255);size:900;"`
      BrokersValue  float64 `json:"brokersValue" form:"brokersValue" gorm:"column:brokers_value;comment:券商阈值;type:float;"`
      AdminValue  float64 `json:"adminValue" form:"adminValue" gorm:"column:admin_value;comment:风控阈值;type:float;"`
      Description string `json:"description" form:"description" gorm:"column:description;comment:说明;type:varchar(255);size:900;"`
}


func (RiskSetting) TableName() string {
  return "risk_settings"
}
