// 自动生成模板Us004FutureInventoryDaily
package future_inventory

import (
	"time"
)

// 如果含有time.Time 请自行import time包
type Us004FutureInventoryDaily struct {
	ID          uint      `json:"ID"  gorm:"column:id;comment:ID;"`
	Time        time.Time `json:"time" form:"time" gorm:"column:time;comment:;type:date;size:0;"`
	ProductCode string    `json:"productCode" form:"productCode" gorm:"column:product_code;comment:;type:varchar;"`
	Volume      string    `json:"volume" form:"volume" gorm:"column:volume;comment:;type:double precision;size:53;"`
	ExchangeId  string    `json:"exchangeId" form:"exchangeId" gorm:"column:exchange_id;comment:;type:varchar;"`
	Comment     string    `json:"comment" form:"comment" gorm:"column:comment;comment:;type:varchar;"`
	Unit        string    `json:"unit" form:"unit" gorm:"column:unit;comment:;type:varchar;"`
	Type        string    `json:"type" form:"type" gorm:"column:type;comment:;type:varchar;"`
}

func (Us004FutureInventoryDaily) TableName() string {
	return "future.us004_future_inventory_daily"
}
