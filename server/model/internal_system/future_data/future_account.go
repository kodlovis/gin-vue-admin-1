// 自动生成模板FutureAccount
package future_data

// 如果含有time.Time 请自行import time包
type FutureAccount struct {
	ID          uint   `json:"id"  gorm:"column:id;comment:ID;"`
	Username    string `json:"username" form:"username" gorm:"column:username;comment:;type:character varying;"`
	Password    string `json:""  gorm:"comment:;"`
	Status      int    `json:"status" form:"status" gorm:"column:status;comment:;type:bigint;size:64;"`
	Proxy       string `json:"proxy" form:"proxy" gorm:"column:proxy;comment:;type:varchar;"`
	Description string `json:"description" form:"description" gorm:"column:description;comment:;type:varchar;"`
	Group       string `json:"group" form:"group" gorm:"column:group;comment:;type:varchar;"`
	//UpdateAt    time.Time `json:"updateAt" form:"updateAt" gorm:"column:update_at;comment:;type:date;size:0;"`
	IsUseProxy int `json:"isUseProxy" form:"isUseProxy" gorm:"column:is_use_proxy;comment:;type:bigint;size:64;"`
}

func (FutureAccount) TableName() string {
	return "future.us002_future_account"
}
