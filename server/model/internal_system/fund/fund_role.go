package fund

// 如果含有time.Time 请自行import time包
type FundUser struct {
	Uid         int    `json:"uid" gorm:"comment:用户id"`
	Companys    string `json:"companys" gorm:"comment:公司"`
	Departments string `json:"departments" gorm:"comment:平台"`
	Roles       string `json:"roles" gorm:"comment:资金角色"`
}

type SearchUser struct {
	Uid int `json:"uid"`
}

func (FundUser) TableName() string {
	return "fund.fund_user"
}
