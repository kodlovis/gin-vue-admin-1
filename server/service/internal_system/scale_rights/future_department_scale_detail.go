package scale_rights

import (
	"gin-vue-admin/global"
	mp "gin-vue-admin/model/internal_system/scale_rights"
	rp "gin-vue-admin/model/request/internal_system/scale_rights"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateFutureDepartmentScaleDetail
//@description: 创建FutureDepartmentScaleDetail记录
//@param: futureDepartmentScaleDetail model.FutureDepartmentScaleDetail
//@return: err error

func CreateFutureDepartmentScaleDetail(futureDepartmentScaleDetail mp.FutureDepartmentScaleDetail) (err error) {
	err = global.GVA_DB.Create(&futureDepartmentScaleDetail).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteFutureDepartmentScaleDetail
//@description: 删除FutureDepartmentScaleDetail记录
//@param: futureDepartmentScaleDetail model.FutureDepartmentScaleDetail
//@return: err error

func DeleteFutureDepartmentScaleDetail(futureDepartmentScaleDetail mp.FutureDepartmentScaleDetail) (err error) {
	err = global.GVA_DB.Delete(&futureDepartmentScaleDetail).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteFutureDepartmentScaleDetailByIds
//@description: 批量删除FutureDepartmentScaleDetail记录
//@param: ids request.IdsReq
//@return: err error

func DeleteFutureDepartmentScaleDetailByIds(ids rp.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]mp.FutureDepartmentScaleDetail{}, "id in ?", ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateFutureDepartmentScaleDetail
//@description: 更新FutureDepartmentScaleDetail记录
//@param: futureDepartmentScaleDetail *model.FutureDepartmentScaleDetail
//@return: err error

func UpdateFutureDepartmentScaleDetail(futureDepartmentScaleDetail mp.FutureDepartmentScaleDetail) (err error) {
	err = global.GVA_DB.Save(&futureDepartmentScaleDetail).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetFutureDepartmentScaleDetail
//@description: 根据id获取FutureDepartmentScaleDetail记录
//@param: id uint
//@return: err error, futureDepartmentScaleDetail model.FutureDepartmentScaleDetail

func GetFutureDepartmentScaleDetail(id uint) (err error, futureDepartmentScaleDetail mp.FutureDepartmentScaleDetail) {
	err = global.GVA_DB.Where("id = ?", id).Order("time desc").First(&futureDepartmentScaleDetail).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetFutureDepartmentScaleDetailInfoList
//@description: 分页获取FutureDepartmentScaleDetail记录
//@param: info request.FutureDepartmentScaleDetailSearch
//@return: err error, list interface{}, total int64

func GetFutureDepartmentScaleDetailInfoList(info rp.FutureDepartmentScaleDetailSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&mp.FutureDepartmentScaleDetail{}).Order("time desc")
	var futureDepartmentScaleDetails []mp.FutureDepartmentScaleDetail
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Time != "" {
		db = db.Where("time = ?", info.Time)
	}
	if info.DepartmentName != "" {
		db = db.Where("department_name LIKE ?", "%"+info.DepartmentName+"%")
	}
	if info.Type != "" {
		db = db.Where("type LIKE ?", "%"+info.Type+"%")
	}
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&futureDepartmentScaleDetails).Error
	return err, futureDepartmentScaleDetails, total
}
