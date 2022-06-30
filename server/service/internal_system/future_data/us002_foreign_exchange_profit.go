package future_data

import (
	"gin-vue-admin/global"
	model "gin-vue-admin/model/internal_system/future_data"
	request "gin-vue-admin/model/request/internal_system/future_data"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateUs002ForeignExchangeProfit
//@description: 创建Us002ForeignExchangeProfit记录
//@param: us002ForeignExchangeProfit model.Us002ForeignExchangeProfit
//@return: err error

func CreateUs002ForeignExchangeProfit(us002ForeignExchangeProfit model.Us002ForeignExchangeProfit) (err error) {
	err = global.GVA_DB.Create(&us002ForeignExchangeProfit).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteUs002ForeignExchangeProfit
//@description: 删除Us002ForeignExchangeProfit记录
//@param: us002ForeignExchangeProfit model.Us002ForeignExchangeProfit
//@return: err error

func DeleteUs002ForeignExchangeProfit(us002ForeignExchangeProfit model.Us002ForeignExchangeProfit) (err error) {
	err = global.GVA_DB.Delete(&us002ForeignExchangeProfit).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteUs002ForeignExchangeProfitByIds
//@description: 批量删除Us002ForeignExchangeProfit记录
//@param: ids request.IdsReq
//@return: err error

func DeleteUs002ForeignExchangeProfitByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.Us002ForeignExchangeProfit{}, "id in ?", ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateUs002ForeignExchangeProfit
//@description: 更新Us002ForeignExchangeProfit记录
//@param: us002ForeignExchangeProfit *model.Us002ForeignExchangeProfit
//@return: err error

func UpdateUs002ForeignExchangeProfit(us002ForeignExchangeProfit model.Us002ForeignExchangeProfit) (err error) {
	err = global.GVA_DB.Save(&us002ForeignExchangeProfit).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetUs002ForeignExchangeProfit
//@description: 根据id获取Us002ForeignExchangeProfit记录
//@param: id uint
//@return: err error, us002ForeignExchangeProfit model.Us002ForeignExchangeProfit

func GetUs002ForeignExchangeProfit(id uint) (err error, us002ForeignExchangeProfit model.Us002ForeignExchangeProfit) {
	err = global.GVA_DB.Where("id = ?", id).First(&us002ForeignExchangeProfit).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetUs002ForeignExchangeProfitInfoList
//@description: 分页获取Us002ForeignExchangeProfit记录
//@param: info request.Us002ForeignExchangeProfitSearch
//@return: err error, list interface{}, total int64

func GetUs002ForeignExchangeProfitInfoList(info request.Us002ForeignExchangeProfitSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&model.Us002ForeignExchangeProfit{}).Preload("Department").Preload("Currencys")
	var us002ForeignExchangeProfits []model.Us002ForeignExchangeProfit
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Time != "" {
		db = db.Where("time = ?", info.Time)
	}
	if info.Currency != "" {
		db = db.Where("currency = ?", info.Currency)
	}
	if info.DepartmentCode != "" {
		db = db.Where("department_code = ?", info.DepartmentCode)
	}
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&us002ForeignExchangeProfits).Error
	return err, us002ForeignExchangeProfits, total
}
