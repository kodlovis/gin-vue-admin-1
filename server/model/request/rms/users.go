package rms

import "gin-vue-admin/model/rms"

type UsersSearch struct{
    rms.Users
    PageInfo
    ID string `json:"userid" form:"userid"`
    RoleID int `json:"roleid" form:"roleid"`
    Password string `json:"password" form:"password"`
    Mac string `json:"mac" form:"mac"`
}

type UserMacList struct{
	UserMacList []rms.UserMac `json:"item"`
}
type UserProductList struct{
	UserProductList []rms.UserProduct `json:"list"`
}