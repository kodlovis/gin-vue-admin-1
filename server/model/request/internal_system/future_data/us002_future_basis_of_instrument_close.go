package future_data

import "gin-vue-admin/model/internal_system/future_data"

type Us002FutureBasisOfInstrumentCloseSearch struct {
	future_data.Us002FutureBasisOfInstrumentClose
	PageInfo
}

type PositionDeliveryMonthInstrument struct {
	future_data.PositionDeliveryMonthInstrument
	PageInfo
	UserID uint `json:"userid" form:"userid"`
}

type NoRiskValue struct {
	Time               string  `json:"time" form:"time" gorm:"column:time;comment:;type:varchar;"`
	FarMonthInstrument string  `json:"farMonthInstrument" form:"farMonthInstrument" gorm:"column:ticker_f;comment:;type:varchar;"`
	Instrument         string  `json:"instrument" form:"instrument" gorm:"column:ticker_n;comment:;type:varchar;"`
	NoRiskValue        float64 `json:"noRiskValue" form:"noRiskValue" gorm:"column:safe_spread;comment:;type:varchar;"`
}

type BasisOInstrumentList struct {
	Us002FutureBasisOfInstrumentClose []future_data.Us002FutureBasisOfInstrumentClose `json:"item"`
}
