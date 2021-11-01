// 自动生成模板FutureRightsDaily
package future_data

// 如果含有time.Time 请自行import time包
type FutureRightsDaily struct {
	ID               uint    `json:"id"  gorm:"column:id;comment:ID;"`
	Time             string  `json:"time" form:"time" gorm:"column:time;comment:;type:varchar;"`
	CumulativeRights float64 `json:"cumulativeRights" form:"cumulativeRights" gorm:"column:cumulative_rights;comment:;type:double;"`
	Department       string  `json:"department" form:"department" gorm:"column:department;comment:;type:varchar;"`
	Product          string  `json:"product" form:"product" gorm:"column:product;comment:;type:varchar;"`
}

func (FutureRightsDaily) TableName() string {
	return "future.us002_future_rights_daily"
}
