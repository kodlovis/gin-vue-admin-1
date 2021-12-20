package scale_rights

import (
	"gin-vue-admin/global"
	mp "gin-vue-admin/model/internal_system/scale_rights"
	rp "gin-vue-admin/model/request/internal_system/scale_rights"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateFutureDeparmentRights
//@description: 创建FutureDeparmentRights记录
//@param: futureDeparmentRights model.FutureDeparmentRights
//@return: err error

func CreateFutureDeparmentRights(futureDeparmentRights mp.FutureDeparmentRights) (err error) {
	err = global.GVA_DB.Create(&futureDeparmentRights).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteFutureDeparmentRights
//@description: 删除FutureDeparmentRights记录
//@param: futureDeparmentRights model.FutureDeparmentRights
//@return: err error

func DeleteFutureDeparmentRights(futureDeparmentRights mp.FutureDeparmentRights) (err error) {
	err = global.GVA_DB.Delete(&futureDeparmentRights).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteFutureDeparmentRightsByIds
//@description: 批量删除FutureDeparmentRights记录
//@param: ids request.IdsReq
//@return: err error

func DeleteFutureDeparmentRightsByIds(ids rp.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]mp.FutureDeparmentRights{}, "id in ?", ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateFutureDeparmentRights
//@description: 更新FutureDeparmentRights记录
//@param: futureDeparmentRights *model.FutureDeparmentRights
//@return: err error

func UpdateFutureDeparmentRights(futureDeparmentRights mp.FutureDeparmentRights) (err error) {
	err = global.GVA_DB.Save(&futureDeparmentRights).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetFutureDeparmentRights
//@description: 根据id获取FutureDeparmentRights记录
//@param: id uint
//@return: err error, futureDeparmentRights model.FutureDeparmentRights

func GetFutureDeparmentRights(id uint) (err error, futureDeparmentRights mp.FutureDeparmentRights) {
	err = global.GVA_DB.Where("id = ?", id).Order("time desc").First(&futureDeparmentRights).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetFutureDeparmentRightsInfoList
//@description: 分页获取FutureDeparmentRights记录
//@param: info request.FutureDeparmentRightsSearch
//@return: err error, list interface{}, total int64

func GetFutureDeparmentRightsInfoList(info rp.FutureDeparmentRightsSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&mp.FutureDeparmentRights{}).Order("time desc")
	var futureDeparmentRightss []mp.FutureDeparmentRights
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Time != "" {
		db = db.Where("time = ?", info.Time)
	}
	if info.Department != "" {
		db = db.Where("department LIKE ?", "%"+info.Department+"%")
	}
	if info.Type != "" {
		db = db.Where("type LIKE ?", "%"+info.Type+"%")
	}
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&futureDeparmentRightss).Error
	return err, futureDeparmentRightss, total
}
