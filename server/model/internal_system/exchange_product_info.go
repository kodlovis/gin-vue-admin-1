// 自动生成模板ExchangeProductInfo
package internal_system

// 如果含有time.Time 请自行import time包
type ExchangeProductInfo struct {
	ID                 uint    `json:"id"  gorm:"column:id;comment:ID;"`
	ExchangeName       string  `json:"exchangeName" form:"exchangeName" gorm:"column:exchange_name;comment:;type:varchar;"`
	ProductCode        string  `json:"productCode" form:"productCode" gorm:"column:product_code;comment:;type:varchar;"`
	ProductName        string  `json:"productName" form:"productName" gorm:"column:product_name;comment:;type:varchar;"`
	DeliveryUnitLot    int     `json:"deliveryUnitLot" form:"deliveryUnitLot" gorm:"column:delivery_unit_lot;comment:;type:int;"`
	DeliveryUnitTon    int     `json:"deliveryUnitTon" form:"deliveryUnitTon" gorm:"column:delivery_unit_ton;comment:;type:int;"`
	ContractMultiplier int     `json:"contractMultiplier" form:"contractMultiplier" gorm:"column:contract_multiplier;comment:;type:int;"`
	TaxRate            float64 `json:"taxRate" form:"taxRate" gorm:"column:tax_rate;comment:;type:double;"`
}

func (ExchangeProductInfo) TableName() string {
	return "future.exchange_product_info"
}
