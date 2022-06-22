// 自动生成模板RightsDetail
package future_data

// 如果含有time.Time 请自行import time包
type RightsDetail struct {
	ID                  uint    `json:"id"  gorm:"column:id;comment:ID;"`
	Time                string  `json:"time" form:"time" gorm:"column:time;comment:;type:varchar;"`
	ProductName         string  `json:"productName" form:"productName" gorm:"column:product_name;comment:;type:varchar;"`
	DepartmentName      string  `json:"departmentName" form:"departmentName" gorm:"column:department_name;comment:;type:varchar;"`
	ThreeCharges        float64 `json:"threeCharges" form:"threeCharges" gorm:"column:three_charges;comment:;type:double precision;size:53;"`
	SundryExpense       float64 `json:"sundryExpense" form:"sundryExpense" gorm:"column:sundry_expense;comment:;type:double precision;size:53;"`
	FutureCharges       float64 `json:"futureCharges" form:"futureCharges" gorm:"column:future_charges;comment:;type:double precision;size:53;"`
	CumulativeRights    float64 `json:"cumulativeRights" form:"cumulativeRights" gorm:"column:cumulative_rights;comment:;type:double precision;size:53;"`
	AdjustExpense       float64 `json:"adjustExpense" form:"adjustExpense" gorm:"column:adjust_expense;comment:;type:double precision;size:53;"`
	CustomerProfit      float64 `json:"customerProfit" form:"customerProfit" gorm:"column:customer_profit;comment:;type:double precision;size:53;"`
	ExchangeGainsLosses float64 `json:"exchangeGainsLosses" form:"exchangeGainsLosses" gorm:"column:exchange_gains_losses;comment:;type:double precision;size:53;"`
	FinancingInterest   float64 `json:"financingInterest" form:"financingInterest" gorm:"column:financing_interest;comment:;type:double precision;size:53;"`
	UserID              uint    `json:"userID" form:"userID" gorm:"column:user_id;comment:;type:double precision;size:53;"`

	UserInfo SysUser `json:"userInfo" gorm:"foreignKey:ID;references:UserID;comment:"`
}

func (RightsDetail) TableName() string {
	return "future.rights_detail"
}
