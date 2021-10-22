package position_reward

import (
	"gin-vue-admin/global"
	mp "gin-vue-admin/model/internalSystem/position_reward"
	rp "gin-vue-admin/model/request/internalSystem/position_reward"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateAccountPositionDaily
//@description: 创建AccountPositionDaily记录
//@param: accountPositionDaily model.AccountPositionDaily
//@return: err error

func CreateAccountPositionDaily(accountPositionDaily mp.AccountPositionDaily) (err error) {
	err = global.GVA_DB.Create(&accountPositionDaily).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteAccountPositionDaily
//@description: 删除AccountPositionDaily记录
//@param: accountPositionDaily model.AccountPositionDaily
//@return: err error

func DeleteAccountPositionDaily(accountPositionDaily mp.AccountPositionDaily) (err error) {
	err = global.GVA_DB.Delete(&accountPositionDaily).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteAccountPositionDailyByIds
//@description: 批量删除AccountPositionDaily记录
//@param: ids request.IdsReq
//@return: err error

func DeleteAccountPositionDailyByIds(ids rp.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]mp.AccountPositionDaily{}, "id in ?", ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateAccountPositionDaily
//@description: 更新AccountPositionDaily记录
//@param: accountPositionDaily *model.AccountPositionDaily
//@return: err error

func UpdateAccountPositionDaily(accountPositionDaily mp.AccountPositionDaily) (err error) {
	err = global.GVA_DB.Save(&accountPositionDaily).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetAccountPositionDaily
//@description: 根据id获取AccountPositionDaily记录
//@param: id uint
//@return: err error, accountPositionDaily model.AccountPositionDaily

func GetAccountPositionDaily(id uint) (err error, accountPositionDaily mp.AccountPositionDaily) {
	err = global.GVA_DB.Where("id = ?", id).First(&accountPositionDaily).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetAccountPositionDailyInfoList
//@description: 分页获取AccountPositionDaily记录
//@param: info request.AccountPositionDailySearch
//@return: err error, list interface{}, total int64

func GetAccountPositionDailyInfoList(info rp.AccountPositionDailySearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&mp.AccountPositionDaily{})
	var us001AccountPositionDailys []mp.AccountPositionDaily
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.TradingDate != "" {
		db = db.Where("trading_date = ?", info.TradingDate)
	}
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&us001AccountPositionDailys).Error
	return err, us001AccountPositionDailys, total
}
