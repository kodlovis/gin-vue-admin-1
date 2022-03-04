// 自动生成模板Us004FutureStockLme
package future_inventory

import "time"

// 如果含有time.Time 请自行import time包
type Us004FutureStockLme struct {
	ID               uint      `json:"ID"  gorm:"column:id;comment:ID;"`
	MarketDay        time.Time `json:"marketDay" form:"marketDay" gorm:"column:market_day;comment:;type:date;size:0;"`
	FuturesVariety   string    `json:"futuresVariety" form:"futuresVariety" gorm:"column:futures_variety;comment:;type:varchar;"`
	WarehouseReceipt float64   `json:"warehouseReceipt" form:"warehouseReceipt" gorm:"column:warehouse_receipt;comment:;type:double;"`
	CancelStock      float64   `json:"cancelStock" form:"cancelStock" gorm:"column:cancel_stock;comment:;type:double;"`
}

func (Us004FutureStockLme) TableName() string {
	return "future.us004_future_stock_lme"
}
