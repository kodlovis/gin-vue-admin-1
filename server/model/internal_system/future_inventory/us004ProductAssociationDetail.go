// 自动生成模板Us004ProductAssociationDetail
package future_inventory

// 如果含有time.Time 请自行import time包
type Us004ProductAssociationDetail struct {
	ID             uint   `json:"ID"  gorm:"column:id;comment:ID;"`
	ProductName    string `json:"productName" form:"productName" gorm:"column:product_name;comment:;type:varchar;"`
	SubProductName string `json:"subProductName" form:"subProductName" gorm:"column:sub_product_name;comment:"`
	Unit           string `json:"unit" form:"unit" gorm:"column:unit;comment:;type:character varying;"`
}

func (Us004ProductAssociationDetail) TableName() string {
	return "future.us004_product_association_detail"
}