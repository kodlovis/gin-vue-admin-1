package rms

import "gin-vue-admin/model/rms"

type MacSearch struct{
    rms.Mac
    PageInfo
}