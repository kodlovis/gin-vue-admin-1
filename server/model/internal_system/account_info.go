package internal_system

// 如果含有time.Time 请自行import time包
type AccountInfo struct {
	ID          uint   `json:"id"  gorm:"column:id;comment:ID;"`
	AccountId   string `json:"accountId" form:"accountId" gorm:"column:account_id;comment:;type:character varying;"`
	FutureName  string `json:"futureName" form:"futureName" gorm:"column:future_name;comment:;type:character varying;"`
	CompanyCode string `json:"companyCode" form:"companyCode" gorm:"column:company_code;comment:;type:character varying;"`
	Comment     string `json:"comment" form:"comment" gorm:"column:comment;comment:;type:character varying;"`
	Type        string `json:"type" form:"comment" gorm:"column:type;comment:;type:character varying;"`
}

func (AccountInfo) TableName() string {
	return "future.account_info"
}
