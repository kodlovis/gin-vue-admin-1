package position_reward

import (
	"gin-vue-admin/global"
	mp "gin-vue-admin/model/internal_system/position_reward"
	rp "gin-vue-admin/model/request/internal_system/position_reward"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateExchangeTradingfeeReward
//@description: 创建ExchangeTradingfeeReward记录
//@param: exchangeTradingfeeReward model.ExchangeTradingfeeReward
//@return: err error

func CreateExchangeTradingfeeReward(exchangeTradingfeeReward mp.ExchangeTradingfeeReward) (err error) {
	err = global.GVA_DB.Create(&exchangeTradingfeeReward).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteExchangeTradingfeeReward
//@description: 删除ExchangeTradingfeeReward记录
//@param: exchangeTradingfeeReward model.ExchangeTradingfeeReward
//@return: err error

func DeleteExchangeTradingfeeReward(exchangeTradingfeeReward mp.ExchangeTradingfeeReward) (err error) {
	err = global.GVA_DB.Delete(&exchangeTradingfeeReward).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteExchangeTradingfeeRewardByIds
//@description: 批量删除ExchangeTradingfeeReward记录
//@param: ids request.IdsReq
//@return: err error

func DeleteExchangeTradingfeeRewardByIds(ids rp.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]mp.ExchangeTradingfeeReward{}, "id in ?", ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateExchangeTradingfeeReward
//@description: 更新ExchangeTradingfeeReward记录
//@param: exchangeTradingfeeReward *model.ExchangeTradingfeeReward
//@return: err error

func UpdateExchangeTradingfeeReward(exchangeTradingfeeReward mp.ExchangeTradingfeeReward) (err error) {
	err = global.GVA_DB.Save(&exchangeTradingfeeReward).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetExchangeTradingfeeReward
//@description: 根据id获取ExchangeTradingfeeReward记录
//@param: id uint
//@return: err error, exchangeTradingfeeReward model.ExchangeTradingfeeReward

func GetExchangeTradingfeeReward(id uint) (err error, exchangeTradingfeeReward mp.ExchangeTradingfeeReward) {
	err = global.GVA_DB.Where("id = ?", id).First(&exchangeTradingfeeReward).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetExchangeTradingfeeRewardInfoList
//@description: 分页获取ExchangeTradingfeeReward记录
//@param: info request.ExchangeTradingfeeRewardSearch
//@return: err error, list interface{}, total int64

func GetExchangeTradingfeeRewardInfoList(info rp.ExchangeTradingfeeRewardSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&mp.ExchangeTradingfeeReward{}).Order("effective_date desc")
	var exchangeTradingfeeRewards []mp.ExchangeTradingfeeReward
	// 如果有条件搜索 下方会自动创建搜索语句
	if !info.EffectiveDate.IsZero() {
		db = db.Where("effective_date = ?", info.EffectiveDate)
	}
	if info.ExchangeId != "" {
		db = db.Where("exchange_id LIKE ?", "%"+info.ExchangeId+"%")
	}
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&exchangeTradingfeeRewards).Error
	return err, exchangeTradingfeeRewards, total
}
