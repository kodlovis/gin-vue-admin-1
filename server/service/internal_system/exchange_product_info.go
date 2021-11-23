package service

import (
	"gin-vue-admin/global"
	mi "gin-vue-admin/model/internal_system"
	ri "gin-vue-admin/model/request/internal_system"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateExchangeProductInfo
//@description: 创建ExchangeProductInfo记录
//@param: exchangeProductInfo model.ExchangeProductInfo
//@return: err error

func CreateExchangeProductInfo(exchangeProductInfo mi.ExchangeProductInfo) (err error) {
	err = global.GVA_DB.Create(&exchangeProductInfo).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteExchangeProductInfo
//@description: 删除ExchangeProductInfo记录
//@param: exchangeProductInfo model.ExchangeProductInfo
//@return: err error

func DeleteExchangeProductInfo(exchangeProductInfo mi.ExchangeProductInfo) (err error) {
	err = global.GVA_DB.Delete(&exchangeProductInfo).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteExchangeProductInfoByIds
//@description: 批量删除ExchangeProductInfo记录
//@param: ids request.IdsReq
//@return: err error

func DeleteExchangeProductInfoByIds(ids ri.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]mi.ExchangeProductInfo{}, "id in ?", ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateExchangeProductInfo
//@description: 更新ExchangeProductInfo记录
//@param: exchangeProductInfo *model.ExchangeProductInfo
//@return: err error

func UpdateExchangeProductInfo(exchangeProductInfo mi.ExchangeProductInfo) (err error) {
	err = global.GVA_DB.Save(&exchangeProductInfo).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetExchangeProductInfo
//@description: 根据id获取ExchangeProductInfo记录
//@param: id uint
//@return: err error, exchangeProductInfo model.ExchangeProductInfo

func GetExchangeProductInfo(id uint) (err error, exchangeProductInfo mi.ExchangeProductInfo) {
	err = global.GVA_DB.Where("id = ?", id).First(&exchangeProductInfo).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetExchangeProductInfoInfoList
//@description: 分页获取ExchangeProductInfo记录
//@param: info request.ExchangeProductInfoSearch
//@return: err error, list interface{}, total int64

func GetExchangeProductInfoInfoList(info ri.ExchangeProductInfoSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&mi.ExchangeProductInfo{})
	var exchangeProductInfos []mi.ExchangeProductInfo
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.ExchangeName != "" {
		db = db.Where("exchange_name LIKE ?", "%"+info.ExchangeName+"%")
	}
	if info.ProductCode != "" {
		db = db.Where("product_code LIKE ?", "%"+info.ProductCode+"%")
	}
	if info.ProductName != "" {
		db = db.Where("product_name LIKE ?", "%"+info.ProductName+"%")
	}
	if info.DeliveryUnitLot != 0 {
		db = db.Where("delivery_unit_lot = ?", info.DeliveryUnitLot)
	}
	if info.DeliveryUnitTon != 0 {
		db = db.Where("delivery_unit_ton = ?", info.DeliveryUnitTon)
	}
	if info.ContractMultiplier != 0 {
		db = db.Where("contract_multiplier = ?", info.ContractMultiplier)
	}
	if info.TaxRate != 0 {
		db = db.Where("tax_rate = ?", info.TaxRate)
	}
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&exchangeProductInfos).Error
	return err, exchangeProductInfos, total
}
