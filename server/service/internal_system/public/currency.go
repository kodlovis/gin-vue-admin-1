package public

import (
	"gin-vue-admin/global"
	model "gin-vue-admin/model/internal_system/public"
	request "gin-vue-admin/model/request/internal_system/public"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateUs100Currency
//@description: 创建Us100Currency记录
//@param: us100Currency model.Us100Currency
//@return: err error

func CreateUs100Currency(us100Currency model.Us100Currency) (err error) {
	err = global.GVA_DB.Create(&us100Currency).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteUs100Currency
//@description: 删除Us100Currency记录
//@param: us100Currency model.Us100Currency
//@return: err error

func DeleteUs100Currency(us100Currency model.Us100Currency) (err error) {
	err = global.GVA_DB.Delete(&us100Currency).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteUs100CurrencyByIds
//@description: 批量删除Us100Currency记录
//@param: ids request.IdsReq
//@return: err error

func DeleteUs100CurrencyByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.Us100Currency{}, "id in ?", ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateUs100Currency
//@description: 更新Us100Currency记录
//@param: us100Currency *model.Us100Currency
//@return: err error

func UpdateUs100Currency(us100Currency model.Us100Currency) (err error) {
	err = global.GVA_DB.Save(&us100Currency).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetUs100Currency
//@description: 根据id获取Us100Currency记录
//@param: id uint
//@return: err error, us100Currency model.Us100Currency

func GetUs100Currency(id uint) (err error, us100Currency model.Us100Currency) {
	err = global.GVA_DB.Where("id = ?", id).First(&us100Currency).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetUs100CurrencyInfoList
//@description: 分页获取Us100Currency记录
//@param: info request.Us100CurrencySearch
//@return: err error, list interface{}, total int64

func GetUs100CurrencyInfoList(info request.Us100CurrencySearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&model.Us100Currency{})
	var us100Currencys []model.Us100Currency
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&us100Currencys).Error
	return err, us100Currencys, total
}
