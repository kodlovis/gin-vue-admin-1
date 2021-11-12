// 自动生成模板FutureDeparmentRights
package scale_rights

// 如果含有time.Time 请自行import time包
type FutureDeparmentRights struct {
	ID         uint   `json:"id"  gorm:"column:id;comment:ID;"`
	Time       string `json:"time" form:"time" gorm:"column:time;comment:;type:varchar;"`
	Department string `json:"department" form:"department" gorm:"column:department;comment:;type:varchar;"`
	Type       string `json:"type" form:"type" gorm:"column:type;comment:;type:varchar;"`
	Profit     string `json:"profit" form:"profit" gorm:"column:profit;comment:;type:double precision;size:53;"`
}

func (FutureDeparmentRights) TableName() string {
	return "future.us002_future_deparment_rights"
}
