package futureData

import (
	"gin-vue-admin/global"
	mif "gin-vue-admin/model/internalSystem/futureData"
	rif "gin-vue-admin/model/request/internalSystem/futureData"
	"math"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateRightsDetail
//@description: 创建RightsDetail记录
//@param: rightsDetail model.RightsDetail
//@return: err error

func CreateRightsDetail(rightsDetail mif.RightsDetail) (err error) {
	err = global.GVA_DB.Create(&rightsDetail).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteRightsDetail
//@description: 删除RightsDetail记录
//@param: rightsDetail model.RightsDetail
//@return: err error

func DeleteRightsDetail(rightsDetail mif.RightsDetail) (err error) {
	err = global.GVA_DB.Delete(&rightsDetail).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteRightsDetailByIds
//@description: 批量删除RightsDetail记录
//@param: ids request.IdsReq
//@return: err error

func DeleteRightsDetailByIds(ids rif.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]mif.RightsDetail{}, "id in ?", ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateRightsDetail
//@description: 更新RightsDetail记录
//@param: rightsDetail *model.RightsDetail
//@return: err error

func UpdateRightsDetail(rightsDetail mif.RightsDetail) (err error) {
	err = global.GVA_DB.Save(&rightsDetail).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetRightsDetail
//@description: 根据id获取RightsDetail记录
//@param: id uint
//@return: err error, rightsDetail model.RightsDetail

func GetRightsDetail(id uint) (err error, rightsDetail mif.RightsDetail) {
	err = global.GVA_DB.Where("id = ?", id).First(&rightsDetail).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetRightsDetailInfoList
//@description: 分页获取RightsDetail记录
//@param: info request.RightsDetailSearch
//@return: err error, list interface{}, total int64

func GetRightsDetailInfoList(info rif.RightsDetailSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&mif.RightsDetail{}).Order("time desc, department_name")
	var rightsDetails []mif.RightsDetail
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Time != "" {
		db = db.Where("time = ?", info.Time)
	}
	if info.ProductName != "" {
		db = db.Where("`product_name` LIKE ?", "%"+info.ProductName+"%")
	}
	if info.DepartmentName != "" {
		db = db.Where("`department_name` LIKE ?", "%"+info.DepartmentName+"%")
	}
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&rightsDetails).Error
	for i := 0; i < len(rightsDetails); i++ {
		rightsDetails[i].ThreeCharges = (math.Floor(rightsDetails[i].ThreeCharges*10 + 0.5)) / 10
		rightsDetails[i].FutureCharges = (math.Floor(rightsDetails[i].FutureCharges*10 + 0.5)) / 10
		rightsDetails[i].SundryExpense = (math.Floor(rightsDetails[i].SundryExpense*10 + 0.5)) / 10
	}
	return err, rightsDetails, total
}
