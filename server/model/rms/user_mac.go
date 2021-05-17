
package rms

// 如果含有time.Time 请自行import time包
type UserMac struct {
      UserId  int `json:"userId" form:"userId" gorm:"column:users_id;comment:用户ID;type:int;"`
      MacId  int `json:"macId" form:"macId" gorm:"column:mac_id;comment:macID;type:int;"`
}


func (UserMac) TableName() string {
  return "users_mac"
}
