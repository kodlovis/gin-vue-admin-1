// 自动生成模板FutureAccount
package future_data

// 如果含有time.Time 请自行import time包
type FutureAccount struct {
	ID       uint   `json:"id"  gorm:"column:id;comment:ID;"`
	Username string `json:"username" form:"username" gorm:"column:username;comment:;type:character varying;"`
	Password string `json:""  gorm:"comment:;"`
	Status   int    `json:"status" form:"status" gorm:"column:status;comment:;type:bigint;size:64;"`
}

func (FutureAccount) TableName() string {
	return "future.us002_future_account"
}
