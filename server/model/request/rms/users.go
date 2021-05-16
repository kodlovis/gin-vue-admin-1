package rms

import "gin-vue-admin/model/rms"

type UsersSearch struct{
    rms.Users
    PageInfo
    ID string `json:"userid" form:"userid"`
    Password string `json:"password" form:"password"`
    Mac string `json:"mac" form:"mac"`
}