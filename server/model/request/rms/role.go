package rms

import "gin-vue-admin/model/rms"

type RoleSearch struct{
    rms.Role
    PageInfo
}