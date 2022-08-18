// 自动生成模板LetterOfCreditDetail
package future_data

import (
	"gin-vue-admin/model/internal_system/public"
	"time"
)

// 如果含有time.Time 请自行import time包
type LetterOfCreditDetail struct {
	ID             uint `json:"ID"  gorm:"column:id;comment:ID;"`
	CreatedAt      time.Time
	CreatedUser    uint                 `json:"createdUser" form:"createdUser" gorm:"column:created_user;comment:;type:bigint;size:64;"`
	CreditId       string               `json:"creditId" form:"creditId" gorm:"column:credit_id;comment:;type:character varying;"`
	MaturityAmount float64              `json:"maturityAmount" form:"maturityAmount" gorm:"column:maturity_amount;comment:;type:double precision;size:53;"`
	MaturityDate   time.Time            `json:"maturityDate" form:"maturityDate" gorm:"column:maturity_date;comment:;type:date;size:0;"`
	InitialRate    float64              `json:"initialRate" form:"initialRate" gorm:"column:initial_rate;comment:;type:double precision;size:53;"`
	PurchaseRate   float64              `json:"purchaseRate" form:"purchaseRate" gorm:"column:purchase_rate;comment:;type:double precision;size:53;"`
	Currency       string               `json:"currency" form:"currency" gorm:"column:currency;comment:;type:varchar;"`
	ProductCode    string               `json:"productCode" form:"productCode" gorm:"column:product_code;comment:;type:varchar;"`
	UserInfo       SysUser              `json:"userInfo" gorm:"foreignKey:CreatedUser;references:ID;comment:"`
	Currencys      public.Us100Currency `json:"currencys" gorm:"foreignKey:Currency;references:CurrencyCode;comment:"`
}

func (LetterOfCreditDetail) TableName() string {
	return "master_data.us002_letter_of_credit_detail"
}
