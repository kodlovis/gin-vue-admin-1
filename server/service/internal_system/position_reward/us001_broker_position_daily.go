package position_reward

import (
	"gin-vue-admin/global"
	mp "gin-vue-admin/model/internal_system/position_reward"
	rp "gin-vue-admin/model/request/internal_system/position_reward"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateBrokerPositionDaily
//@description: 创建BrokerPositionDaily记录
//@param: BrokerPositionDaily model.BrokerPositionDaily
//@return: err error

func CreateBrokerPositionDaily(BrokerPositionDaily mp.BrokerPositionDaily) (err error) {
	err = global.GVA_DB.Create(&BrokerPositionDaily).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteBrokerPositionDaily
//@description: 删除BrokerPositionDaily记录
//@param: BrokerPositionDaily model.BrokerPositionDaily
//@return: err error

func DeleteBrokerPositionDaily(BrokerPositionDaily mp.BrokerPositionDaily) (err error) {
	err = global.GVA_DB.Delete(&BrokerPositionDaily).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteBrokerPositionDailyByIds
//@description: 批量删除BrokerPositionDaily记录
//@param: ids request.IdsReq
//@return: err error

func DeleteBrokerPositionDailyByIds(ids rp.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]mp.BrokerPositionDaily{}, "id in ?", ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateBrokerPositionDaily
//@description: 更新BrokerPositionDaily记录
//@param: BrokerPositionDaily *model.BrokerPositionDaily
//@return: err error

func UpdateBrokerPositionDaily(BrokerPositionDaily mp.BrokerPositionDaily) (err error) {
	err = global.GVA_DB.Save(&BrokerPositionDaily).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetBrokerPositionDaily
//@description: 根据id获取BrokerPositionDaily记录
//@param: id uint
//@return: err error, BrokerPositionDaily model.BrokerPositionDaily

func GetBrokerPositionDaily(id uint) (err error, BrokerPositionDaily mp.BrokerPositionDaily) {
	err = global.GVA_DB.Where("id = ?", id).First(&BrokerPositionDaily).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetBrokerPositionDailyInfoList
//@description: 分页获取BrokerPositionDaily记录
//@param: info request.BrokerPositionDailySearch
//@return: err error, list interface{}, total int64

func GetBrokerPositionDailyInfoList(info rp.BrokerPositionDailySearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&mp.BrokerPositionDaily{})
	var BrokerPositionDailys []mp.BrokerPositionDaily
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.TradingDate != "" {
		db = db.Where("trading_date = ?", info.TradingDate)
	}
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&BrokerPositionDailys).Error
	return err, BrokerPositionDailys, total
}
