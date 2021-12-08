package future_data

import (
	"gin-vue-admin/global"
	mif "gin-vue-admin/model/internal_system/future_data"
	rif "gin-vue-admin/model/request/internal_system/future_data"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateTradeDetail
//@description: 创建TradeDetail记录
//@param: tradeDetail model.TradeDetail
//@return: err error

func CreateTradeDetail(tradeDetail mif.TradeDetail) (err error) {
	err = global.GVA_DB.Create(&tradeDetail).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteTradeDetail
//@description: 删除TradeDetail记录
//@param: tradeDetail model.TradeDetail
//@return: err error

func DeleteTradeDetail(tradeDetail mif.TradeDetail) (err error) {
	err = global.GVA_DB.Delete(&tradeDetail).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteTradeDetailByIds
//@description: 批量删除TradeDetail记录
//@param: ids request.IdsReq
//@return: err error

func DeleteTradeDetailByIds(ids rif.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]mif.TradeDetail{}, "id in ?", ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateTradeDetail
//@description: 更新TradeDetail记录
//@param: tradeDetail *model.TradeDetail
//@return: err error

func UpdateTradeDetail(tradeDetail mif.TradeDetail) (err error) {
	err = global.GVA_DB.Save(&tradeDetail).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetTradeDetail
//@description: 根据id获取TradeDetail记录
//@param: id uint
//@return: err error, tradeDetail model.TradeDetail

func GetTradeDetail(id uint) (err error, tradeDetail mif.TradeDetail) {
	err = global.GVA_DB.Where("id = ?", id).First(&tradeDetail).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetTradeDetailInfoList
//@description: 分页获取TradeDetail记录
//@param: info request.TradeDetailSearch
//@return: err error, list interface{}, total int64

func GetTradeDetailInfoList(info rif.TradeDetailSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&mif.TradeDetail{})
	var tradeDetails []mif.TradeDetail
	// 如果有条件搜索 下方会自动创建搜索语句
	if !info.Time.IsZero() {
		db = db.Where("time = ?", info.Time)
	}
	if info.AccountId != "" {
		db = db.Where("account_id LIKE ?", "%"+info.AccountId+"%")
	}
	if info.Instrument != "" {
		db = db.Where("instrument LIKE ?", "%"+info.Instrument+"%")
	}
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&tradeDetails).Error
	return err, tradeDetails, total
}
