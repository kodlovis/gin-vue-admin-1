package future_data

import "gin-vue-admin/model/internal_system/future_data"

type SpotDetailSearch struct {
	future_data.SpotDetail
	PageInfo
}

type SpotDetail struct {
	ID             uint    `json:"id"`
	Time           string  `json:"time"`
	ProductName    string  `json:"productName"`
	AccountId      string  `json:"accountId"`
	ProfitByFloat  float64 `json:"profitByFloat"`
	ProfitByTrade  string  `json:"profitByTrade"`
	TradeFee       float64 `json:"tradeFee"`
	DepartmentName string  `json:"departmentName"`
}
