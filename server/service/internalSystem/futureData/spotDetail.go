package futureData

import (
	"gin-vue-admin/global"
	mif "gin-vue-admin/model/internalSystem/futureData"
	rif "gin-vue-admin/model/request/internalSystem/futureData"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateSpotDetail
//@description: 创建SpotDetail记录
//@param: spotDetail model.SpotDetail
//@return: err error

func CreateSpotDetail(spotDetail mif.SpotDetail) (err error) {
	err = global.GVA_DB.Create(&spotDetail).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteSpotDetail
//@description: 删除SpotDetail记录
//@param: spotDetail model.SpotDetail
//@return: err error

func DeleteSpotDetail(spotDetail mif.SpotDetail) (err error) {
	err = global.GVA_DB.Delete(&spotDetail).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteSpotDetailByIds
//@description: 批量删除SpotDetail记录
//@param: ids request.IdsReq
//@return: err error

func DeleteSpotDetailByIds(ids rif.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]mif.SpotDetail{}, "id in ?", ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateSpotDetail
//@description: 更新SpotDetail记录
//@param: spotDetail *model.SpotDetail
//@return: err error

func UpdateSpotDetail(spotDetail mif.SpotDetail) (err error) {
	err = global.GVA_DB.Save(&spotDetail).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetSpotDetail
//@description: 根据id获取SpotDetail记录
//@param: id uint
//@return: err error, spotDetail model.SpotDetail

func GetSpotDetail(id uint) (err error, spotDetail mif.SpotDetail) {
	err = global.GVA_DB.Where("id = ?", id).First(&spotDetail).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetSpotDetailInfoList
//@description: 分页获取SpotDetail记录
//@param: info request.SpotDetailSearch
//@return: err error, list interface{}, total int64

func GetSpotDetailInfoList(info rif.SpotDetailSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&mif.SpotDetail{})
	var spotDetails []mif.SpotDetail
	// 如果有条件搜索 下方会自动创建搜索语句
	if !info.Time.IsZero() {
		db = db.Where("time = ?", info.Time)
	}
	if info.ProductName != "" {
		db = db.Where("product_name like ?", "%"+info.ProductName+"%")
	}
	if info.AccountId != "" {
		db = db.Where("account_id LIKE ?", "%"+info.AccountId+"%")
	}
	if info.DepartmentName != "" {
		db = db.Where("department_name LIKE ?", "%"+info.DepartmentName+"%")
	}
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&spotDetails).Error
	return err, spotDetails, total
}
