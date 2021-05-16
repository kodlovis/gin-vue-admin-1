package rms

import "gin-vue-admin/model/rms"

type ProductSearch struct{
    rms.Product
    PageInfo
}