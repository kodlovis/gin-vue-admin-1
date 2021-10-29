package future_data

import (
	"gin-vue-admin/global"
	mif "gin-vue-admin/model/internalSystem/future_data"
	rif "gin-vue-admin/model/request/internalSystem/future_data"
	"gin-vue-admin/utils"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateFutureAccount
//@description: 创建FutureAccount记录
//@param: futureAccount model.FutureAccount
//@return: err error

func CreateFutureAccount(futureAccount mif.FutureAccount) (err error) {
	futureAccount.Password = string(utils.Base64Encode([]byte(futureAccount.Password)))

	err = global.GVA_DB.Create(&futureAccount).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteFutureAccount
//@description: 删除FutureAccount记录
//@param: futureAccount model.FutureAccount
//@return: err error

func DeleteFutureAccount(futureAccount mif.FutureAccount) (err error) {
	err = global.GVA_DB.Delete(&futureAccount).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteFutureAccountByIds
//@description: 批量删除FutureAccount记录
//@param: ids request.IdsReq
//@return: err error

func DeleteFutureAccountByIds(ids rif.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]mif.FutureAccount{}, "id in ?", ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateFutureAccount
//@description: 更新FutureAccount记录
//@param: futureAccount *model.FutureAccount
//@return: err error

func UpdateFutureAccount(futureAccount mif.FutureAccount) (err error) {
	err = global.GVA_DB.Save(&futureAccount).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetFutureAccount
//@description: 根据id获取FutureAccount记录
//@param: id uint
//@return: err error, futureAccount model.FutureAccount

func GetFutureAccount(id uint) (err error, futureAccount mif.FutureAccount) {
	err = global.GVA_DB.Where("id = ?", id).First(&futureAccount).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetFutureAccountInfoList
//@description: 分页获取FutureAccount记录
//@param: info request.FutureAccountSearch
//@return: err error, list interface{}, total int64

func GetFutureAccountInfoList(info rif.FutureAccountSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&mif.FutureAccount{})
	var futureAccounts []mif.FutureAccount
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&futureAccounts).Error
	return err, futureAccounts, total
}
