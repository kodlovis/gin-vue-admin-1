// 自动生成模板FuturePositionSizeDetail
package scale_rights

import (
	"gin-vue-admin/global"
	"time"
)

// 如果含有time.Time 请自行import time包
type FuturePositionSizeDetail struct {
	global.GVA_MODEL
	Time            time.Time `json:"time" form:"time" gorm:"column:time;comment:;type:date;size:0;"`
	Department      string    `json:"department" form:"department" gorm:"column:department;comment:;type:varchar;"`
	TradeType       string    `json:"tradeType" form:"tradeType" gorm:"column:trade_type;comment:;type:varchar;"`
	RiskType        string    `json:"riskType" form:"riskType" gorm:"column:risk_type;comment:;type:varchar;"`
	LongInstrument  string    `json:"longInstrument" form:"longInstrument" gorm:"column:long_instrument;comment:;type:character varying;"`
	LongBasis       float64   `json:"longBasis" form:"longBasis" gorm:"column:long_basis;comment:;type:double precision;size:53;"`
	LongMarket      float64   `json:"longMarket" form:"longMarket" gorm:"column:long_market;comment:;type:double precision;size:53;"`
	Price           float64   `json:"price" form:"price" gorm:"column:price;comment:;type:double precision;size:53;"`
	PriceMarket     float64   `json:"priceMarket" form:"priceMarket" gorm:"column:price_market;comment:;type:double precision;size:53;"`
	LongAccount     string    `json:"longAccount" form:"longAccount" gorm:"column:long_account;comment:;type:character varying;"`
	LongPosition    float64   `json:"longPosition" form:"longPosition" gorm:"column:long_position;comment:;type:double precision;size:53;"`
	LongComment     string    `json:"longComment" form:"longComment" gorm:"column:long_comment;comment:;type:character varying;"`
	ShortInstrument string    `json:"shortInstrument" form:"shortInstrument" gorm:"column:short_instrument;comment:;type:character varying;"`
	ShortBasis      float64   `json:"shortBasis" form:"shortBasis" gorm:"column:short_basis;comment:;type:double precision;size:53;"`
	ShortMarket     float64   `json:"shortMarket" form:"shortMarket" gorm:"column:short_market;comment:;type:double precision;size:53;"`
	ShortAccount    string    `json:"shortAccount" form:"shortAccount" gorm:"column:short_account;comment:;type:character varying;"`
	ShortPosition   float64   `json:"shortPosition" form:"shortPosition" gorm:"column:short_position;comment:;type:double precision;size:53;"`
	ShortComment    string    `json:"shortComment" form:"shortComment" gorm:"column:short_comment;comment:;type:character varying;"`
}

func (FuturePositionSizeDetail) TableName() string {
	return "future.us002_future_position_size_detail"
}
