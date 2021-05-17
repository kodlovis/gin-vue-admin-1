
package rms

// 如果含有time.Time 请自行import time包
type UserProduct struct {
      UserId  int `json:"userId" form:"userId" gorm:"column:users_id;comment:用户ID;type:int;"`
      ProductId  int `json:"productId" form:"productId" gorm:"column:product_id;comment:产品ID;type:int;"`
}


func (UserProduct) TableName() string {
  return "users_product"
}
