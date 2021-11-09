package service

import (
	"gin-vue-admin/global"
	mi "gin-vue-admin/model/internal_system"
	ri "gin-vue-admin/model/request/internal_system"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateAccountInfo
//@description: 创建AccountInfo记录
//@param: accountInfo model.AccountInfo
//@return: err error

func CreateAccountInfo(accountInfo mi.AccountInfo) (err error) {
	err = global.GVA_DB.Create(&accountInfo).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteAccountInfo
//@description: 删除AccountInfo记录
//@param: accountInfo model.AccountInfo
//@return: err error

func DeleteAccountInfo(accountInfo mi.AccountInfo) (err error) {
	err = global.GVA_DB.Delete(&accountInfo).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteAccountInfoByIds
//@description: 批量删除AccountInfo记录
//@param: ids request.IdsReq
//@return: err error

func DeleteAccountInfoByIds(ids ri.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]mi.AccountInfo{}, "id in ?", ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateAccountInfo
//@description: 更新AccountInfo记录
//@param: accountInfo *model.AccountInfo
//@return: err error

func UpdateAccountInfo(accountInfo mi.AccountInfo) (err error) {
	err = global.GVA_DB.Save(&accountInfo).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetAccountInfo
//@description: 根据id获取AccountInfo记录
//@param: id uint
//@return: err error, accountInfo model.AccountInfo

func GetAccountInfo(id uint) (err error, accountInfo mi.AccountInfo) {
	err = global.GVA_DB.Where("id = ?", id).First(&accountInfo).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetAccountInfoInfoList
//@description: 分页获取AccountInfo记录
//@param: info request.AccountInfoSearch
//@return: err error, list interface{}, total int64

func GetAccountInfoInfoList(info ri.AccountInfoSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&mi.AccountInfo{})
	var accountInfos []mi.AccountInfo
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&accountInfos).Error
	return err, accountInfos, total
}
