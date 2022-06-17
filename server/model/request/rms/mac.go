package rms

import "gin-vue-admin/model/rms"

type MacSearch struct {
	rms.Mac
	PageInfo
	names []string `json:"names"`
}
type MacList struct {
	MacList []rms.Mac `json:"macs"`
}
