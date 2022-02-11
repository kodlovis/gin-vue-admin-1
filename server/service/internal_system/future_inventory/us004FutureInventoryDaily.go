package future_inventory

import (
	"gin-vue-admin/global"
	mif "gin-vue-admin/model/internal_system/future_inventory"
	rif "gin-vue-admin/model/request/internal_system/future_inventory"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateUs004FutureInventoryDaily
//@description: 创建Us004FutureInventoryDaily记录
//@param: us004FutureInventoryDaily model.Us004FutureInventoryDaily
//@return: err error

func CreateUs004FutureInventoryDaily(us004FutureInventoryDaily mif.Us004FutureInventoryDaily) (err error) {
	err = global.GVA_DB.Create(&us004FutureInventoryDaily).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteUs004FutureInventoryDaily
//@description: 删除Us004FutureInventoryDaily记录
//@param: us004FutureInventoryDaily model.Us004FutureInventoryDaily
//@return: err error

func DeleteUs004FutureInventoryDaily(us004FutureInventoryDaily mif.Us004FutureInventoryDaily) (err error) {
	err = global.GVA_DB.Delete(&us004FutureInventoryDaily).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteUs004FutureInventoryDailyByIds
//@description: 批量删除Us004FutureInventoryDaily记录
//@param: ids request.IdsReq
//@return: err error

func DeleteUs004FutureInventoryDailyByIds(ids rif.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]mif.Us004FutureInventoryDaily{}, "id in ?", ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateUs004FutureInventoryDaily
//@description: 更新Us004FutureInventoryDaily记录
//@param: us004FutureInventoryDaily *model.Us004FutureInventoryDaily
//@return: err error

func UpdateUs004FutureInventoryDaily(us004FutureInventoryDaily mif.Us004FutureInventoryDaily) (err error) {
	err = global.GVA_DB.Save(&us004FutureInventoryDaily).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetUs004FutureInventoryDaily
//@description: 根据id获取Us004FutureInventoryDaily记录
//@param: id uint
//@return: err error, us004FutureInventoryDaily model.Us004FutureInventoryDaily

func GetUs004FutureInventoryDaily(id uint) (err error, us004FutureInventoryDaily mif.Us004FutureInventoryDaily) {
	err = global.GVA_DB.Where("id = ?", id).First(&us004FutureInventoryDaily).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetUs004FutureInventoryDailyInfoList
//@description: 分页获取Us004FutureInventoryDaily记录
//@param: info request.Us004FutureInventoryDailySearch
//@return: err error, list interface{}, total int64

func GetUs004FutureInventoryDailyInfoList(info rif.Us004FutureInventoryDailySearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&mif.Us004FutureInventoryDaily{}).Order("time desc")
	var us004FutureInventoryDailys []mif.Us004FutureInventoryDaily
	// 如果有条件搜索 下方会自动创建搜索语句
	if !info.Time.IsZero() {
		db = db.Where("time = ?", info.Time)
	}
	if info.ProductCode != "" {
		db = db.Where("product_code LIKE ?", "%"+info.ProductCode+"%")
	}
	if info.ExchangeId != "" {
		db = db.Where("exchange_id LIKE ?", "%"+info.ExchangeId+"%")
	}
	if info.Comment != "" {
		db = db.Where("comment LIKE ?", "%"+info.Comment+"%")
	}
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&us004FutureInventoryDailys).Error
	return err, us004FutureInventoryDailys, total
}
