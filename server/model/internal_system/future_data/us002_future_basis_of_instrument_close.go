// 自动生成模板TradeDetail
package future_data

import (
	"time"
)

// 如果含有time.Time 请自行import time包
type Us002FutureBasisOfInstrumentClose struct {
	ID                      uint    `json:"ID"  gorm:"column:id;comment:ID;"`
	Time                    string  `json:"time" form:"time" gorm:"column:time;comment:;"`
	DeliveryMonthInstrument string  `json:"deliveryMonthInstrument" form:"deliveryMonthInstrument" gorm:"column:delivery_month_instrument;comment:;type:varchar;"`
	BenchmarkInstrument     string  `json:"benchmarkInstrument" form:"benchmarkInstrument" gorm:"column:benchmark_instrument;comment:;type:varchar;"`
	AdjustBasis             float64 `json:"adjustBasis" form:"adjustBasis" gorm:"column:adjust_basis;comment:;type:double;"`
	BasisWithoutRisk        float64 `json:"basisWithoutRisk" form:"basisWithoutRisk" gorm:"column:basis_without_risk;comment:;type:double;"`
	ProductCode             string  `json:"productCode" form:"productCode" gorm:"column:product_code;comment:;type:varchar;"`
}

func (Us002FutureBasisOfInstrumentClose) TableName() string {
	return "master_data.us002_future_basis_of_instrument_close"
}

type PositionDeliveryMonthInstrument struct {
	Time               string               `json:"time" form:"time" gorm:"column:time;comment:;"`
	Instrument         string               `json:"instrument" form:"instrument" gorm:"column:instrument;comment:;type:varchar;"`
	Exchange_id        string               `json:"exchange_id" form:"exchange_id" gorm:"column:exchange_id;comment:;type:varchar;"`
	Product            string               `json:"product" form:"product" gorm:"column:product;comment:;type:varchar;"`
	Launch_date        time.Time            `json:"launch_date" form:"launch_date" gorm:"column:launch_date;comment:;"`
	Last_trade_date    time.Time            `json:"last_trade_date" form:"last_trade_date" gorm:"column:last_trade_date;comment:;"`
	Month              int64                `json:"month" form:"month" gorm:"column:month;comment:;type:date;size:0;"`
	Last_ddate         string               `json:"last_ddate" form:"last_ddate" gorm:"column:last_ddate;comment:;type:varchar;"`
	Is_throw           int64                `json:"is_throw" form:"is_throw" gorm:"column:is_throw;comment:;"`
	FarMonthInstrument []FarMonthInstrument `json:"farMonthInstrument" gorm:"foreignKey:Product;references:Product;comment:"`
	NoRiskValue        float64              `json:"noRiskValue" form:"noRiskValue" gorm:"column:noRiskValue;comment:;"`
}

type FarMonthInstrument struct {
	Instrument string `json:"instrument" form:"instrument" gorm:"column:instrument;comment:;type:varchar;"`
	Product    string `json:"product" form:"product" gorm:"column:product;comment:;type:varchar;"`
}
