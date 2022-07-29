package fund

import (
	"gin-vue-admin/global"
	model "gin-vue-admin/model/internal_system/fund"
	request "gin-vue-admin/model/request/internal_system/fund"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateRemainingSum
//@description: 创建RemainingSum记录
//@param: remainingSum model.RemainingSum
//@return: err error

func CreateRemainingSum(remainingSum model.RemainingSum) (err error) {
	err = global.GVA_DB.Create(&remainingSum).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteRemainingSum
//@description: 删除RemainingSum记录
//@param: remainingSum model.RemainingSum
//@return: err error

func DeleteRemainingSum(remainingSum model.RemainingSum) (err error) {
	err = global.GVA_DB.Delete(&remainingSum).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteRemainingSumByIds
//@description: 批量删除RemainingSum记录
//@param: ids request.IdsReq
//@return: err error

func DeleteRemainingSumByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.RemainingSum{}, "id in ?", ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateRemainingSum
//@description: 更新RemainingSum记录
//@param: remainingSum *model.RemainingSum
//@return: err error

func UpdateRemainingSum(remainingSum model.RemainingSum) (err error) {
	err = global.GVA_DB.Save(&remainingSum).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetRemainingSum
//@description: 根据id获取RemainingSum记录
//@param: id uint
//@return: err error, remainingSum model.RemainingSum

func GetRemainingSum(id uint) (err error, remainingSum model.RemainingSum) {
	err = global.GVA_DB.Where("id = ?", id).First(&remainingSum).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetRemainingSumInfoList
//@description: 分页获取RemainingSum记录
//@param: info request.RemainingSumSearch
//@return: err error, list interface{}, total int64

func GetRemainingSumInfoList(info request.RemainingSumSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&model.RemainingSum{})
	var remainingSums []model.RemainingSum
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&remainingSums).Error
	return err, remainingSums, total
}

func SearchRemainingSum(info model.SearchRemainingSum) (err error, list interface{}) {
	// 创建db
	db := global.GVA_DB.Model(&model.RemainingSum{})
	var remainingSums []model.RemainingSum
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Where("type = ? and up_date <= ? and company = ?", info.Type, info.UpDate, info.Company).Last(&remainingSums).Error

	return err, remainingSums
}

func SearchSum(info model.SearchRemainingSum) (err error, list interface{}) {
	// 创建db
	db := global.GVA_DB.Model(&model.RemainingSum{})
	var remainingSums []model.RemainingSum
	// 如果有条件搜索 下方会自动创建搜索语句
	var remainingSumOne []model.RemainingSum
	err = db.Where("type = ? and up_date <= ? ", info.Type, info.UpDate).Last(&remainingSumOne).Error
	db2 := global.GVA_DB.Model(&model.RemainingSum{})
	db2.Where("type = ? and up_date > ? ", info.Type, info.UpDate).Find(&remainingSums)
	remainingSums = append(remainingSums, remainingSumOne...)
	return err, remainingSums
}

func SearchUser(info model.SearchUser) (err error, list interface{}) {
	// 创建db
	db := global.GVA_DB.Model(&model.FundUser{})
	var FundUsers []model.FundUser

	err = db.Where("uid = ? ", info.Uid).Last(&FundUsers).Error

	return err, FundUsers
}
