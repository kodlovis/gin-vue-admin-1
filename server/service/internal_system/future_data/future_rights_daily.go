package future_data

import (
	"gin-vue-admin/global"
	mif "gin-vue-admin/model/internal_system/future_data"
	rif "gin-vue-admin/model/request/internal_system/future_data"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateFutureRightsDaily
//@description: 创建FutureRightsDaily记录
//@param: futureRightsDaily model.FutureRightsDaily
//@return: err error

func CreateFutureRightsDaily(futureRightsDaily mif.FutureRightsDaily) (err error) {
	err = global.GVA_DB.Create(&futureRightsDaily).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteFutureRightsDaily
//@description: 删除FutureRightsDaily记录
//@param: futureRightsDaily model.FutureRightsDaily
//@return: err error

func DeleteFutureRightsDaily(futureRightsDaily mif.FutureRightsDaily) (err error) {
	err = global.GVA_DB.Delete(&futureRightsDaily).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteFutureRightsDailyByIds
//@description: 批量删除FutureRightsDaily记录
//@param: ids request.IdsReq
//@return: err error

func DeleteFutureRightsDailyByIds(ids rif.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]mif.FutureRightsDaily{}, "id in ?", ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateFutureRightsDaily
//@description: 更新FutureRightsDaily记录
//@param: futureRightsDaily *model.FutureRightsDaily
//@return: err error

func UpdateFutureRightsDaily(futureRightsDaily mif.FutureRightsDaily) (err error) {
	err = global.GVA_DB.Save(&futureRightsDaily).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetFutureRightsDaily
//@description: 根据id获取FutureRightsDaily记录
//@param: id uint
//@return: err error, futureRightsDaily model.FutureRightsDaily

func GetFutureRightsDaily(id uint) (err error, futureRightsDaily mif.FutureRightsDaily) {
	err = global.GVA_DB.Where("id = ?", id).Order("time desc").First(&futureRightsDaily).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetFutureRightsDailyInfoList
//@description: 分页获取FutureRightsDaily记录
//@param: info request.FutureRightsDailySearch
//@return: err error, list interface{}, total int64

func GetFutureRightsDailyInfoList(info rif.FutureRightsDailySearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&mif.FutureRightsDaily{}).Order("time desc")
	var futureRightsDailys []mif.FutureRightsDaily
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Time != "" {
		db = db.Where("time LIKE ?", "%"+info.Time+"%")
	}
	if info.Department != "" {
		db = db.Where("department LIKE ?", "%"+info.Department+"%")
	}
	if info.Product != "" {
		db = db.Where("product LIKE ?", "%"+info.Product+"%")
	}
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&futureRightsDailys).Error
	return err, futureRightsDailys, total
}
