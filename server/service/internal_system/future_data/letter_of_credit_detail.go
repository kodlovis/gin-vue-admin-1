package future_data

import (
	"gin-vue-admin/global"
	model "gin-vue-admin/model/internal_system/future_data"
	request "gin-vue-admin/model/request/internal_system/future_data"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateLetterOfCreditDetail
//@description: 创建LetterOfCreditDetail记录
//@param: letterOfCreditDetail model.LetterOfCreditDetail
//@return: err error

func CreateLetterOfCreditDetail(letterOfCreditDetail model.LetterOfCreditDetail) (err error) {
	err = global.GVA_DB.Create(&letterOfCreditDetail).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteLetterOfCreditDetail
//@description: 删除LetterOfCreditDetail记录
//@param: letterOfCreditDetail model.LetterOfCreditDetail
//@return: err error

func DeleteLetterOfCreditDetail(letterOfCreditDetail model.LetterOfCreditDetail) (err error) {
	err = global.GVA_DB.Delete(&letterOfCreditDetail).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteLetterOfCreditDetailByIds
//@description: 批量删除LetterOfCreditDetail记录
//@param: ids request.IdsReq
//@return: err error

func DeleteLetterOfCreditDetailByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.LetterOfCreditDetail{}, "id in ?", ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateLetterOfCreditDetail
//@description: 更新LetterOfCreditDetail记录
//@param: letterOfCreditDetail *model.LetterOfCreditDetail
//@return: err error

func UpdateLetterOfCreditDetail(letterOfCreditDetail model.LetterOfCreditDetail) (err error) {
	err = global.GVA_DB.Save(&letterOfCreditDetail).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetLetterOfCreditDetail
//@description: 根据id获取LetterOfCreditDetail记录
//@param: id uint
//@return: err error, letterOfCreditDetail model.LetterOfCreditDetail

func GetLetterOfCreditDetail(id uint) (err error, letterOfCreditDetail model.LetterOfCreditDetail) {
	err = global.GVA_DB.Where("id = ?", id).First(&letterOfCreditDetail).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetLetterOfCreditDetailInfoList
//@description: 分页获取LetterOfCreditDetail记录
//@param: info request.LetterOfCreditDetailSearch
//@return: err error, list interface{}, total int64

func GetLetterOfCreditDetailInfoList(info request.LetterOfCreditDetailSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&model.LetterOfCreditDetail{}).Preload("UserInfo").Preload("Currencys").Preload("DepartmentInfo").Preload("ProductInfo").Order("created_at desc")
	var letterOfCreditDetails []model.LetterOfCreditDetail
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.CreatedUser != 0 {
		db = db.Where("created_user = ?", info.CreatedUser)
	}
	if info.Currency != "" {
		db = db.Where("currency = ?", info.Currency)
	}
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&letterOfCreditDetails).Error
	return err, letterOfCreditDetails, total
}
func GetLetterOfCreditDetailListWithNoPurchaseRate(info request.LetterOfCreditDetailSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&model.LetterOfCreditDetail{}).Preload("UserInfo").Preload("Currencys").Preload("DepartmentInfo").Preload("ProductInfo").Where("purchase_rate = 0").Order("created_at desc")
	var letterOfCreditDetails []model.LetterOfCreditDetail
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.CreatedUser != 0 {
		db = db.Where("created_user = ?", info.CreatedUser)
	}
	if info.Currency != "" {
		db = db.Where("currency = ?", info.Currency)
	}
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&letterOfCreditDetails).Error
	return err, letterOfCreditDetails, total
}

func UpdateLetterOfCreditPurchaseRate(letterOfCreditDetail model.LetterOfCreditDetail) (err error) {
	err = global.GVA_DB.Model(&letterOfCreditDetail).Where("id = ?", letterOfCreditDetail.ID).Select("purchase_rate").Updates(map[string]interface{}{"purchase_rate": letterOfCreditDetail.PurchaseRate}).Error
	return err
}
