package rms

import "gin-vue-admin/model/rms"

type RiskSettingSearch struct{
    rms.RiskSetting
    PageInfo
}