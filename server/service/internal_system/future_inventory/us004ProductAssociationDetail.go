package future_inventory

import (
	"gin-vue-admin/global"
	mif "gin-vue-admin/model/internal_system/future_inventory"
	rif "gin-vue-admin/model/request/internal_system/future_inventory"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateUs004ProductAssociationDetail
//@description: 创建Us004ProductAssociationDetail记录
//@param: us004ProductAssociationDetail model.Us004ProductAssociationDetail
//@return: err error

func CreateUs004ProductAssociationDetail(us004ProductAssociationDetail mif.Us004ProductAssociationDetail) (err error) {
	err = global.GVA_DB.Create(&us004ProductAssociationDetail).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteUs004ProductAssociationDetail
//@description: 删除Us004ProductAssociationDetail记录
//@param: us004ProductAssociationDetail model.Us004ProductAssociationDetail
//@return: err error

func DeleteUs004ProductAssociationDetail(us004ProductAssociationDetail mif.Us004ProductAssociationDetail) (err error) {
	err = global.GVA_DB.Delete(&us004ProductAssociationDetail).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteUs004ProductAssociationDetailByIds
//@description: 批量删除Us004ProductAssociationDetail记录
//@param: ids request.IdsReq
//@return: err error

func DeleteUs004ProductAssociationDetailByIds(ids rif.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]mif.Us004ProductAssociationDetail{}, "id in ?", ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateUs004ProductAssociationDetail
//@description: 更新Us004ProductAssociationDetail记录
//@param: us004ProductAssociationDetail *model.Us004ProductAssociationDetail
//@return: err error

func UpdateUs004ProductAssociationDetail(us004ProductAssociationDetail mif.Us004ProductAssociationDetail) (err error) {
	err = global.GVA_DB.Save(&us004ProductAssociationDetail).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetUs004ProductAssociationDetail
//@description: 根据id获取Us004ProductAssociationDetail记录
//@param: id uint
//@return: err error, us004ProductAssociationDetail model.Us004ProductAssociationDetail

func GetUs004ProductAssociationDetail(id uint) (err error, us004ProductAssociationDetail mif.Us004ProductAssociationDetail) {
	err = global.GVA_DB.Where("id = ?", id).First(&us004ProductAssociationDetail).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetUs004ProductAssociationDetailInfoList
//@description: 分页获取Us004ProductAssociationDetail记录
//@param: info request.Us004ProductAssociationDetailSearch
//@return: err error, list interface{}, total int64

func GetUs004ProductAssociationDetailInfoList(info rif.Us004ProductAssociationDetailSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&mif.Us004ProductAssociationDetail{})
	var us004ProductAssociationDetails []mif.Us004ProductAssociationDetail
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.ProductName != "" {
		db = db.Where("product_name = ?", info.ProductName)
	}
	if info.SubProductName != "" {
		db = db.Where("sub_product_name = ?", info.SubProductName)
	}
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&us004ProductAssociationDetails).Error
	return err, us004ProductAssociationDetails, total
}
